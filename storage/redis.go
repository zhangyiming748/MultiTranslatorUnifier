package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// 全局 Redis 客户端变量
var RedisClient *redis.Client

// 初始化 Redis 连接
func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // 使用容器名和端口
		Password: "123456",
		DB:       0, // 使用默认数据库 (0)
	})

	// 使用 context.Background() 创建一个空的 context
	ctx := context.Background()

	// 测试连接
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("Redis 连接失败: %w", err)
	}

	fmt.Println("Redis 连接成功!")
	return nil
}
