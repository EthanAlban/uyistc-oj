package redisOP

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"unioj/conf"
	logger "unioj/logs"
)

var Conn redis.Conn

func Init() {
	var err error
	Conn, err = redis.Dial("tcp", conf.GetStringConfig("redis_host"))
	if err != nil {
		logger.LogError("redis初始化失败...")
		logger.LogError(fmt.Sprintf("err:%v", err))
		fmt.Println("redis初始化失败...")
		return
	}
	logger.LogInfo("redis初始化成功...")
	fmt.Println("[2] INFO redis初始化成功...")
}

func RedisSetKey(keyname string, keyval string) error {
	Conn, _ = redis.Dial("tcp", conf.GetStringConfig("redis_host"))
	_, err := Conn.Do("set", keyname, keyval)
	if err != nil {
		logger.LogError(fmt.Sprintf("%v", err))
		fmt.Println(err)
		return err
	}
	return err
}

func RedisGetKey(keyname string) (string, error) {
	Conn, _ = redis.Dial("tcp", conf.GetStringConfig("redis_host"))
	reply, err := Conn.Do("get", keyname)
	str, err := redis.String(reply, err)
	if err != nil {
		logger.LogError(fmt.Sprintf("%v", err))
		fmt.Println(err)
	}
	return str, err
}
