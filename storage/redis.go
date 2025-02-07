package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// 全局 Redis 客户端变量
var redisClient *redis.Client

// 初始化 Redis 连接
func InitRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // 使用容器名和端口
		Password: "123456",
		DB:       0, // 使用默认数据库 (0)
	})

	// 使用 context.Background() 创建一个空的 context
	ctx := context.Background()

	// 测试连接
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("Redis 连接失败: %w", err)
	}

	fmt.Println("Redis 连接成功!")
	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func InsertTranslationToRedis(src, dst, from string) error {
	ctx := context.Background()
	data := map[string]string{
		"dst":  dst,
		"from": from,
	}
	// 使用 HMSet 命令插入 Hash 数据
	err := GetRedis().HMSet(ctx, src, data).Err()
	if err != nil {
		log.Println("插入 Redis Hash 失败", err)
		return err
	}
	log.Println("插入 Redis Hash 成功")
	return nil
}

// GetTranslationFromRedis 方法，从 Redis Hash 中获取翻译数据 (已存在，无需修改)
func GetTranslationFromRedis(src string) (dst string, from string, err error) {
	ctx := context.Background()
	// 使用 HMGet 命令获取指定字段的值
	result, err := GetRedis().HMGet(ctx, src, "dst", "from").Result()
	if err != nil {
		log.Println("读取 Redis Hash 失败", err)
		return "", "", err
	}

	// 处理结果
	if len(result) == 2 {
		dst, _ = result[0].(string) // 类型断言
		from, _ = result[1].(string)
	}

	return dst, from, nil
}

// SearchTranslationFromRedis 方法, 根据src查找dst, 同时打印from (新增)
func SearchTranslationFromRedis(src string) (dst string, err error) {
	ctx := context.Background()
	// 使用 HMGet 命令获取指定字段的值
	result, err := GetRedis().HMGet(ctx, src, "dst", "from").Result()
	if err != nil {
		log.Println("读取 Redis Hash 失败", err)
		return "", err
	}

	// 处理结果
	if len(result) == 2 {
		dst, _ = result[0].(string) // 类型断言
		from, _ := result[1].(string)
		fmt.Printf("译文: %s, 来源: %s\n", dst, from) // 打印译文和来源
	}

	return dst, nil
}
