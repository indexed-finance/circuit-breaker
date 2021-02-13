# Architecture

`circuit-breaker` consists of a Golang microservice backend with two main services known as `block-listener` and `contract-watcher`, along with a database backend using PostgreSQL. To minimize RPC calls and maximize processing speed, a "SimpleMultiCall" contract is used to batch together RPC calls.

## simplemulticall

The "SimpleMultiCall" contract is used to fetch the following information for all tokens helds by an IndexPool:

* Balances of all tokens held by the IndexPool
* Denormalized weights for all tokens held by the IndexPool
* Total supplies for all tokens held by the IndexPool

## postgresql

The postgresql database is used to store information for each monitored IndexPool. We store information for the 512 most recent blocks. The information tracked is as follows:

* Name of the pool
* Contract address of the pool
* The last block the information was updated at
* Decimals of all tokens held by the IndexPool
  * This information is gathered only once when the IndexPool record is first created in the database
* Balances of all tokens held by the IndexPool
  * Gathered every block
* Denormalized weights for all tokens held by the IndexPool
  * Gathered every block
* Total supplies for all tokens held by the IndexPool
  * Gathered every block

## block-listener

The `block-listener` service uses the `newHeads` ETH-RPC call to subscribe to notifications about the current blockchain head. Every time we receive a notification about a new block, a goroutine is spawned for each IndexPool we are monitoring to allow processing in parallel. The goroutines poll the SimpleMultiCall contract for the following information:

* Token balances for all tokens held by the IndexPool
* Denormalized weights for all tokens held by IndexPool
* Total supply of the various tokens held by the IndexPool

This information is then persisted into the database. After which we compare the information from this block to the information from the previous block, specifically looking for differences in total supplies. If the total supplies have **increased** or **decreased** by a percentage outside of an acceptable range (specified via a configuration file), we break the circuit for the given IndexPool temporarily disabling swaps. This is followed up by sending an SMS alert to all configured parties indicating that a circuit was broken.

## contract-watcher

The `contract-watcher` service is responsible for monitoring events, specifically the `LOG_SWAP` event, emitted by IndexPool contracts. Each monitored IndexPool has its own dedicated goroutine listening for events. Whenever a new event is dedicated, we first check to see if it is marked as "Removed" which indicates the event was removed by a blockchain reorganization. If it has been removed we skip processing that particular event.

We then fetch the following information:

* Swap fee for the block previous to the one the transaction which emitted the event was mined in
  * This is fetched from the blockchain
* Spot price for the swap direction in the block the transaction was mined in
  * This is fetched from the blockchain
* Spot price for the same swap direction for the block previous to the one the transaction which emitted the event was mined in
  * This is fetched from the database
  * This used the swap fee from the previous block in the calculation

We then use this information to calculate the price fluctuation between the last two blocks, and the price has **decreased** by a percentage outside of an acceptable range (specified via a configuration file), we break the circuit for the given IndexPool temporarily disabling swaps. This is followed up by sending an SMS alert to all configured parties indicating that a circuit was broken.