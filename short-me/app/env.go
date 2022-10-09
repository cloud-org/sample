package main

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	S Storage
}

// 读取环境变量 初始化 redis client
func getEnv() *Env {
	var (
		addr   string
		passwd string
		db     int
		err    error
		r      *RedisCli
	)
	addr = os.Getenv("APP_REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}

	passwd = os.Getenv("APP_REDIS_PASSWORD")
	if passwd == "" {
		passwd = ""
	}
	if os.Getenv("APP_REDIS_DB") == "" {
		db = 0
	} else {
		if db, err = strconv.Atoi(os.Getenv("APP_REDIS_DB")); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("connect to redis ", addr, passwd, db)

	r = NewRedisCli(addr, passwd, db)

	return &Env{S: r}
}
