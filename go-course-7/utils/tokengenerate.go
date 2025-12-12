// 这个全部用AI写的！！！
package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 定义JWT密钥（生产环境务必放在环境变量中，不要硬编码！）
const jwtSecret = "your-secret-key-123456" // 建议替换为32位以上随机字符串

// CustomClaims 自定义JWT载荷（存储用户名、过期时间等）
type CustomClaims struct {
	Username             string `json:"username"` // 存储用户名
	jwt.RegisteredClaims        // 继承JWT标准载荷（包含过期时间等）
}

// MakeToken 生成Token
// 参数：username（用户名）、expireTime（过期时间）
// 返回：生成的Token字符串、错误
func MakeToken(username string, expireTime time.Time) (string, error) {
	// 1. 构造自定义载荷
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // Token过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), // Token签发时间
			Issuer:    "go-course-7",                  // 签发者（自定义，如项目名）
		},
	}

	// 2. 生成Token（使用HS256加密算法）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. 用密钥签名，得到最终Token字符串
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New("Token签名失败：" + err.Error())
	}

	return tokenStr, nil
}
