package config

import (
	"github.com/spf13/viper"
)


type Config struct {
	AppConfig AppConfig  `mapstructure:"app" yaml:"App"`
	DBConfig MysqlConfig `mapstructure:"mysql" yaml:"Mysql"`
	RedisConfig RedisConfig `mapstructure:"redis" yaml:"Redis"`
	RabbitMQConfig RabbitMQConfig `mapstructure:"rabbitmq" yaml:"RabbitMQ"`
}

type AppConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
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
	Vhost   string `yaml:"vhost"`
}





// 初始化配置
func InitConfig() (*Config, error) {
	var config Config
	// 用viper加载配置文件
	viper.SetConfigType("yaml")
	// viper.AutomaticEnv() // 读取环境变量
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}