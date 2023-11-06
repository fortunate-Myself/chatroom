package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)
var pool *redis.Pool

func initPool(address string,maxIdle,maxActive int,idleTimeout time.Duration){

	pool = &redis.Pool{
		MaxIdle:maxIdle,//空闲链接数
		MaxActive:maxActive,//和数据库的最大链接，0表示没限制
		IdleTimeout:idleTimeout,//最大空闲时间
		Dial:func()(redis.Conn,error){//初始化链接的代码，链接哪个ip
			return redis.Dial("tcp",address)
		},
	}
}