package config

import (
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"k8s.io/client-go/kubernetes"
)

type K8sMap struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewK8sMap() *K8sMap {
	return &K8sMap{}
}

func (k *K8sMap) JdInitDepMap() *services.DeploymentMap {
	return &services.DeploymentMap{}
}

func (k *K8sMap) JdInitPodMap() *services.PodMap {
	return &services.PodMap{}
}

func (k *K8sMap) JdInitNamespaceMap() *services.NamespaceMap {
	return &services.NamespaceMap{}
}

func (k *K8sMap) JdInitServiceMap() *services.ServiceMap {
	return &services.ServiceMap{}
}

func (k *K8sMap) JdInitIngressMap() *services.IngressMap {
	return &services.IngressMap{}
}

func (k *K8sMap) JdInitSecretMap() *services.SecretMap {
	return &services.SecretMap{}
}

func (k *K8sMap) JdInitConfigMapMap() *services.ConfigMapMap {
	return &services.ConfigMapMap{}
}
