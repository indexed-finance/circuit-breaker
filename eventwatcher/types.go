package eventwatcher

const (
	// LogSwapABI is the ABI for LOG_SWAP events
	LogSwapABI = `{
		"anonymous": false,
		"inputs": [
		  {
			"indexed": true,
			"internalType": "address",
			"name": "caller",
			"type": "address"
		  },
		  {
			"indexed": true,
			"internalType": "address",
			"name": "tokenIn",
			"type": "address"
		  },
		  {
			"indexed": true,
			"internalType": "address",
			"name": "tokenOut",
			"type": "address"
		  },
		  {
			"indexed": false,
			"internalType": "uint256",
			"name": "tokenAmountIn",
			"type": "uint256"
		  },
		  {
			"indexed": false,
			"internalType": "uint256",
			"name": "tokenAmountOut",
			"type": "uint256"
		  }
		],
		"name": "LOG_SWAP",
		"type": "event"
	}`
)
