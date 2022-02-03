package config

import (
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
)

type K8sHandler struct {
}

func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}

func (n *K8sHandler) JdInitDepHandlers() *services.DepHandler {
	return &services.DepHandler{}
}

func (n *K8sHandler) JdInitPodHandlers() *services.PodHandler {
	return &services.PodHandler{}
}

func (n *K8sHandler) JdInitNamespaceHandlers() *services.NamespaceHandler {
	return &services.NamespaceHandler{}
}
