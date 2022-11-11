package api

import "go_project/gin_project/gin_test/service"

type ApiGroup struct {
	LoginApi
	UploadApi
}

var ApiGroupApp = new(ApiGroup)

var (
	userService = service.ServiceGroupApp.UserService
)
