package common

import "gorm.io/gorm"

var (
	DB_USER     = "root"
	DB_PASSWORD = "19981008"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = 3306
	DB_NAME     = "gin"

	GVA_DB *gorm.DB
)
