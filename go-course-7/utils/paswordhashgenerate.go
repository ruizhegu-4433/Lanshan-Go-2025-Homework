package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// 生成密码的哈希值
func HashPassword(rawPassword string) (string, error) {
	if rawPassword == "" {
		return "", errors.New("密码不能为空")
	}
	hashBytes, err := bcrypt.GenerateFromPassword(
		[]byte(rawPassword), // 转字节数组
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", errors.New("密码哈希生成失败：" + err.Error())
	}

	return string(hashBytes), nil
}
