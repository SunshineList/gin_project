package initialize

import (
	"fmt"
	"go_project/gin_project/gin_test/common"
	"go_project/gin_project/gin_test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func DbInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", common.DB_USER, common.DB_PASSWORD, common.DB_HOST, common.DB_PORT, common.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 这里放需要注册的模型
		model.User{},
	)
	if err != nil {
		fmt.Println("注册失败")
		os.Exit(0) // 退出程序  打包后是exe 可以执行Exit退出程序
	}
}
