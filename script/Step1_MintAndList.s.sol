// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/TestNFT.sol";
import "../src/NFTTradeMarket.sol";
import {IERC721} from "@openzeppelin/token/ERC721/IERC721.sol";

contract Step1_MintAndList is Script {
    address NFT_ADDR = 0xFe2aBc15793204656C6bA51D0d6b4741B16A1895;
    address MARKET_ADDR = 0x8fF9fdf34811CBB03600b032929c83028Fa89095;

    uint256 public constant LIST_PRICE = 0.0001 ether;
    uint256 public constant NEW_PRICE = 0.00015 ether;

    function run() external {
        uint256 sellerPk = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(sellerPk);

        // 1. 批量铸造 10 个 NFT
        TestNFT(NFT_ADDR).batchMint(10);
        console.log(" Batch mint 10 NFT done");

        // 2. 授权市场合约转移 NFT
        IERC721(NFT_ADDR).setApprovalForAll(MARKET_ADDR, true);
        console.log(" Approval done");

        // 3. 逐条挂单 token1 ~ token5
        NFTTradeMarket(MARKET_ADDR).listNFT(NFT_ADDR, 1, LIST_PRICE);
        NFTTradeMarket(MARKET_ADDR).listNFT(NFT_ADDR, 2, LIST_PRICE);
        NFTTradeMarket(MARKET_ADDR).listNFT(NFT_ADDR, 3, LIST_PRICE);
        NFTTradeMarket(MARKET_ADDR).listNFT(NFT_ADDR, 4, LIST_PRICE);
        NFTTradeMarket(MARKET_ADDR).listNFT(NFT_ADDR, 5, LIST_PRICE);
        console.log(" List token 1~5 done");

        // 4. 修改 token1 挂单价格为 0.15 ETH
        NFTTradeMarket(MARKET_ADDR).editListOrder(NFT_ADDR, 1, NEW_PRICE);
        console.log(" Edit token1 price done");

        vm.stopBroadcast();
    }
}


