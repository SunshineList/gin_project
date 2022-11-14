package model

import (
	"gorm.io/gorm"
	"time"
)

/*
	基类model 直接继承可获取如下字段
*/

type BaseModel struct {
	ID          uint           `gorm:"primarykey" json:"id"`                     // 主键ID
	CreatedTime time.Time      `gorm:"autoCreateTime" json:"created_time"`       // 创建时间
	UpdatedTime time.Time      `gorm:"autoUpdateTime:milli" json:"updated_time"` // 更新时间
	DeletedTime gorm.DeletedAt `gorm:"index" json:"deleted_time"`                // 删除时间
}
