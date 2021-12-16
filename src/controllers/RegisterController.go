package controllers

import (
	"encoding/json"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type RegisterController struct {
	Db *jdft.GormAdapter `inject:"-"`
}

func NewRegisterController() *RegisterController {
	return &RegisterController{}
}

func (r *RegisterController) Register(ctx *gin.Context) int {
	type regUser struct {
		UserName  string `json:"name"`
		Password1 string `json:"password1"`
		Password2 string `json:"password2"`
	}
	var ru regUser
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(data, &ru)
	if err != nil {
		return -400
	}
	var user models.User
	r.Db.Where("username = ?", ru.UserName).First(&user)
	if user.Id > 0 {
		return -102
	}
	user.UserName = ru.UserName
	user.Password = ru.Password1
	result := r.Db.Create(&user)
	if result.Error != nil {
		return -400
	}
	return 1
}

func (r *RegisterController) Build(jdft *jdft.Jdft) {
	jdft.Handle("POST", "register", r.Register)
}
