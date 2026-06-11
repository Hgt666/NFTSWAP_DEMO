// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/NFTTradeMarket.sol";

contract Step2_TradeOperate is Script {
    address NFT_ADDR = 0xFe2aBc15793204656C6bA51D0d6b4741B16A1895;
    address MARKET_ADDR = 0x8fF9fdf34811CBB03600b032929c83028Fa89095;

    // 价格常量
    uint256 public constant BID1 = 0.12 ether;
    uint256 public constant BID2 = 0.16 ether;
    uint256 public constant BUY_PRICE = 0.15 ether;
    uint256 public constant BID_TOKEN3 = 0.13 ether;

    function run() external {
        NFTTradeMarket market = NFTTradeMarket(MARKET_ADDR);

        // 账户私钥
        uint256 sellerPk = vm.envUint("PRIVATE_KEY");
        uint256 buyer1Pk = 0x66b649ffa33288788b18bdcc8df75d4c78d4f63e51680eac0cd2b073d97642f2;
        uint256 buyer2Pk = 0x107fcbade17ab87bdba4b10d8a9e384c2217851e498d590c1a0373043d09dc74;
        uint256 buyer3Pk = 0x8a8b9667b00069213058ddb2d6c78170ab0b128227e9944d8d216d7271b41b60;

        // 给测试账户充值 ETH
        // vm.deal(vm.addr(buyer1Pk), 10 ether);
        // vm.deal(vm.addr(buyer2Pk), 10 ether);
        // vm.deal(vm.addr(buyer3Pk), 10 ether);

        // 1. 撤销 token5 挂单（此时订单全部有效）
        vm.startBroadcast(sellerPk);
        market.cancelListOrder(NFT_ADDR, 5);
        vm.stopBroadcast();
        console.log(" Cancel token5 list done");

        // 2. 买家3 市价购买 token2（挂单有效，优先成交）
        vm.startBroadcast(buyer3Pk);
        market.buyMarket{value: BUY_PRICE}(NFT_ADDR, 2);
        vm.stopBroadcast();
        console.log("Market buy token2 done");

        // 3. 买家1 对 token1 / token3 出价（token2 已成交，不再操作）
        vm.startBroadcast(buyer1Pk);
        market.placeBid{value: BID1}(NFT_ADDR, 1);
        market.placeBid{value: BID_TOKEN3}(NFT_ADDR, 3);
        vm.stopBroadcast();
        console.log(" Buyer1 bid done");

        // 4. 买家2 对 token1 出更高价（0.16 > 0.15，满足撮合条件）
        vm.startBroadcast(buyer2Pk);
        market.placeBid{value: BID2}(NFT_ADDR, 1);
        vm.stopBroadcast();
        console.log(" Buyer2 higher bid done");

        // 5. 触发 token1 限价自动撮合
        vm.startBroadcast();
        market.matchLimitOrder(NFT_ADDR, 1);
        vm.stopBroadcast();
        console.log(" Limit match token1 done");

        // 6. 卖家接受 token3 出价
        vm.startBroadcast(sellerPk);
        market.acceptBid(NFT_ADDR, 3);
        vm.stopBroadcast();
        console.log(" Seller accept bid token3 done");

        console.log("===== All test data generated successfully =====");
    }
}