package controllers

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type LoginController struct {
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (l *LoginController)Login(ctx *gin.Context)int{
	data,_ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println(string(data))
	return 1
}

func (l *LoginController) Build(jdft *jdft.Jdft) {
	jdft.Handle("POST", "login", l.Login)
}
