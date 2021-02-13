package database

import (
	"encoding/json"
	"errors"
	"math/big"
	"strconv"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"modernc.org/sortutil"
)

// Pool is a single IndexPool
type Pool struct {
	gorm.Model
	Name            string `gorm:"unique"`
	ContractAddress string
	// denotes the last block that we did an update at
	// this is to make sure we dont needlessly update the information
	LastUpdateAt uint64
	// maps block num -> info
	Infos datatypes.JSONMap
	// TokenDecimals tracks the decimal values for a given token
	// at the moment we expect this to be fixed and not updated that often
	// decimal value is of type uint8
	TokenDecimals datatypes.JSONMap
}

// PoolInfo records pool information for a given block
type PoolInfo struct {
	// tokenName -> balance (big.Int)
	Balances datatypes.JSONMap `json:"balances"`
	// tokenName -> weight (big.Int)
	DenormalizedWeights datatypes.JSONMap `json:"denormalized_weights"`
	// TokenTotalSupplies tracks the total supply for all tokens held by the indice
	// this is not to be confused with the IndexPool tokens themselves. For example
	// DEFI5 holds UNI, SNX, AAVE, CMP and CRV so this will track the supply fo rthose
	// supply value is of type string
	TokenTotalSupplies datatypes.JSONMap `json:"token_total_supplies"`
}

// NewPool creates a new entry in the database for an index pool
func (db *Database) NewPool(
	name string,
	contractAddress string,
	block uint64,
	balances, denormWeights, tokenTotalSupplies, decimals map[string]interface{},
) error {
	name = strings.ToLower(name)
	db.namedLock.Lock(name)
	defer db.namedLock.Unlock(name)
	infos := make(map[string]interface{})
	infos[strconv.FormatUint(block, 10)] = PoolInfo{
		Balances:            balances,
		DenormalizedWeights: denormWeights,
		TokenTotalSupplies:  tokenTotalSupplies,
	}
	return db.db.Create(&Pool{
		Name:            strings.ToLower(name),
		ContractAddress: contractAddress,
		LastUpdateAt:    block,
		TokenDecimals:   decimals,
		Infos:           infos,
	}).Error
}

// RecordInfo is used to record updates infos for the given block
func (db *Database) RecordInfo(
	name string,
	block uint64,
	balances, denormWeights, tokenTotalSupplies map[string]interface{},
) error {
	name = strings.ToLower(name)
	db.namedLock.Lock(name)
	defer db.namedLock.Unlock(name)
	return db.db.Transaction(func(tx *gorm.DB) error {
		var pool Pool
		if err := tx.Model(&Pool{}).Where(
			"name = ?", name,
		).First(&pool).Error; err != nil {
			return err
		}
		if pool.LastUpdateAt == block {
			// no need to update
			return nil
		}
		pool.LastUpdateAt = block
		pool.Infos[strconv.FormatUint(block, 10)] = PoolInfo{
			Balances:            balances,
			DenormalizedWeights: denormWeights,
			TokenTotalSupplies:  tokenTotalSupplies,
		}
		return tx.Save(&pool).Error
	})
}

// GetInfo returns the record PoolInfo at the given block
func (db *Database) GetInfo(name string, block uint64) (PoolInfo, error) {
	name = strings.ToLower(name)
	db.namedLock.RLock(name)
	defer db.namedLock.RUnlock(name)
	var pool Pool
	if err := db.db.Where("name = ?", name).First(&pool, datatypes.JSONQuery("infos").HasKey(strconv.FormatUint(block, 10))).Error; err != nil {
		return PoolInfo{}, err
	}
	data, err := json.Marshal(pool.Infos[strconv.FormatUint(block, 10)])
	if err != nil {
		return PoolInfo{}, err
	}
	var info PoolInfo
	return info, json.Unmarshal(data, &info)
}

// GetPool returns a pool by its index name
func (db *Database) GetPool(name string) (*Pool, error) {
	name = strings.ToLower(name)
	db.namedLock.RLock(name)
	defer db.namedLock.RUnlock(name)
	var pool Pool
	return &pool, db.db.Model(&Pool{}).Where("name = ?", name).First(&pool).Error
}

// GetTokenDecimals returns the decimal count associated with various tokens
func (db *Database) GetTokenDecimals(poolName, tokenName string) (uint8, error) {
	poolName = strings.ToLower(poolName)
	tokenName = strings.ToLower(tokenName)
	db.namedLock.RLock(poolName)
	defer db.namedLock.RUnlock(poolName)
	var pool Pool
	if err := db.db.Where("name = ?", poolName).First(
		&pool,
		datatypes.JSONQuery("token_decimals").HasKey(tokenName),
	).Error; err != nil {
		return 0, err
	}
	strDec, ok := pool.TokenDecimals[tokenName].(string)
	if !ok {
		return 0, errors.New("tokenDecimal is not of type string")
	}
	dec, err := strconv.ParseUint(strDec, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint8(dec), nil
}

// GetNumInfos returns the number of infos recorded for a pool
func (db *Database) GetNumInfos(poolName string) (int, error) {
	poolName = strings.ToLower(poolName)
	db.namedLock.RLock(poolName)
	defer db.namedLock.RUnlock(poolName)
	var pool Pool
	if err := db.db.Model(&Pool{}).Where("name = ?", poolName).First(&pool).Error; err != nil {
		return 0, err
	}
	return len(pool.Infos), nil
}

// PurgeOldInfos is used to purge old information from the database
func (db *Database) PurgeOldInfos(poolName string, purgeCount int) error {
	poolName = strings.ToLower(poolName)
	db.namedLock.Lock(poolName)
	defer db.namedLock.Unlock(poolName)

	return db.db.Transaction(func(tx *gorm.DB) error {
		var pool Pool
		if err := tx.Model(&Pool{}).Where("name = ?", poolName).First(&pool).Error; err != nil {
			return err
		}
		// no need to purge
		if len(pool.Infos) < purgeCount {
			return nil
		}
		//var bigInts sortutil.BigIntSlice
		bigInts := make(sortutil.BigIntSlice, 0, len(pool.Infos))
		for blockNum := range pool.Infos {
			bigBlockNum, ok := new(big.Int).SetString(blockNum, 10)
			if !ok {
				return errors.New("failed to convert block number from string to big.Int")
			}
			bigInts = append(bigInts, bigBlockNum)
		}
		bigInts.Sort()
		for i := 0; i < purgeCount; i++ {
			delete(pool.Infos, bigInts[i].String())

		}
		return tx.Save(&pool).Error
	})
}
