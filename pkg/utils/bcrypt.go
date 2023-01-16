package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	PassWordCost = 12 //密码加密难度
)

// EncipherPassword 加密密码
func EncipherPassword(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return "", err
	}
	return string(pwd), nil
}

// VerifyPassword 检验密码
func VerifyPassword(oldPwd, nowPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(oldPwd), []byte(nowPwd))
	return err == nil
}
