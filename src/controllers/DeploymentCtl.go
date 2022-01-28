package controllers

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

type DeploymentCtl struct {
	K8sClient  *kubernetes.Clientset       `inject:"-"`
	DepService *services.DeploymentService `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (d *DeploymentCtl) GetList(c *gin.Context) (int, jdft.Json) {
	//d.DepService.ListNamespace("istio-system")
	list := d.DepService.ListNamespace("istio-system")
	return 1, list
}

func (d *DeploymentCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "deployments", d.GetList)
}
