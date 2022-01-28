package services

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (c *CommonService) GetImages(dep v1.Deployment) string {
	return c.GetImagesByPod(dep.Spec.Template.Spec.Containers)
}

func (c *CommonService) GetImagesByPod(containers []corev1.Container) string {
	if len(containers) < 1 {
		return ""
	}
	images := ""
	for _, container := range containers {
		images = images + container.Image + "/"
	}
	return images[:len(images)]
}
