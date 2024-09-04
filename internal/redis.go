package internal

//
//import (
//	"fmt"
//	"main/utils"
//
//	"github.com/go-redis/redis"
//	_ "github.com/lib/pq"
//)
//
//type RedisInstance struct{}
//
//func NewRedisInstance() *RedisInstance {
//	return &RedisInstance{}
//}
//
//func (r *RedisInstance) ConnectRedis(p utils.RedisConfig) *redis.Client {
//	redisClient := redis.NewClient(&redis.Options{
//		Addr:         p.Address,
//		Password:     p.Password,
//		MinIdleConns: p.MinIdleConn,
//		PoolSize:     p.MaxOpenConn,
//	})
//
//	_, err := redisClient.Ping().Result()
//	if err != nil {
//		panic(err)
//	}
//	utils.ShowInfoLogs(fmt.Sprintf("Connected to redis %s", p.Address))
//	return redisClient
//}
