-- 区块断点记录表（实现断点续传）
CREATE TABLE `sync_block_point` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sync_key` varchar(64) NOT NULL COMMENT '同步标识(唯一)',
  `last_block_num` bigint NOT NULL DEFAULT 0 COMMENT '最后已处理区块高度',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_sync_key` (`sync_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区块同步断点';

-- 初始化一条记录
INSERT INTO sync_block_point(sync_key, last_block_num) VALUES ('nft_market_event', 0);


-- 挂单表 nft_list_order（加索引）
CREATE TABLE `nft_list_order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nft_contract` varchar(42) NOT NULL COMMENT 'NFT合约地址',
  `token_id` bigint NOT NULL COMMENT 'NFT TokenID',
  `seller` varchar(42) NOT NULL COMMENT '卖家地址',
  `price_wei` varchar(66) NOT NULL COMMENT '价格(wei)',
  `price_eth` decimal(20,8) NOT NULL DEFAULT 0 COMMENT '价格(ETH)',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '1有效 0失效',
  `block_number` bigint NOT NULL COMMENT '区块高度',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_contract_token` (`nft_contract`,`token_id`),
  KEY `idx_seller` (`seller`),
  KEY `idx_status` (`status`),
  KEY `idx_block` (`block_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='NFT挂单表';



-- 出价表 nft_bid_record（加索引）
CREATE TABLE `nft_bid_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nft_contract` varchar(42) NOT NULL,
  `token_id` bigint NOT NULL,
  `bidder` varchar(42) NOT NULL COMMENT '出价人',
  `bid_price_wei` varchar(66) NOT NULL,
  `bid_price_eth` decimal(20,8) NOT NULL DEFAULT 0,
  `status` tinyint NOT NULL DEFAULT 1,
  `block_number` bigint NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_contract_token` (`nft_contract`,`token_id`),
  KEY `idx_bidder` (`bidder`),
  KEY `idx_block` (`block_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='NFT出价记录';


-- 成交表 nft_trade_record（分表预备 + 索引）
CREATE TABLE `nft_trade_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nft_contract` varchar(42) NOT NULL,
  `token_id` bigint NOT NULL,
  `seller` varchar(42) NOT NULL,
  `buyer` varchar(42) NOT NULL,
  `trade_price_wei` varchar(66) NOT NULL,
  `trade_price_eth` decimal(20,8) NOT NULL DEFAULT 0,
  `trade_type` tinyint NOT NULL COMMENT '1限价 2市价 3卖家接单',
  `block_number` bigint NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_contract_token` (`nft_contract`,`token_id`),
  KEY `idx_seller` (`seller`),
  KEY `idx_buyer` (`buyer`),
  KEY `idx_block` (`block_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='NFT成交记录';


-- NFT 元数据表（对接 IPFS）
CREATE TABLE `nft_metadata` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nft_contract` varchar(42) NOT NULL,
  `token_id` bigint NOT NULL,
  `ipfs_cid` varchar(255) NOT NULL COMMENT 'IPFS CID',
  `name` varchar(128) DEFAULT '' COMMENT 'NFT名称',
  `description` text COMMENT '描述',
  `image` varchar(255) DEFAULT '' COMMENT '图片链接',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_contract_token` (`nft_contract`,`token_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='NFT元数据';