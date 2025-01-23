package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	// Kafka broker 地址
	brokers := []string{"172.18.103.160:9092", "172.18.103.163:9092", "172.18.103.157:9092", "172.18.103.158:9092", "172.18.103.161:9092", "172.18.103.162:9092", "172.18.103.159:9092"}

	// 创建新的消费者配置
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_1 // 设置 Kafka 版本
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 使用 range 策略分配分区
	config.Net.DialTimeout = 10 * time.Second                              // 设置连接超时时间
	config.Net.SASL.Enable = false                                         // 如果未启用 SASL 认证，确保禁用

	// 创建新的消费组
	group, err := sarama.NewConsumerGroup(brokers, "your-consumer-group", config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer func() {
		if err := group.Close(); err != nil {
			log.Fatalf("Error closing consumer group: %v", err)
		}
	}()

	// 订阅的 topic
	topics := []string{"abtest_push"}

	// 创建一个 WaitGroup 来等待消费者完成
	var wg sync.WaitGroup
	wg.Add(1)

	// 创建一个 channel 来接收信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 创建上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 启动消费者
	go func() {
		defer wg.Done()
		for {
			// 消费消息
			if err := group.Consume(ctx, topics, &consumerHandler{}); err != nil {
				log.Fatalf("Error from consumer: %v", err)
			}
			// 检查是否需要退出
			if ctx.Err() != nil {
				return
			}
		}
	}()

	// 等待中断信号
	<-signals

	// 取消上下文，停止消费
	cancel()

	// 等待消费者完成
	wg.Wait()
}

// 自定义消费者处理器
type consumerHandler struct{}

// Setup 在消费者启动时调用
func (h *consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer group setup")
	return nil
}

// Cleanup 在消费者退出时调用
func (h *consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer group cleanup")
	return nil
}

// ConsumeClaim 消费消息
func (h *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Message claimed: topic = %s, partition = %d, offset = %d, key = %s, value = %s\n",
			message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		// 标记消息为已处理
		session.MarkMessage(message, "")
	}
	return nil
}
