package models

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type AuthRedis struct {
	Key  string
	Code string
}

//set key:value
func Set(key, code string) bool {
	c := RedisPool().Get()
	if c == nil {
		log.Fatal("redis connect Error faction back nil")
		return false
	}
	//defer c.Close()
	_, err := c.Do("SET", key, code, "EX", "6000") //SET 后接传入参数， EX 后接销毁时间，180s后销毁数据
	if err != nil {
		log.Fatal("Redis Add data Error: ", err)
		return false
	}
	return true
}

//Get value
func Get(email string) interface{} {
	c := RedisPool().Get()
	if c == nil {
		log.Fatal("redis connect Error faction back nil")
	}
	v, err := redis.String(c.Do("GET", email))
	if err != nil {
		log.Fatal("Redis Get data Error: ", err)
	}
	return v
}

//Check is Exist
func Exists(key string) interface{} {
	c := RedisPool().Get()
	if c == nil {
		log.Fatal("redis connect Error faction back nil")
	}
	e, err := c.Do("EXISTS", key)
	if err != nil {
		log.Fatal("Redis EXIST Error: ", err)
	}
	return e
}
