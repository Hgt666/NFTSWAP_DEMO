// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/TestNFT.sol";
import "../src/NFTTradeMarket.sol";

contract Deploy is Script {
    function run() external {
        // 开启广播，使用当前钱包签名交易
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署 NFT 铸造合约
        TestNFT nft = new TestNFT(
            "TestNFT Collection",
            "TNFT",
            "ipfs://test-meta/"
        );
        console.log("TestNFT Address:", address(nft));

        // 2. 部署 NFT 交易市场
        NFTTradeMarket market = new NFTTradeMarket();
        console.log("NFTTradeMarket Address:", address(market));

        vm.stopBroadcast();
    }
}