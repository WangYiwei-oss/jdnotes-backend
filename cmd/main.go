package main

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdnotes-backend/src/common"
	"github.com/WangYiwei-oss/jdnotes-backend/src/config"
	"github.com/WangYiwei-oss/jdnotes-backend/src/controllers"
	"github.com/WangYiwei-oss/jdnotes-backend/src/middlewares"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
)

func migration() {
	//jdft.NewGormAdapter().AutoMigrate(&models.User{})
}

func main() {
	//migration()	//调试用
	common.GetFdNotify().Mount("D:\\test").Start()
	jdft.NewJdft().
		DefaultBean().
		Attach(middlewares.CrossMiddleWare()).
		Beans(
			services.NewDeploymentMap(),
			services.NewPodMap(),
			services.NewNamespaceMap(),
			config.NewK8sHandler(),
			config.NewK8sConfig(),
			config.NewK8sMap(),
			config.NewServiceConfig()).
		Mount("v1",
			controllers.NewLoginController(),
			controllers.NewRegisterController(),
			controllers.NewDeploymentCtl(),
			controllers.NewPodCtl(),
			controllers.NewNamespaceCtl(),
		).                                    //挂载路由
		CronTask("0/3 * * * * *", func() {}). //定时器函数
		Launch()
}
