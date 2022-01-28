package controllers

import (
	"encoding/json"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type LoginController struct {
	Db *jdft.Gorm `inject:"-"`
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (l *LoginController) Login(ctx *gin.Context) int {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var user1 models.User
	err := json.Unmarshal(data, &user1)
	if err != nil {
		return -400
	}
	l.Db.Where("username = ? AND password = ?", user1.UserName, user1.Password).First(&user1)
	if user1.UserId == 0 {
		return -100
	}
	return 1
}

func (l *LoginController) Build(jdft *jdft.Jdft) {
	jdft.Handle("POST", "login", l.Login)
}
