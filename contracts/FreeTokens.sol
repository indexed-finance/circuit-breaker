pragma solidity >=0.6.0 <0.8.0;


interface FreeTokens {
    function getFreeTokens(address recipient, uint256 amunt) external;
}