package controllers

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
)

type PodCtl struct {
	PodService *services.PodService `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (p *PodCtl) GetList(c *gin.Context) (int, jdft.Json) {
	list := p.PodService.ListNamespace("istio-system")
	return 1, list
}

func (p *PodCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "pods", p.GetList)
}
