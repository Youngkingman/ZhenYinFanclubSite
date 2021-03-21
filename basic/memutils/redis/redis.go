package redis

import (
	"fmt"
	"log"
	"sync"
	"time"

	redis "github.com/go-redis/redis"
)

//Config redis的配置文件
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Auth     string `json:"auth"`
	Db       int    `json:"db"`
	PoolSize int    `json:"poolSize"`
}

var config Config

//SetConfig 设置配置文件
func SetConfig(conf Config) {
	config = conf
}

var redisOnce sync.Once
var client *redis.Client

//ErrKeyNotExists not exists
var ErrKeyNotExists = redis.Nil

//GetRedisClient get the client of redis
func getRedisClient() *redis.Client {
	redisOnce.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:       fmt.Sprintf("%s:%d", config.Host, config.Port),
			Password:   config.Auth,
			DB:         config.Db,
			MaxRetries: 2,
			PoolSize:   config.PoolSize,
		})
	})
	return client
}

//Get get redis value
func Get(key string) (string, error) {
	return getRedisClient().Get(key).Result()
}

//Set set redis key value
func Set(key string, val string, expiration time.Duration) error {
	err := getRedisClient().Set(key, val, expiration).Err()
	if err != nil {
		log.Printf("redis set key: %s val: %s fail: %s", key, val, err)
	}
	return err
}

//Del delete the key
func Del(key string) error {
	return getRedisClient().Del(key).Err()
}

//TTL change the Ttl
func TTL(key string) (time.Duration, error) {
	r := getRedisClient().TTL(key)
	return r.Val(), r.Err()
}

//Client return the raw redis client
func Client() *redis.Client {
	return getRedisClient()
}

//RPush RPush
func RPush(key string, value string) error {

	return getRedisClient().RPush(key, value).Err()
}

//LPush LPush
func LPush(key string, value string) error {
	return getRedisClient().LPush(key, value).Err()
}

//RPop RPop
func RPop(key string) (string, error) {
	r := getRedisClient().RPop(key)
	return r.Val(), r.Err()
}

//LPop LPop
func LPop(key string) (string, error) {
	r := getRedisClient().LPop(key)
	return r.Val(), r.Err()
}

//LLen LLen
func LLen(key string) (int64, error) {
	r := getRedisClient().LLen(key)
	return r.Val(), r.Err()
}

//LRem LRem
func LRem(key string, count int64, value string) error {
	return getRedisClient().LRem(key, count, value).Err()
}

//Expire set redis key expire
func Expire(key string, expiration time.Duration) error {
	err := getRedisClient().Expire(key, expiration).Err()
	return err
}
