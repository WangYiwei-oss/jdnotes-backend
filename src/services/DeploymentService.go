package services

import (
	"fmt"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	v1 "k8s.io/api/apps/v1"
)

type DeploymentService struct {
	DepMap        *DeploymentMap `inject:"-"`
	CommonService *CommonService `inject:"-"`
	DepHandler    *DepHandler    `inject:"-"`
}

func NewDeploymentService() *DeploymentService {
	return &DeploymentService{}
}

func (d *DeploymentService) getDeploymentCondition(dep *v1.Deployment) string {
	for _, item := range dep.Status.Conditions {
		if string(item.Type) == "Available" && string(item.Status) != "True" {
			return item.Message
		}
	}
	return ""
}

//判断deployment是否就绪的方式: 副本数等于可用副本数
func (d *DeploymentService) getDeploymentIsComplete(dep *v1.Deployment) bool {
	return dep.Status.Replicas == dep.Status.AvailableReplicas
}

func (d *DeploymentService) ListNamespace(namespace string) (ret []*models.Deployment) {
	depList, err := d.DepMap.ListByNamespace(namespace)
	if err != nil {
		fmt.Println(err)
	}
	//jdft.Error(err)
	for _, item := range depList {
		ret = append(ret, &models.Deployment{
			Name:      item.Name,
			Namespace: item.Namespace,
			Replicas: [3]int32{item.Status.Replicas,
				item.Status.AvailableReplicas,
				item.Status.UnavailableReplicas,
			},
			Images:     d.CommonService.GetImagesFromDeployment(*item),
			IsComplete: d.getDeploymentIsComplete(item),
			Message:    d.getDeploymentCondition(item),
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return
}
