package model

import (
	"gorm.io/gorm"
	"time"
)

/*
	基类model 直接继承可获取如下字段
*/

type BaseModel struct {
	ID          uint           `gorm:"primarykey"` // 主键ID
	CreatedTime time.Time      // 创建时间
	UpdatedTime time.Time      // 更新时间
	DeletedTime gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
