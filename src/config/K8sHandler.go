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

func (n *K8sHandler) JdInitServiceHandlers() *services.ServiceHandler {
	return &services.ServiceHandler{}
}

func (n *K8sHandler) JdInitIngressHandlers() *services.IngressHandler {
	return &services.IngressHandler{}
}

func (n *K8sHandler) JdInitSecretHandlers() *services.SecretHandler {
	return &services.SecretHandler{}
}

func (n *K8sHandler) JdInitConfigMapHandlers() *services.ConfigMapHandler {
	return &services.ConfigMapHandler{}
}
