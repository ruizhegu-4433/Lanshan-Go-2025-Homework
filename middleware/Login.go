package middleware

import (
	"go-course-7/dao"
	"go-course-7/model"
	"go-course-7/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 用户登录
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求错误",
			})
			c.Abort()
			return
		}
		// 检查用户是否存在且密码是否正确
		exist, _, err := dao.FindUser(req.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "服务器错误",
			})
			c.Abort()
			return
		}
		if exist {
			userpassword, finderr := dao.FindPassword(req.Username)
			passwordhash, hasherr := utils.HashPassword(req.Password)
			if finderr != nil && hasherr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "服务器错误",
				})
				c.Abort()
				return
			} else if passwordhash != userpassword {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "密码错误",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "密码正确",
			})
			// 生成jwt token
			token, err := utils.MakeToken(req.Username, time.Now().Add(10*time.Minute))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "网络服务错误",
				})
				return
			}
			// 返回token
			c.JSON(http.StatusOK, gin.H{
				"message": "login",
				"token":   token,
			})
			c.Next()
		}
	}
}
