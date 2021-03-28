pragma solidity >=0.7.0 <0.9.0;
pragma experimental ABIEncoderV2;

/**
    * @notice this interface is taken from indexed-core commit hash dae7f231d0f58bfc0993f6c01199cd6b74b01895
 */
interface IndexPoolI {
  function getDenormalizedWeight(address token) external view returns (uint256);
  function getBalance(address token) external view returns (uint256);
  function getUsedBalance(address token) external view returns (uint256);
  function getSpotPrice(address tokenIn, address tokenOut) external view returns (uint256);    
}

interface ERC20I {
    function totalSupply() external view returns (uint256);
}

/**
    * @notice SimpleMultiCall is a multicall-like contract for reading IndexPool information
    * @notice it is intended to minimize the need for manual abi encoding/decoding
    * @notice and leverage Golang's abigen to do the heavy lifting
 */
contract SimpleMultiCall {

    struct Bundle {
        address pool;
        address[] tokens;
        uint256[] denormalizedWeights;
        uint256[] balances;
        uint256[] totalSupplies;
    }
    
    function getBundle(
        address poolAddress,
        address[] memory tokens
    )
        public
        view
        returns (Bundle memory)
    {
        // order of the elements will be based on the ordering of the tokens
        // so we can ignore the return values of the address array
        (, uint256[] memory weights) = getDenormalizedWeights(poolAddress, tokens);
        (, uint256[] memory balances) = getBalances(poolAddress, tokens);
        (, uint256[] memory totalSupplies) = getTotalSupplies(tokens);
        return Bundle({
            pool: poolAddress,
            tokens: tokens,
            denormalizedWeights: weights,
            balances: balances,
            totalSupplies: totalSupplies
        });
    }

    function getBundles(
        address[] memory pools,
        address[][] memory tokens
    )
        public
        view
        returns (Bundle[] memory)
    {
        require(pools.length == tokens.length, "mismatching pools and tokens length");
        Bundle[] memory bundles = new Bundle[](pools.length);
        for (uint256 i = 0; i < tokens.length; i++) {
            bundles[i] = getBundle(pools[i], tokens[i]);
        }
        return bundles;
    }

    // index pool methods

    function getDenormalizedWeights(
        address poolAddress,
        address[] memory tokens
    ) 
        public 
        view
        returns (address[] memory, uint256[] memory) 
    {
        uint256[] memory weights = new uint256[](tokens.length);
        for (uint256 i = 0; i < tokens.length; i++) {
            weights[i] = IndexPoolI(poolAddress).getDenormalizedWeight(tokens[i]);
        }
        return (tokens, weights);
    }

    function getBalances(
        address poolAddress,
        address[] memory tokens
    ) 
        public 
        view
        returns (address[] memory, uint256[] memory) 
    {
        uint256[] memory balances = new uint256[](tokens.length);
        for (uint256 i = 0; i < tokens.length; i++) {
            balances[i] = IndexPoolI(poolAddress).getBalance(tokens[i]);
        }
        return (tokens, balances);
    }

    function getUsedBalances(
        address poolAddress,
        address[] memory tokens
    ) 
        public 
        view
        returns (address[] memory, uint256[] memory) 
    {
        uint256[] memory balances = new uint256[](tokens.length);
        for (uint256 i = 0; i < tokens.length; i++) {
            balances[i] = IndexPoolI(poolAddress).getUsedBalance(tokens[i]);
        }
        return (tokens, balances);
    }

    function getSpotPrices(
        address poolAddress,
        address[] memory inTokens,
        address[] memory outTokens
    )
        public
        view 
        returns (address[] memory, address[] memory, uint256[] memory)
    {
        require(inTokens.length == outTokens.length);
        uint256[] memory prices = new uint256[](inTokens.length);
        for (uint256 i = 0; i < inTokens.length; i++) {
            prices[i] = IndexPoolI(poolAddress).getSpotPrice(inTokens[i], outTokens[i]);
        }
        return (inTokens, outTokens, prices);
    }

    // erc20 methods

    function getTotalSupplies(
        address[] memory tokens
    )
        public
        view
        returns (address[] memory, uint256[] memory)
    {
        uint256[] memory supplies = new uint256[](tokens.length);
        for (uint256 i = 0; i < tokens.length; i++) {
            supplies[i] = ERC20I(tokens[i]).totalSupply();
        }
        return (tokens, supplies);
    }
}