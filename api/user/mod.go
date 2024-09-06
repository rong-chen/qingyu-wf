package user

type User struct {
	Username string `json:"username" gorm:"username"`
	NickName string `json:"nickname" gorm:"nickname"`
	Password string `json:"password" gorm:"password"`
	Phone    string `json:"phone" gorm:"phone"`
	Email    string `json:"email" gorm:"email"`
	BirthDay string `json:"birthDay" gorm:"birthDay"`
	Gender   string `json:"gender" gorm:"gender"` // 1男，2女，3保密
}

func (User) TableName() string {
	return "user"
}

func (User) AutoMigrateFunc() {}
