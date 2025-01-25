package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 密码，如果没有设置密码则为空
		DB:       0,                // 使用默认 DB
	})

	// 使用 context 来控制超时等
	ctx := context.Background()

	// 检查连接是否成功
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis: %v", err)
	}

	// // 设置一个键值对
	// err = rdb.Set(ctx, "key", "value222", 0).Err()
	// if err != nil {
	// 	log.Fatalf("设置键值对失败: %v", err)
	// }

	// // 读取键值对
	// val, err := rdb.Get(ctx, "key").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key 不存在")
	// } else if err != nil {
	// 	log.Fatalf("读取键值对失败: %v", err)
	// } else {
	// 	fmt.Println("key 的值为:", val)
	// }

	// 定义前缀
	readReids(ctx, err, rdb)
	// writeReids(ctx, err, rdb)

}

func writeReids(ctx context.Context, err error, rdb *redis.Client) {
	// 定义批量设置的键值对数量
	numKeys := 1000

	// 创建 Pipeline
	pipe := rdb.Pipeline()

	// 开始时间（用于计算性能）
	start := time.Now()

	// 批量设置键值对
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("aaaakey%d", i)
		value := fmt.Sprintf("value%d", i)
		pipe.Set(ctx, key, value, 0) // 0 表示不过期
	}

	// 执行 Pipeline
	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Fatalf("批量设置失败: %v", err)
	}

	// 计算耗时
	elapsed := time.Since(start)
	fmt.Printf("批量设置 %d 个键值对成功，耗时: %v\n", numKeys, elapsed)
}

func readReids(ctx context.Context, err error, rdb *redis.Client) {
	prefix := "aa*"

	// 使用 SCAN 遍历匹配前缀的键
	var count int
	var cursor uint64
	var keys []string
	for {
		// 执行 SCAN 命令

		keys, cursor, err = rdb.Scan(ctx, cursor, prefix, 100).Result()
		if err != nil {
			log.Fatalf("SCAN 失败: %v", err)
		}

		// 增加计数器
		count += len(keys)

		// 如果 cursor 为 0，表示遍历完成
		if cursor == 0 {
			break
		}
	}

	// 输出匹配的键数量
	fmt.Printf("匹配前缀 '%s' 的键总数量: %d\n", prefix, count)

}
