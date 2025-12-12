package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Username string `json:"username" gorm:"column:username;type:varchar(50);not null;unique"`
	Password string `json:"password" gorm:"column:password;type:varchar(100);not null"`
}
