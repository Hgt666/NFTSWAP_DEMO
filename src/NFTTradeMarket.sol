// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import {
    ReentrancyGuard
} from "@openzeppelin/utils/ReentrancyGuard.sol";
import {IERC721} from "@openzeppelin/token/ERC721/IERC721.sol";
import {IERC721Receiver} from "@openzeppelin/token/ERC721/IERC721Receiver.sol";

/**
 * @title NFTTradeMarket
 * @dev NFT 交易市场：支持 List挂售、Bid出价、限价/市价、编辑订单、取消订单、自动撮合
 */
contract NFTTradeMarket is ReentrancyGuard, IERC721Receiver {
    // ===================== 数据结构定义 =====================
    // 挂售订单 List：NFT 卖家挂单
    struct ListOrder {
        address seller; // 卖家地址
        address nftContract; // NFT 合约地址
        uint256 tokenId; // NFT 编号
        uint256 price; // 限价(最低售价)
        bool active; // 订单是否有效
    }

    // 出价订单 Bid：买家针对指定 NFT 出价
    struct BidOrder {
        address bidder; // 出价人
        address nftContract;
        uint256 tokenId;
        uint256 price; // 出价金额
        bool active; // 出价是否有效
    }

    // Key: nftContract + tokenId => List 挂售单
    mapping(address => mapping(uint256 => ListOrder)) public listOrders;
    // Key: nftContract + tokenId => Bid 出价单列表(简化: 单NFT仅保留最高出价)
    mapping(address => mapping(uint256 => BidOrder)) public highestBid;

    // ===================== 事件定义(后端索引核心) =====================
    event OrderListed(
        address indexed seller,
        address indexed nftContract,
        uint256 indexed tokenId,
        uint256 price
    );

    event OrderEdited(
        address indexed seller,
        address indexed nftContract,
        uint256 indexed tokenId,
        uint256 newPrice
    );

    event OrderCancelled(
        address indexed seller,
        address indexed nftContract,
        uint256 indexed tokenId
    );

    event BidPlaced(
        address indexed bidder,
        address indexed nftContract,
        uint256 indexed tokenId,
        uint256 price
    );

    event BidCancelled(
        address indexed bidder,
        address indexed nftContract,
        uint256 indexed tokenId
    );

    event MatchSuccess(
        address indexed seller,
        address indexed buyer,
        address indexed nftContract,
        uint256 tokenId,
        uint256 price,
        uint8 tradeType // 1=限价成交 2=市价成交 3=卖家接受出价
    );

    // ===================== 构造函数 =====================
    constructor() {}

    // ERC721 接收回调
    function onERC721Received(
        address,
        address,
        uint256,
        bytes calldata
    ) external pure override returns (bytes4) {
        return this.onERC721Received.selector;
    }

    // ===================== 1. 卖家操作：挂售 List (限价单) =====================
    /**
     * @dev 挂单售卖 NFT (限价单)
     * @param nftContract NFT 合约地址
     * @param tokenId NFT 编号
     * @param price 最低售价(限价)
     */
    function listNFT(
        address nftContract,
        uint256 tokenId,
        uint256 price
    ) external nonReentrant {
        IERC721 nft = IERC721(nftContract);
        address owner = nft.ownerOf(tokenId);

        // 校验：必须是 NFT 持有者
        require(owner == msg.sender, "Not NFT owner");
        // 校验：该 NFT 未挂单
        ListOrder storage order = listOrders[nftContract][tokenId];
        require(!order.active, "Order already listed");
        // 价格必须大于0
        require(price > 0, "Price must > 0");

        // 将 NFT 转入市场合约托管
        nft.safeTransferFrom(msg.sender, address(this), tokenId);

        // 写入挂售订单
        listOrders[nftContract][tokenId] = ListOrder({
            seller: msg.sender,
            nftContract: nftContract,
            tokenId: tokenId,
            price: price,
            active: true
        });

        emit OrderListed(msg.sender, nftContract, tokenId, price);
    }

    // ===================== 2. 编辑挂售订单(修改限价) =====================
    function editListOrder(
        address nftContract,
        uint256 tokenId,
        uint256 newPrice
    ) external nonReentrant {
        ListOrder storage order = listOrders[nftContract][tokenId];
        // 校验：订单有效 + 仅卖家可编辑
        require(order.active, "Order not active");
        require(order.seller == msg.sender, "Not seller");
        require(newPrice > 0, "Price must > 0");

        // 更新价格
        order.price = newPrice;
        emit OrderEdited(msg.sender, nftContract, tokenId, newPrice);
    }

    // ===================== 3. 取消挂售订单 =====================
    function cancelListOrder(
        address nftContract,
        uint256 tokenId
    ) external nonReentrant {
        ListOrder storage order = listOrders[nftContract][tokenId];
        require(order.active, "Order not active");
        require(order.seller == msg.sender, "Not seller");

        // 归还 NFT 给卖家
        IERC721(nftContract).safeTransferFrom(
            address(this),
            msg.sender,
            tokenId
        );

        // 关闭订单
        order.active = false;
        emit OrderCancelled(msg.sender, nftContract, tokenId);
    }

    // ===================== 4. 买家操作：发起 Bid 出价 =====================
    function placeBid(
        address nftContract,
        uint256 tokenId
    ) external payable nonReentrant {
        uint256 bidPrice = msg.value;
        require(bidPrice > 0, "Bid must > 0");

        BidOrder storage oldBid = highestBid[nftContract][tokenId];
        // 新出价必须高于历史最高出价
        if (oldBid.active) {
            require(bidPrice > oldBid.price, "Bid not higher than current");
            // 退还上一个出价人的 ETH
            // payable(oldBid.bidder).transfer(oldBid.price);
            (bool success, )= payable(oldBid.bidder).call{value: oldBid.price}("");
            require(success, "Transfer failed");
        }

        // 更新为最新最高出价
        highestBid[nftContract][tokenId] = BidOrder({
            bidder: msg.sender,
            nftContract: nftContract,
            tokenId: tokenId,
            price: bidPrice,
            active: true
        });

        emit BidPlaced(msg.sender, nftContract, tokenId, bidPrice);
    }

    // ===================== 5. 取消自己的 Bid 出价 =====================
    function cancelBid(
        address nftContract,
        uint256 tokenId
    ) external nonReentrant {
        BidOrder storage bid = highestBid[nftContract][tokenId];
        require(bid.active, "Bid not active");
        require(bid.bidder == msg.sender, "Not bidder");

        // 退还出价 ETH
        payable(msg.sender).transfer(bid.price);
        // 清空出价
        bid.active = false;
        emit BidCancelled(msg.sender, nftContract, tokenId);
    }

    // ===================== 6. 自动撮合：市价单(买家直接吃限价单) =====================
    /**
     * @dev 市价购买：买家直接按卖家限价成交
     */
    function buyMarket(
        address nftContract,
        uint256 tokenId
    ) external payable nonReentrant {
        ListOrder storage listOrder = listOrders[nftContract][tokenId];
        require(listOrder.active, "No active list order");
        uint256 limitPrice = listOrder.price;
        // 支付金额 >= 限价 即可成交
        require(msg.value >= limitPrice, "Insufficient payment");

        address seller = listOrder.seller;
        address buyer = msg.sender;

        // 1. 转账 ETH 给卖家
        payable(seller).transfer(limitPrice);
        // 多余金额退回买家
        if (msg.value > limitPrice) {
            payable(buyer).transfer(msg.value - limitPrice);
        }

        // 2. 转移 NFT 给买家
        IERC721(nftContract).safeTransferFrom(address(this), buyer, tokenId);

        // 3. 关闭挂售单
        listOrder.active = false;

        // 4. 清空该 NFT 所有出价(订单完成)
        highestBid[nftContract][tokenId].active = false;

        // tradeType = 2 市价成交
        emit MatchSuccess(seller, buyer, nftContract, tokenId, limitPrice, 2);
    }

    // ===================== 7. 自动撮合：卖家接受最高出价(Bid 成交) =====================
    /**
     * @dev 卖家主动接受当前最高出价，完成交易
     */
    function acceptBid(
        address nftContract,
        uint256 tokenId
    ) external nonReentrant {
        ListOrder storage listOrder = listOrders[nftContract][tokenId];
        BidOrder storage bidOrder = highestBid[nftContract][tokenId];

        require(listOrder.active, "List order not active");
        require(bidOrder.active, "No valid bid");
        require(listOrder.seller == msg.sender, "Not seller");

        address seller = listOrder.seller;
        address buyer = bidOrder.bidder;
        uint256 dealPrice = bidOrder.price;

        // 1. ETH 转给卖家
        payable(seller).transfer(dealPrice);
        // 2. NFT 转给买家
        IERC721(nftContract).safeTransferFrom(address(this), buyer, tokenId);

        // 3. 关闭所有订单
        listOrder.active = false;
        bidOrder.active = false;

        // tradeType = 3 接受出价成交
        emit MatchSuccess(seller, buyer, nftContract, tokenId, dealPrice, 3);
    }

    // ===================== 8. 限价撮合(兜底逻辑) 外部主动触发 =====================
    /**
     * @dev 限价撮合：当出价 >= 卖家限价，可任意地址触发自动成交
     */
    function matchLimitOrder(
        address nftContract,
        uint256 tokenId
    ) external nonReentrant {
        ListOrder storage listOrder = listOrders[nftContract][tokenId];
        BidOrder storage bidOrder = highestBid[nftContract][tokenId];

        require(listOrder.active && bidOrder.active, "Order not valid");
        // 出价 >= 卖家限价，满足撮合条件
        require(bidOrder.price >= listOrder.price, "Bid lower than limit");

        address seller = listOrder.seller;
        address buyer = bidOrder.bidder;
        uint256 dealPrice = bidOrder.price;

        // 资产转移
        payable(seller).transfer(dealPrice);
        IERC721(nftContract).safeTransferFrom(address(this), buyer, tokenId);

        // 关闭订单
        listOrder.active = false;
        bidOrder.active = false;

        // tradeType = 1 限价自动撮合成交
        emit MatchSuccess(seller, buyer, nftContract, tokenId, dealPrice, 1);
    }

    // ===================== 视图函数(前端/后端查询) =====================
    function getListOrder(
        address nftContract,
        uint256 tokenId
    ) external view returns (ListOrder memory) {
        return listOrders[nftContract][tokenId];
    }

    function getHighestBid(
        address nftContract,
        uint256 tokenId
    ) external view returns (BidOrder memory) {
        return highestBid[nftContract][tokenId];
    }
}
