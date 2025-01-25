package main

import (
	"context"
	"fmt"
	"log"

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

	// 定义前缀
	prefix := "aa*"

	// 定义 Lua 脚本
	luaScript := `
		local pattern = ARGV[1]
		local count = 0
		local cursor = "0"

		repeat
			local result = redis.call("SCAN", cursor, "MATCH", pattern, "COUNT", 1000)
			cursor = result[1]
			local keys = result[2]
			count = count + #keys
		until cursor == "0"

		return count
	`

	// 执行 Lua 脚本
	val, err := rdb.Eval(ctx, luaScript, nil, prefix).Result()
	if err != nil {
		log.Fatalf("执行 Lua 脚本失败: %v", err)
	}

	// 输出匹配的键数量
	fmt.Printf("匹配前缀 '%s' 的键总数量: %d\n", prefix, val)
}
