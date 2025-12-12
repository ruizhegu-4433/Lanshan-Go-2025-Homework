package api

import (
	"go-course-7/dao"
	"go-course-7/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	exist, _, err := dao.FindUser(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器错误",
		})
		return
	}
	if exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户已存在",
		})
		return
	}
	adderr := dao.AddUser(req.Username, req.Password)
	if adderr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

// 查找
func Find(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	exist, user, err := dao.FindUser(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器错误",
		})
		return
	}
	if exist {
		c.JSON(http.StatusBadRequest, user)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
	}
}

// 更新
func Updateinformation(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	_, err := dao.Update(req.Username, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// 删除
func DeleteUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	_, err := dao.Delete(req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
