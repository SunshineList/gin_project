package request

import (
	"gin_project/utils"
)

type UserInfo struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
}

func (u UserInfo) GetSex(sex string) string {
	return utils.ChangeSex(sex)
}

func (u UserInfo) GetStatus(status string) string {
	return utils.ChangeStatus(status)
}
