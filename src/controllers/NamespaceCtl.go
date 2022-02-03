package controllers

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
)

type NamespaceCtl struct {
	NamespaceService *services.NamespaceService `inject:"-"`
}

func NewNamespaceCtl() *NamespaceCtl {
	return &NamespaceCtl{}
}

func (n *NamespaceCtl) GetList(c *gin.Context) (int, jdft.Json) {
	//d.DepService.ListNamespace("istio-system")
	list := n.NamespaceService.ListNamespaces()
	return 1, list
}

func (n *NamespaceCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "namespaces", n.GetList)
}
