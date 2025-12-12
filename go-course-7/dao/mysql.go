package dao

import (
	"log" // 日志包，用于输出错误/成功信息

	"go-course-7/model"
	"go-course-7/utils"

	"gorm.io/driver/mysql" // GORM的MySQL驱动，负责和MySQL通信
	"gorm.io/gorm"         // GORM核心包，提供ORM操作能力
)

var DB *gorm.DB

// 链接数据库
func Connectmysql() {
	dsn := "root:wxhdyh4433@tcp(127.0.0.1:3306)/datas?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败: ", err)
		return
	}
	DB = db
	log.Println("链接数据库成功")
}

// 迁移数据库
func CreateTable() {
	if DB == nil {
		log.Fatal("未链接数据库")
		return
	}
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("自动迁移失败: ", err)
		return
	}
	log.Println("自动迁移成功")
}

// 查找用户
func FindUser(username string) (bool, model.User, error) {
	if DB == nil {
		log.Println("未链接数据库")
		return false, model.User{}, DB.Error
	}
	var user model.User
	result := DB.Where("username = ? ", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("没有查询到指定项")
			return false, model.User{}, DB.Error
		}
		log.Println("出现错误", result.Error)
		return false, model.User{}, DB.Error
	}
	return true, user, nil
}

// 增添用户
func AddUser(username, password string) error {
	if DB == nil {
		log.Println("未链接数据库")
		return DB.Error
	}
	passwordhash, err := utils.HashPassword(password)
	if err != nil {
		log.Println("哈希失败,新增用户失败")
		return err
	}
	newUser := model.User{
		Username: username,     // 用户名
		Password: passwordhash, // 密码
	}
	result := DB.Create(&newUser)
	if result.Error != nil {
		log.Println("新增用户失败", result.Error)
		return DB.Error
	}
	return nil
}

// 获得哈希后的密码
func FindPassword(username string) (string, error) {
	if DB == nil {
		log.Println("未链接数据库")
		return "", DB.Error
	}
	var user model.User
	result := DB.Where("username = ? ", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("没有查询到指定项")
			return "", result.Error
		}
		log.Println("出现错误", result.Error)
		return "", result.Error
	}
	return user.Password, nil
}

// 更新
func Update(username string, param model.User) (bool, error) {
	if DB == nil {
		log.Println("未链接数据库")
		return false, DB.Error
	}
	var user model.User
	result := DB.Where("username = ? ", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("没有查询到指定项")
			return false, DB.Error
		}
		log.Println("出现错误", result.Error)
		return false, DB.Error
	}
	updateResult := DB.Model(&user).Updates(param)
	if updateResult.Error != nil {
		log.Printf("更新用户[%s]信息失败：%v", username, updateResult.Error)
		return false, updateResult.Error
	}
	return true, nil
}

// 删除用户
func Delete(username string) (bool, error) {
	if DB == nil {
		log.Println("未链接数据库")
		return false, DB.Error
	}
	var user model.User
	result := DB.Where("username = ? ", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("没有查询到指定项")
			return false, DB.Error
		}
		log.Println("出现错误", result.Error)
		return false, DB.Error
	}
	deleteResult := DB.Unscoped().Delete(&user)
	if deleteResult.Error != nil {
		log.Printf("硬删除用户[%s]失败：执行物理删除操作时出错 - %v", username, deleteResult.Error)
		return false, deleteResult.Error
	}
	return true, nil
}
