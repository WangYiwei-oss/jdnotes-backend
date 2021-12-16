package models

type User struct {
	Id       int    `gorm:"column:id" json:"id"`
	UserName string `gorm:"column:username" json:"name"`
	Password string `gorm:"column:password" json:"password"`
}
