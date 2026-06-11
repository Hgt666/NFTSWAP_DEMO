// nft_list_order 挂单表
CREATE TABLE IF NOT EXISTS nft_list_order (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nft_contract TEXT NOT NULL,    -- NFT合约地址
    token_id INTEGER NOT NULL,     -- NFT tokenId
    seller TEXT NOT NULL,         -- 卖家地址
    price TEXT NOT NULL,          -- 挂单价格(wei字符串)
    status TINYINT NOT NULL DEFAULT 1, -- 1-有效 0-已取消/已成交
    block_number INTEGER NOT NULL, -- 区块高度
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(nft_contract, token_id)
);



// nft_bid_recode 出价记录表
CREATE TABLE IF NOT EXISTS nft_bid_record (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nft_contract TEXT NOT NULL,
    token_id INTEGER NOT NULL,
    bidder TEXT NOT NULL,          -- 出价人地址
    bid_price TEXT NOT NULL,       -- 出价金额(wei)
    status TINYINT NOT NULL DEFAULT 1, -- 1-有效 0-失效
    block_number INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

// nft_trade_record 交易记录表
CREATE TABLE IF NOT EXISTS nft_trade_record (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nft_contract TEXT NOT NULL,
    token_id INTEGER NOT NULL,
    seller TEXT NOT NULL,
    buyer TEXT NOT NULL,
    trade_price TEXT NOT NULL,     -- 成交价格(wei)
    trade_type TINYINT NOT NULL,   -- 1:限价撮合 2:市价购买 3:卖家接受出价
    block_number INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);