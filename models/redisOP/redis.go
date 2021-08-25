package redisOP

import (
	"github.com/gomodule/redigo/redis"
	"github.com/wonderivan/logger"
	"unioj/conf"
)

var Conn redis.Conn

func Init() {
	var err error
	Conn, err = redis.Dial("tcp", conf.GetStringConfig("redis_host"))
	if err != nil {
		logger.Error("redis初始化失败...")
		return
	}
	logger.Debug("[2] INFO redis初始化成功...")
}

func RedisSetKey(keyname string, keyval string) error {
	Conn, _ = redis.Dial("tcp", conf.GetStringConfig("redis_host"))
	_, err := Conn.Do("set", keyname, keyval)
	if err != nil {
		logger.Error(err)
		return err
	}
	return err
}

func RedisGetKey(keyname string) (string, error) {
	Conn, _ = redis.Dial("tcp", conf.GetStringConfig("redis_host"))
	reply, err := Conn.Do("get", keyname)
	str, err := redis.String(reply, err)
	if err != nil {
		logger.Error(err)
	}
	return str, err
}

func RedisSetExpireLimit(keyname string, expireSec, limit int) (bool, error) {
	Conn, _ = redis.Dial("tcp", conf.GetStringConfig("redis_host"))
	//如果键不存在设置新的
	reply, err := Conn.Do("set", keyname, 0, "NX", "EX", expireSec)
	_, err = redis.String(reply, err)
	if err != nil {
		logger.Debug(err)
	}
	exist, err := Conn.Do("ttl", keyname)
	ex, err := redis.Int(exist, err)
	current_ := 0
	if ex > 0 {
		current, err := Conn.Do("incr", keyname)
		current_, err = redis.Int(current, err)
		if err != nil {
			logger.Debug(err, current)
		}
	} else {
		Conn.Do("del", keyname)
		logger.Warn("删除异常键...")
	}
	//
	if current_ > limit {
		return true, nil
	} else {
		return false, nil
	}
}
