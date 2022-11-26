package model

import (
	"errors"
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

func ToResponse(org interface{}, new interface{}) (s interface{}, err error) {
	res := map[string]interface{}{}

	if reflect.TypeOf(org).Kind() != reflect.Struct || reflect.TypeOf(new).Kind() != reflect.Struct {
		return nil, errors.New("结构错误")
	}

	n := reflect.TypeOf(new)
	o := reflect.TypeOf(org)
	oValue := reflect.ValueOf(org)
	nValue := reflect.ValueOf(new)

	for i := 0; i < n.NumField(); i++ {
		var jName = n.Field(i).Tag.Get("json")
		for k := 0; k < o.NumField(); k++ {
			if n.Field(i).Name == o.Field(k).Name {
				_, ok := n.MethodByName("Get" + n.Field(i).Name)
				if ok {
					args := []reflect.Value{reflect.ValueOf(fmt.Sprintf("%v", oValue.Field(k).Interface()))}
					res[jName] = nValue.MethodByName("Get" + n.Field(i).Name).Call(args)[0].Interface()
				} else {
					res[jName] = oValue.Field(k).Interface()
				}
				break
			} else {
				res[jName] = nil
			}
		}
	}

	return res, nil
}
