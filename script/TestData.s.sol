// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/NFTTradeMarket.sol";

contract Step2_TradeOperate is Script {
    address NFT_ADDR = 0x5FbDB2315678afecb367f032d93F642f64180aa3;
    address MARKET_ADDR = 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512;

    uint256 public constant BID1 = 0.12 ether;
    uint256 public constant BID2 = 0.16 ether;
    uint256 public constant BUY_PRICE = 0.15 ether;
    uint256 public constant BID_TOKEN3 = 0.13 ether;

    function run() external {
        NFTTradeMarket market = NFTTradeMarket(MARKET_ADDR);

        uint256 sellerPk = vm.envUint("PRIVATE_KEY");
        uint256 buyer1Pk = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d;
        uint256 buyer2Pk = 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a;
        uint256 buyer3Pk = 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6;

        vm.deal(vm.addr(buyer1Pk), 10 ether);
        vm.deal(vm.addr(buyer2Pk), 10 ether);
        vm.deal(vm.addr(buyer3Pk), 10 ether);

        // ========== 第一步：先取消 token5 挂单（此时所有订单都是有效状态） ==========
        // vm.startBroadcast(sellerPk);
        // market.cancelListOrder(NFT_ADDR, 5);
        // vm.stopBroadcast();
        // console.log("Cancel token5 list done");

        // ========== 第二步：所有买家出价 ==========
        vm.startBroadcast(buyer1Pk);
        market.placeBid{value: BID1}(NFT_ADDR, 1);
        market.placeBid{value: BID1}(NFT_ADDR, 2);
        market.placeBid{value: BID_TOKEN3}(NFT_ADDR, 3);
        vm.stopBroadcast();
        console.log("Buyer1 bid done");

        vm.startBroadcast(buyer2Pk);
        market.placeBid{value: BID2}(NFT_ADDR, 1);
        vm.stopBroadcast();
        console.log(" Buyer2 higher bid done");

        // ========== 第三步：市价成交 token2 ==========
        vm.startBroadcast(buyer3Pk);
        market.buyMarket{value: BUY_PRICE}(NFT_ADDR, 2);
        vm.stopBroadcast();
        console.log("Market buy token2 done");

        // ========== 第四步：限价撮合 token1 ==========
        vm.startBroadcast();
        market.matchLimitOrder(NFT_ADDR, 1);
        vm.stopBroadcast();
        console.log("Limit match token1 done");

        // ========== 第五步：卖家接受 token3 出价 ==========
        vm.startBroadcast(sellerPk);
        market.acceptBid(NFT_ADDR, 3);
        vm.stopBroadcast();
        console.log("Seller accept bid token3 done");

        console.log("===== All trade operations finished =====");
    }
}