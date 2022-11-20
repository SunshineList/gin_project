package model

import (
	"fmt"
	"reflect"
)

const (
	MALE = iota
	FEMALE
)

const ACTIVE = 1

type User struct {
	BaseModel         // 实现了基类model
	Username  string  `json:"username" gorm:"index;comment:用户登录名;not null;unique"`
	Name      string  `json:"name"`
	Password  string  `json:"password"`
	Sex       *uint64 `json:"sex"`
	Phone     string  `json:"phone"  gorm:"comment:用户手机号"`
	Status    uint64  `json:"status" gorm:"default:1;comment:用户是否冻结"`
}

func (User) TableName() string {
	return "user"
}

func ToResponse(org interface{}, new interface{}) (s interface{}) {
	var res = map[string]interface{}{}

	if reflect.TypeOf(org).Kind() != reflect.Struct || reflect.TypeOf(new).Kind() != reflect.Struct {
		fmt.Println("错误")
		return
	}

	n := reflect.TypeOf(new)
	o := reflect.TypeOf(org)
	oValue := reflect.ValueOf(org)

	for i := 0; i < n.NumField(); i++ {
		var jName = n.Field(i).Tag.Get("json")
		for k := 0; k < o.NumField(); k++ {
			if n.Field(i).Name == o.Field(k).Name {
				res[jName] = oValue.Field(k)
				break
			} else {
				res[jName] = nil
			}
		}
	}

	fmt.Printf("得到的值 %v", res)

	return res
}
