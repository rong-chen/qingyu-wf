package utils

import "golang.org/x/crypto/bcrypt"

func Vaild(pwd, hashPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}

func Encryption(pwd string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(password), nil
}
