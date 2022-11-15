package request

type UserInfo struct {
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Sex      *uint64 `json:"sex"`
	Phone    string  `json:"phone"`
	Status   uint64  `json:"status"`
}
