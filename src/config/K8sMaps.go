package config

import (
	"github.com/WangYiwei-oss/jdnotes-backend/src/core"
	"k8s.io/client-go/kubernetes"
)

type K8sMap struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewK8sMap() *K8sMap {
	return &K8sMap{}
}

func (k *K8sMap) JdInitDepMap() *core.DeploymentMap {
	return &core.DeploymentMap{}
}

func (k *K8sMap) JdInitPodMap() *core.PodMap {
	return &core.PodMap{}
}
