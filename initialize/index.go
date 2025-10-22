package initialize

import (
	"ZoranYuan_blog/global"
	"fmt"
	"net/http"
	"os"
)

func Init() *http.Server {
	// 初始化服务
	// TODO 进行错误封装，记录日志
	var err error = nil
	global.Config, err = InitConfig()
	if err != nil {
		fmt.Printf("初始化配置失败: %v\n", err)
		os.Exit(1)
		return nil
	} else {
		global.Log, err = InitLogger()
		if err != nil {
			fmt.Printf("初始化配置失败: %v\n", err)
			os.Exit(1)
		}
		global.Redis, _ = InitRedis()
		global.BlackCache, _ = IniteLocalCache()
		global.DB, err = InitDB()
		if err != nil {
			fmt.Printf("初始化配置失败: %v\n", err)
			os.Exit(1)
		}
		addr := global.Config.System.Addr()
		router := InitRouter()

		return InitServer(addr, router)
	}
}
