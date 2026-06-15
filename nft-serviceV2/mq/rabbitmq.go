package mq

import (
	"fmt"
	"log"
	"nft-service/config"

	"github.com/streadway/amqp"
)

// 初始化RabbitMQ连接
func InitRabbitMQ() error {
	// 连接到RabbitMQ服务器
	// 创建一个连接
	url := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
	config.GlobalConfig.RabbitMQConfig.Username, 
	config.GlobalConfig.RabbitMQConfig.Password,
	config.GlobalConfig.RabbitMQConfig.Host,
	config.GlobalConfig.RabbitMQConfig.Port,
	config.GlobalConfig.RabbitMQConfig.Vhost)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("mq连接失败: %s", err)
		log.Fatal("mq初始化失败")
		return err
	}
	defer conn.Close()

	// 创建一个通道
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// 声明一个交换器
	err = ch.ExchangeDeclare(
		"nft_exchange", // 交换器名称
		"direct",      // 交换器类型
		true,          // 是否持久化
		false,         // 是否自动删除
		false,         // 是否内部交换器	// 是否匿名
		false,         // 是否声明队列
		nil,          // 交换器属性
	)
	if err != nil {
		panic(err)
	}

	// 声明一个队列
	q, err := ch.QueueDeclare(
		"nft_queue", // 队列名称
		true,          // 是否持久化
		false,         // 是否自动删除
		false,         // 是否内部队列	// 是否匿名
		false,         // 是否声明交换器
		nil,          // 队列属性
	)
	if err != nil {
		panic(err)
	}

	// 绑定交换器和队列
	err = ch.QueueBind(
		q.Name,       // 队列名称
		"nft_routing_key", // 路由键
		"nft_exchange", // 交换器名称	// 绑定键
		false,         // 是否自动删除
		nil,          // 绑定属性
	)
	if err != nil {
		panic(err)
	}
	return nil
}