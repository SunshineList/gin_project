package model

const (
	MALE = iota
	FEMALE
)

type User struct {
	BaseModel         // 实现了基类model
	Username  string  `json:"userName" gorm:"index;comment:用户登录名;not null"`
	Name      string  `json:"name"`
	Password  string  `json:"password"`
	Sex       *uint64 `json:"sex"`
	Phone     string  `json:"phone"  gorm:"comment:用户手机号"`
	Status    uint64  `json:"status" gorm:"default:1;comment:用户是否冻结"`
}

func (User) TableName() string {
	return "user"
}
