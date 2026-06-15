package db

import (
	"log"
	"nft-service/config"
	"nft-service/models"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 初始化 mysql
func InitMySQL() error {
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GlobalConfig.MysqlConfig.User,
		config.GlobalConfig.MysqlConfig.Password,
		config.GlobalConfig.MysqlConfig.Host,
		config.GlobalConfig.MysqlConfig.Port,
		config.GlobalConfig.MysqlConfig.DBName)

	// 优化连接数配置
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Silent),
	})

	DB = db

	if err != nil {
		log.Fatal("mysql连接失败")
		return err
	}
	// 迁移数据库
	err = db.AutoMigrate(&models.NftBid{}, &models.NftList{}, &models.NftMeta{}, &models.NftTradeRecord{}, &models.SyncIndexStatus{})
	if err != nil {
		log.Fatal("mysql迁移失败")
		return err
	}
	return nil
}
