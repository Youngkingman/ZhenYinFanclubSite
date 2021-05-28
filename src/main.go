package main

import algorithm "basic/models/algorithm/searchmethod"

func main() {
	//初始化数据库
	// mysql.InitSQL(config.Current.MySqlConfig)
	// redis.SetConfig(config.Current.RedisConfig)

	// //开启路由监听
	// r := routers.SetupRouter()
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

	//tract调试
	// f, _ := os.Create("trace.out")
	// defer f.Close()
	// trace.Start(f)
	algorithm.Compare(500, 500, 0.2, 10, 11, [2]int{1, 1}, [2]int{495, 495}, 0)
	// trace.Stop()
}
