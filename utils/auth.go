package utils

import (
	"crypto/ecdsa"
	"io/ioutil"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// reusable ethereum key authorization tooling

// Authorizer wraps ethereum transactopts
type Authorizer struct {
	sync.Mutex // bind.TransactOpts is not thread safe
	*bind.TransactOpts
}

// NewAuthorizer returns an Authorizer object
func NewAuthorizer(keyFile, keyPass string) (*Authorizer, error) {
	fileBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	pk, err := keystore.DecryptKey(fileBytes, keyPass)
	if err != nil {
		return nil, err
	}
	return NewAuthorizerFromPK(pk.PrivateKey), nil
}

// NewAuthorizerFromPK returns an authorizer from a private key
func NewAuthorizerFromPK(pk *ecdsa.PrivateKey) *Authorizer {
	return &Authorizer{TransactOpts: bind.NewKeyedTransactor(pk)}
}

// NewAccount returns a new ethereum account as generated by `geth account new`
func NewAccount(keyFileDir, keyPass string) (accounts.Account, error) {
	return keystore.StoreKey(keyFileDir, keyPass, keystore.StandardScryptN, keystore.StandardScryptP)
}
