package redis

import (
	"fmt"
	"github.com/xiaopengkuang/gin-web/config"
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *GoRedisClient

type GoRedisClient struct {
	client *redis.Client
}

// 获取数据
func (r *GoRedisClient) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

// 设置key value
func (r *GoRedisClient) Set(key, value string, duration time.Duration) {
	r.client.Set(key, value, duration)
}

func (r *GoRedisClient) initRedisClient() error {
	redisConfig := config.AppConfig.GetRedisConfig()
	ok := redisConfig.CheckInfo()
	if !ok {
		return fmt.Errorf("redis config info missing")
	}

	r.client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})

	pong, err := r.client.Ping().Result()
	if err != nil {
		return err
	}

	fmt.Println(pong)
	return nil
}

func InitRedis() error {
	RedisClient = &GoRedisClient{}
	return RedisClient.initRedisClient()
}
