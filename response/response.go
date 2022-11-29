package response

import (
	"errors"
	"fmt"
	"gin_project/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 1
	ERROR   = 0
)

func Result(code int, data interface{}, msg string, content *gin.Context) {
	content.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	if code == ERROR {
		return
	}
}

//成功请求返回结构(带数据的)

func OkAndData(data interface{}, msg string, content *gin.Context) {
	Result(SUCCESS, data, msg, content)
}

// 只返回成功信息

func Ok(content *gin.Context) {
	//Result(SUCCESS, map[string]interface{}{}, content)
	Result(SUCCESS, nil, "操作成功", content)
}

// 失败返回结构体

func FailAndMsg(msg string, content *gin.Context) {
	//Result(ERROR, map[string]interface{}{}, msg, content)
	Result(ERROR, nil, msg, content)
}

// unite handle error message

func HandleReturn(err error, context *gin.Context) {
	if err != nil {
		FailAndMsg(utils.Translate(err), context)
		return
	}
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

	if _, ok := o.FieldByName("BaseModel"); ok {
		// 还有BaseModel Id等字段的处理
		res["id"] = oValue.FieldByName("ID").Interface()
		res["created_time"] = oValue.FieldByName("CreatedTime").Interface()
		res["update_time"] = oValue.FieldByName("UpdatedTime").Interface()
	}

	return res, nil
}
