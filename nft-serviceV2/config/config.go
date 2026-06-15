package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig      AppConfig      `mapstructure:"app" yaml:"App"`
	MysqlConfig    MysqlConfig    `mapstructure:"mysql" yaml:"Mysql"`
	RedisConfig    RedisConfig    `mapstructure:"redis" yaml:"Redis"`
	RabbitMQConfig RabbitMQConfig `mapstructure:"rabbitmq" yaml:"RabbitMQ"`
	Chain          ChainConfig    `mapstructure:"chain" yaml:"Chain"`
}

type AppConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type RabbitMQConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

type ChainConfig struct {
	ChainId        string `yaml:"chain_id"`
	RpcUrl         string `yaml:"rpc_url"`
	StartBlock     uint64 `yaml:"start_block"`
	ScanInterval   uint64 `yaml:"scan_interval"`
	BatchSize      uint64 `yaml:"batch_size"`
	TargetContract string `yaml:"target_contract"`
}

var GlobalConfig Config

// 初始化配置
func InitConfig() error {

	// 用viper加载配置文件
	viper.SetConfigType("yaml")
	// viper.AutomaticEnv() // 读取环境变量
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		return err
	}
	return nil
}
