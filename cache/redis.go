package cache

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"

	. "go-integration/conf"
)

var RedisClient *redis.Pool

func init() {
	log.Println("初始化redis中")
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,                 //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			conn, err := redis.Dial(Redis["type"], Redis["address"])
			if err != nil {
				return nil, err
			}
			if Redis["auth"] != "" {
				if _, err := conn.Do("AUTH", Redis["auth"]); err != nil {
					_ = conn.Close()
					return nil, err
				}
			}
			return conn, nil
		},
	}
	log.Println("初始化redis完毕")
}
