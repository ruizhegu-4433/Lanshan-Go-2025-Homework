package main

import (
	"go-course-7/api"
	"go-course-7/dao"
	"go-course-7/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(middleware.Logger())

	dao.Connectmysql()
	dao.CreateTable()

	publicGroup := r.Group("/api/v1")
	{
		publicGroup.POST("/register", api.Register)
	}

	protectedGroup := r.Group("/api/v1/user")
	protectedGroup.Use(middleware.Login())
	{
		protectedGroup.POST("/find", api.Find)
		protectedGroup.POST("/update", api.Updateinformation)
		protectedGroup.POST("/delete", api.DeleteUser)
	}

	err := r.Run(":8080")
	if err != nil {
		panic("服务启动失败: " + err.Error())
	}
}
