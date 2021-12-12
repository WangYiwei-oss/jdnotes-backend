package controllers

import (
	"encoding/json"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type LoginController struct {
	Db *jdft.GormAdapter `inject:"-"`
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (l *LoginController) Login(ctx *gin.Context) int {
	type User struct {
		Id       int    `gorm:"column:id" json:"id"`
		UserName string `gorm:"column:username" json:"name"`
		Password string `gorm:"column:password" json:"password"`
	}
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var user1 User
	err := json.Unmarshal(data, &user1)
	if err != nil {
		return -400
	}
	l.Db.Where("username = ? AND password = ?", user1.UserName, user1.Password).First(&user1)
	if user1.Id == 0 {
		return -100
	}
	return 1
}

func (l *LoginController) Build(jdft *jdft.Jdft) {
	jdft.Handle("POST", "login", l.Login)
}
