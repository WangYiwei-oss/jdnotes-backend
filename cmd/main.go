package main

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/src/qrcode"
	"github.com/WangYiwei-oss/jdnotes-backend/src/common"
	"github.com/WangYiwei-oss/jdnotes-backend/src/controllers"
)

func main() {
	common.GetFdNotify().Mount("D:\\test").Start()
	jdft.NewJdft().
		Beans(jdft.NewGormAdapter(), qrcode.NewQrCode()).                                   //注册依赖
		Mount("v1", controllers.NewLoginController(), controllers.NewRegisterController()). //挂载路由
		CronTask("0/3 * * * * *", func() {}).                                               //定时器函数
		Launch()
}
