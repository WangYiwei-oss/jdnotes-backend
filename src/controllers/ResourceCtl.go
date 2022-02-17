package controllers

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

type ResourceCtl struct {
	K8sClient       *kubernetes.Clientset     `inject:"-"`
	ResourceService *services.ResourceService `inject:"-"`
}

func NewResourceCtl() *ResourceCtl {
	return &ResourceCtl{}
}

func (r *ResourceCtl) GetList(c *gin.Context) (int, jdft.Json) {
	res, err := r.ResourceService.ListResources()
	if err != nil {
		return -400, err.Error()
	}
	return 1, res
}

func (r *ResourceCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "resources", r.GetList)
}
