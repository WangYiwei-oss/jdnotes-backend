package models

import "github.com/WangYiwei-oss/jdframe/src/jdft"

type User struct {
	jdft.User
	Intro string `gorm:"column:intro" json:"intro"`
	Icon  string `gorm:"column:icon" json:"icon"`
}
