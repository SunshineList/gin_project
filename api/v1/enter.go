package api

import "gin_project/service"

type ApiGroup struct {
	LoginApi
	UploadApi
}

var ApiGroupApp = new(ApiGroup)

var (
	userService = service.ServiceGroupApp.UserService
)
