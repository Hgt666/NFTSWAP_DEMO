# 需求文档
# 1. 项目名称：NFTSwap
# 2. 项目描述：一个基于区块链的NFT交易平台，支持多种NFT类型和交易方式。
# 3、项目功能：NFT 挂单、改价、撤单、出价、取消订单、撮合成交等功能
# 4、需求实现：
    1、链上数据索引
    2、提供 NFT 相关的API 接口，包括 NFT 挂单列表、NFT 详情、单个 NFT 出价记录、NFT 无数据等

# 5、 数据库设计
    1、NFT 挂单表：
        id
        nft_contract - NFT 合约地址
        nft_id - NFT ID
        symbol - NFT symbol
        owner - NFT 所有者地址
        price - NFT 价格
        status - NFT 挂单状态(0:未成交，1:已成交)
        created_at - 创建时间
        updated_at - 更新时间
    2、NFT 出价表
        id 
        nft_id - NFT ID
        price - 出价价格
        buyer - 出价者地址
        created_at - 创建时间
        updated_at - 更新时间
    3、成交记录表
        id 
        nft_id - NFT ID
        buyer - 成交者地址
        seller - 卖家地址
        price - 成交价格
        created_at - 成交时间
    4、NFT 元数据表
        id 
        nft_id - NFT ID
        metadata - NFT 元数据
        created_at - 创建时间
        updated_at - 更新时间
    5、扫链记录表
        id 
        nft_contract - NFT 合约地址
        chain_id - 链ID
        lastIndexBlock - 最后一个区块索引
        lastIndexTxHash - 最后一个交易哈希

# 6、项目分层目录
    - cmd
       - api
            - main.go
       - scanner
            - main.go
    - config
        - config.go
        - config.yaml
    - db
        - mysql.go
        - redis.go
    - mq
        - rabbitmq.go
    - contract
        - abi
        - market.go
    - utils
        - utils.go
    - models
        - nft_list.go
        - nft_bid.go
        - nft_trade_record.go
        - nft_meta.go
        - sync_index_status.go
    - router
        - router.go
    - service
        - nft_service.go
        - nft_bid_service.go
        - nft_trade_record_service.go
        - nft_meta_service.go
        - sync_index_service.go
    - dao
        - nft_list_dao.go
        - nft_bid_dao.go
        - nft_trade_record_dao.go
        - nft_meta_dao.go
        - sync_index_dao.go
    - main.go
    - go.mod
    - go.sum
    - README.md
    
# 7、接口设计
    1、NFT 挂单列表
        GET /api/nft/list
    2、NFT 详情
        GET /api/nft/detail/{nft_id}
    3、单个 NFT 出价记录
        GET /api/nft/bid/{nft_id}
    4、NFT 元数据
        GET /api/nft/meta/{nft_id}
    5、NFT 成交价排序 top 3
        GET /api/nft/top10
    6、NFT 挂单状态
        GET /api/nft/status/{nft_id}



    


