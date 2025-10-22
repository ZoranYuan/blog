package core

import (
	"ZoranYuan_blog/global"
	"ZoranYuan_blog/initialize"

	"go.uber.org/zap"
)

func RunServer() {
	s := initialize.Init()
	global.Log.Info("server run success on ", zap.String("address", global.Config.System.Addr()))
	global.Log.Error(s.ListenAndServe().Error())
}
