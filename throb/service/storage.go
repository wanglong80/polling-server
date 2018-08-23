package service

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"happy.work/throb/cache"
)

// 获取状态数据（直接读缓存，提高QPS）
func StorageCache(name string) map[string]interface{} {
	m := make(map[string]interface{})

	r := cache.Conn()

	if r == nil {
		return nil
	}

	defer r.Close()

	key := "imx_s__" + name

	value := cache.GetValueByKey(r, key)

	if value == "" {
		return nil
	}

	err2 := json.Unmarshal([]byte(value), &m)

	if err2 != nil {
		fmt.Println(err2)
		return nil
	}

	return m
}

// 创建状态
func CreateStorage(name string, m map[string]interface{}) map[string]interface{} {
	// 把数据放入内存
	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer r.Close()

	key := "imx_s__" + name

	m2 := StorageCache(name)

	if m2 == nil {
		m2 = make(map[string]interface{})
	}

	// 与已经存在的数据合并
	for k, v := range m {
		m2[k] = v
	}

	str, _ := json.Marshal(m2)
	cache.SetValue(r, key, string(str))

	return m2
}
