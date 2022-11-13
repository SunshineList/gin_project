package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

/**
全局配置文件
*/

var (
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string

	GVA_DB *gorm.DB

	//AppMode  string
	HttpPort string
)

// 初始化setting文件

func init() {
	file, err := ini.Load("common/config/config.ini") //你的ini文件所在的位置
	if err != nil {
		fmt.Println("配置文件错误无法读取", err)
	}
	LoadServer(file)
	LoadData(file)
}

// 默认配置server服务 must是默认值

func LoadServer(file *ini.File) {
	//AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

// 默认配置数据库

func LoadData(file *ini.File) {
	DB_HOST = file.Section("database").Key("DB_HOST").MustString("")
	DB_PORT = file.Section("database").Key("DB_PORT").MustInt(3306)
	DB_USER = file.Section("database").Key("DB_USER").MustString("")
	DB_PASSWORD = file.Section("database").Key("DB_PASSWORD").MustString("")
	DB_NAME = file.Section("database").Key("DB_NAME").MustString("")
}
