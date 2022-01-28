package services

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdnotes-backend/src/core"
	corev1 "k8s.io/api/core/v1"
)

type PodService struct {
	PodMap        *core.PodMap   `inject:"-"`
	CommonService *CommonService `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (p *PodService) ListNamespace(ns string) []*corev1.Pod {
	list, err := p.PodMap.ListByNamespace(ns)
	jdft.Error(err)
	return list
}
