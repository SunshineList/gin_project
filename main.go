package main

import (
	"go_project/gin_project/gin_test/common"
	globalinit "go_project/gin_project/gin_test/initialize"
)

func main() {

	g := globalinit.InitRoutes()             // 初始化路由
	common.GVA_DB = globalinit.DbInit()      // 初始化数据库
	globalinit.RegisterTables(common.GVA_DB) // 注册model

	g.Run(":9999")
}
