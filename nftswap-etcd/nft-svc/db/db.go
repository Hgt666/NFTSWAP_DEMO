package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

// MysqlConfig MySQL配置
type MysqlConfig struct {
	Dsn string `yaml:"dsn"`
}

// InitMySQL 初始化数据库
func InitMySQL(cfg MysqlConfig) error {
	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 关闭日志
	})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// 连接池配置
	sqlDB.SetMaxOpenConns(80)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	DB = db
	return nil
}