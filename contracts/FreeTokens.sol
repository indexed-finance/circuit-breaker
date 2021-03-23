pragma solidity >=0.8.0 <0.9.0;


interface FreeTokens {
    function getFreeTokens(address recipient, uint256 amunt) external;
}