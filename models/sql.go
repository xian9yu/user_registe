package models

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"xorm.io/xorm"
)

var DB *xorm.Engine

//初始化sql
func InitSQL() {
	var err error
	DB, err = xorm.NewEngine("mysql", "register:register@tcp(127.0.0.1:3306)/register?charset=utf8")
	if err != nil {
		log.Fatal("Disconnected database:", err)
	}

	if err = DB.Sync2(new(User)); err != nil {
		//if err != nil {
		log.Fatal("Failed to sync tables :", err)
		return
	}
	//DB.Close()
	//最大连接池空闲数(连接池空闲数 = ((核心数 * 2) + 有效磁盘数))
	DB.SetMaxIdleConns(3)
	//最大连接数
	DB.SetMaxOpenConns(100)

}

var (
	redisAddr = "127.0.0.1:6379"
	redisPwd  = "123456"
)

//初始化Redis （使用 redigo连接池）
func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     1024,               //Linux下使用 ulimit -n 命令查看，Linux默认最大值为 1024
		MaxActive:   256,                //最大空闲连接数
		Wait:        true,               //超过最大连接，报错or等待
		IdleTimeout: time.Second * 1000, //空闲连接超时时间
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", redisAddr, redis.DialPassword(redisPwd)) },
	}
}
