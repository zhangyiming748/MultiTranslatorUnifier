package storage

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

// 全局 Redis 客户端变量
var RedisClient *redis.Client

// 初始化 Redis 连接
func InitRedis() error {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	addr := strings.Join([]string{host, port}, ":")
	password := os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		db = 0
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr, // 使用容器名和端口
		Password: password,
		DB:       db, // 使用默认数据库 (0)
	})

	// 使用 context.Background() 创建一个空的 context
	ctx := context.Background()

	// 测试连接
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("Redis 连接失败: %w", err)
	}

	fmt.Println("Redis 连接成功!")
	return nil
}

// Redis 查找函数
func RedisGet(key string) (string, error) {
	ctx := context.Background()
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("Key %s 不存在", key)
		}
		return "", fmt.Errorf("Redis Get 错误: %w", err)
	}
	return val, nil
}

// Redis 添加函数
func RedisSet(key string, value string) error {
	ctx := context.Background()
	err := RedisClient.Set(ctx, key, value, 0).Err() // 0 表示永不过期
	if err != nil {
		return fmt.Errorf("Redis Set 错误: %w", err)
	}
	return nil
}
