package user

import "qingyu-wf/global"

type User struct {
	Username string `json:"username" gorm:"username"`
	NickName string `json:"nickname" gorm:"nickname"`
	Password string `json:"-" gorm:"password"`
	Phone    string `json:"phone" gorm:"phone"`
	Email    string `json:"email" gorm:"email"`
	BirthDay string `json:"birthDay" gorm:"birthDay"`
	Gender   string `json:"gender" gorm:"gender"` // 1男，2女，3保密
	global.Model
}

type CreateParams struct {
	Username string `json:"username" gorm:"username" binding:"required"`
	Password string `json:"password" gorm:"password" binding:"required"`
}
type LoginParams struct {
	Username string `json:"username" gorm:"username" binding:"required"`
	Password string `json:"password" gorm:"password" binding:"required"`
}

func (User) TableName() string {
	return "user"
}
