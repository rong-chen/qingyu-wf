package user

import (
	"qinyu-wf/global"
)

func Create(user User) error {
	return global.MySql.Create(&user).Error
}

func SearchDb(key, val string) User {
	var u User
	global.MySql.Where(key+"= ?", val).Find(&u)
	return u
}
