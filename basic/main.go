package main

import (
	"basic/config"
	"basic/memutils/mysql"
	"basic/memutils/redis"
	"basic/routers"
	"net/http"
	"time"
)

func main() {
	//初始化数据库
	mysql.InitSQL(config.Current.MySqlConfig)
	redis.SetConfig(config.Current.RedisConfig)

	//开启路由监听
	r := routers.SetupRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
