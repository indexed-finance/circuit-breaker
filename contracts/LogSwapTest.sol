pragma solidity >=0.7.0 <0.8.0;


contract LogSwapTest {
    event LOG_SWAP(
        address indexed caller,
        address indexed tokenIn,
        address indexed tokenOut,
        uint256 tokenAmountIn,
        uint256 tokenAmountOut
    );

    uint256 internal bigRando = 1;
    string public symbol = "CC10";
    uint8 public decimal = 18;
    uint256 public _swapFee = 25000000000000000; // 2.5%
    bool internal _publicSwap;

    function emitLogSwap(uint256 outAmount, uint256 inAmount) public {
        emit LOG_SWAP(msg.sender, msg.sender, msg.sender, inAmount, outAmount);
    }

    function setPublicSwap(bool enabled) external {
        _publicSwap = enabled;
    }

    // note: this is used to generate a random denorm weight for testing purposes
    // please dont use this in production it is not secure
    function getDenormalizedWeight(address token) public view returns (uint256) {
        uint256 baseWeight = uint256(keccak256(abi.encodePacked(block.timestamp, msg.sender, bigRando, token))) % 100;
        uint256 randWeight = baseWeight * 1 ether;
        return randWeight;
    }

    function getUsedBalance(address token) public view returns (uint256) {
        uint256 baseWeight = uint256(keccak256(abi.encodePacked(block.timestamp, msg.sender, bigRando, token)));
        uint256 randWeight = baseWeight * 0.25 ether;
        return randWeight;   
    }

    function getBalance(address token) public view returns (uint256) {
        uint256 baseWeight = uint256(keccak256(abi.encodePacked(block.timestamp, msg.sender, bigRando, token)));
        uint256 randWeight = baseWeight * 10 ether;
        return randWeight;   
    }

    function getCurrentTokens() public view returns (address[] memory) {
        address[] memory tokens = new address[](1);
        tokens[0] = address(this);
        return tokens;
    }

    function getSpotPrice(address tokenIn, address tokenOut) public view returns (uint256) {
        // override randomized checks as this is specific test logic which is dependent
        // within watched_contract_test.go
        // changing this value could break testing for that particular package
        // in the future the testing should be made less brittle but it is a bit diifficult at the moment
        // due to the complexity of the index pools
        if (block.number >= 3) {
            return 0.12 ether;
        }
        uint256 inPrice = uint256(keccak256(abi.encodePacked(block.timestamp, msg.sender, bigRando, tokenIn))) % 250;
        uint256 outPrice = uint256(keccak256(abi.encodePacked(block.timestamp, msg.sender, bigRando, tokenOut))) % 250;
        return (inPrice * (1 ether)) + (outPrice * (1 ether));
    }

    function totalSupply() public view returns (uint256) {
        uint256 baseWeight = uint256(keccak256(abi.encodePacked(block.timestamp, msg.sender, bigRando)));
        return (baseWeight * 100 ether);      
    }

    function getSwapFee() public view returns (uint256) {
        return _swapFee;
    }

    function incRando() public {
        bigRando += 1;
    }

}