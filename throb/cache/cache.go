package cache

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	// 定义常量
	RedisClient *redis.Pool
)

func init() {
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("cache.maxidle", 100),
		MaxActive:   beego.AppConfig.DefaultInt("cache.maxactive", 1000),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

// 连接 redis
func Conn() redis.Conn {

	r := RedisClient.Get()
	return r
}

// 获取当前最大ID的缓存数据
func GetListValuesByMaxId(r redis.Conn, key string) []interface{} {
	values, err := redis.Values(r.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf", "LIMIT", 0, 1))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return values
}

// 根据最大最小ID区间获取缓存数据
func GetListValuesByRange(r redis.Conn, key string, minId string, maxId string) []interface{} {
	values, err := redis.Values(r.Do("ZRANGEBYSCORE", key, minId, maxId))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return values
}

// 保存列表值
func AddListValue(r redis.Conn, key string, id int64, value string) {
	_, err := r.Do("ZADD", key, id, value)

	if err != nil {
		fmt.Println(err)
	}
}

// 根据ID删除缓存数据
func DeleteListValueById(r redis.Conn, key string, id int64) {
	_, err := r.Do("ZREMRANGEBYSCORE", key, id, id)

	if err != nil {
		fmt.Println(err)
	}
}

// 保存缓存值
func SetValue(r redis.Conn, key string, value string) {
	_, err := r.Do("SET", key, value)

	if err != nil {
		fmt.Println(err)
	}
}

// 删除指定key的缓存数据
func GetValueByKey(r redis.Conn, key string) string {
	value, err := redis.String(r.Do("GET", key))

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return value
}

// 删除指定key的缓存数据
func DeleteKey(r redis.Conn, key string) {
	_, err := r.Do("DEL", key)

	if err != nil {
		fmt.Println(err)
	}
}
