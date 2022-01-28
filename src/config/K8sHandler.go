package config

import "github.com/WangYiwei-oss/jdnotes-backend/src/core"

type K8sHandler struct {
}

func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}

func (n *K8sHandler) JdInitDepHandlers() *core.DepHandler {
	return &core.DepHandler{}
}

func (n *K8sHandler) JdInitPodHandlers() *core.PodHandler {
	return &core.PodHandler{}
}
