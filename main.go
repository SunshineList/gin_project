package main

import (
	common "gin_project/common/config"
	globalinit "gin_project/initialize"
)

func main() {

	g := globalinit.InitRoutes()             // 初始化路由
	common.GVA_DB = globalinit.DbInit()      // 初始化数据库
	globalinit.RegisterTables(common.GVA_DB) // 注册model

	g.Run(common.HttpPort)
}
