package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"outdoor_rental/config"
	"time"
)

const REDIS_UTIL_ERR_PREFIX = "utils/redis.go ->"

var (
	ctx = context.Background() //创建一个空的根上下文
	rdb *redis.Client
)

// 对 Redis 库的操作二次封装, 统一处理错误
var Redis = new(_redis)

type _redis struct{}

// 初始化 redis 连接
func InitRedis() *redis.Client {
	redisCfg := config.Cfg.Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	// 测试连接状况
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic("Reids 连接失败: ", err)
	}
	log.Println("Redis 连接成功 ")
	return rdb
}

//GetVal 获取值
func (*_redis) GetVal(key string) string {
	return rdb.Get(ctx, key).Val()
}

//Set redis设置 key value 过期时间
func (*_redis) Set(key string, value interface{}, expiration time.Duration) {
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"Set: ", zap.Error(err))
		panic(err)
	}
}

// Del redis 删除值
func (*_redis) Del(key string) {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"Del: ", zap.Error(err))
		panic(err)
	}
}

// Keys redis 获取根据匹配项获取键名列表
func (*_redis) Keys(pattern string) []string {
	return rdb.Keys(ctx, pattern).Val()
}
