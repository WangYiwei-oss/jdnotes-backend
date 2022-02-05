package services

import (
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	corev1 "k8s.io/api/core/v1"
	"log"
)

var PodPhase map[corev1.PodPhase]string

func init() {
	PodPhase = map[corev1.PodPhase]string{
		corev1.PodPending:   "Pending",
		corev1.PodRunning:   "Running",
		corev1.PodSucceeded: "Succeeded",
		corev1.PodFailed:    "Failed",
		corev1.PodUnknown:   "Unknown",
	}
}

type PodService struct {
	PodMap        *PodMap        `inject:"-"`
	CommonService *CommonService `inject:"-"`
	PodHandler    *PodHandler    `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (p *PodService) ListNamespace(ns string) (ret []*models.Pod) {
	podlist, err := p.PodMap.ListByNamespace(ns)
	if err != nil {
		log.Println("PodService:", err)
		return
	}
	for _, item := range podlist {
		ret = append(ret, &models.Pod{
			Name:       item.Name,
			Namespace:  item.Namespace,
			Images:     p.CommonService.GetImagesByPod(item.Spec.Containers),
			NodeName:   item.Spec.NodeName,
			IP:         []string{item.Status.PodIP, item.Status.HostIP},
			Phase:      PodPhase[item.Status.Phase],
			IsReady:    p.CommonService.GetPodIsReady(item),
			Message:    item.Status.Message,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return
}
