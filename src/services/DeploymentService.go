package services

import (
	"context"
	"fmt"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentService struct {
	DepMap        *DeploymentMap        `inject:"-"`
	CommonService *CommonService        `inject:"-"`
	DepHandler    *DepHandler           `inject:"-"`
	Client        *kubernetes.Clientset `inject:"-"`
}

func NewDeploymentService() *DeploymentService {
	return &DeploymentService{}
}

func (d *DeploymentService) GetPodByDeployment(dep *v1.Deployment) ([]corev1.Pod, error) {
	podList, err := d.Client.CoreV1().Pods(dep.Namespace).List(context.Background(), metav1.ListOptions{
		LabelSelector: d.CommonService.SimpleMap2String(dep.Spec.Selector.MatchLabels),
	})
	if err != nil {
		return nil, err
	}
	return podList.Items, nil
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

// ListNamespace 接口: 列举namespace下deployment简要列表
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

// DeleteDeployment 接口，删除dep
func (d *DeploymentService) DeleteDeployment(name, namespace string) error {
	return d.Client.AppsV1().Deployments(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}

// GetDeploymentDetail 接口，获取dep详情
func (d *DeploymentService) GetDeploymentDetail(name, namespace string) (*models.DeploymentDetail, error) {
	dep, err := d.DepMap.GetDeploymentByNamespace(namespace, name)
	if err != nil {
		return nil, err
	}
	podlist, err := d.GetPodByDeployment(dep)
	if err != nil {
		return nil, err
	}
	pods := make([]*models.Pod, len(podlist))
	for i, item := range podlist {
		pods[i] = &models.Pod{
			Name:       item.Name,
			Namespace:  item.Namespace,
			Images:     d.CommonService.GetImagesByPod(item.Spec.Containers),
			NodeName:   item.Spec.NodeName,
			IP:         []string{item.Status.PodIP, item.Status.HostIP},
			Phase:      PodPhase[item.Status.Phase],
			IsReady:    d.CommonService.GetPodIsReady(&item),
			Message:    item.Status.Message,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
	}
	return &models.DeploymentDetail{
		Name:        dep.Name,
		Namespace:   dep.Namespace,
		Labels:      d.CommonService.SimpleMap2String(dep.Labels),
		Annotations: d.CommonService.SimpleMap2String(dep.Annotations),
		Replicas: [3]int32{dep.Status.Replicas,
			dep.Status.AvailableReplicas,
			dep.Status.UnavailableReplicas,
		},
		Images:     d.CommonService.GetImagesFromDeployment(*dep),
		IsComplete: d.getDeploymentIsComplete(dep),
		Message:    d.getDeploymentCondition(dep),
		CreateTime: dep.CreationTimestamp.Format("2006-01-02 15:04:05"),
		Pods:       pods,
	}, nil
}

// CreateIngress 创建 Ingress
func (d *DeploymentService) CreateDeployment(post *models.DeploymentPost) error {
	containers := make([]corev1.Container, 0)
	for _, container := range post.DeploymentTemplate.Containers {
		containers = append(containers, corev1.Container{
			Name:  container.Name,
			Image: container.Image,
		})
	}
	dep := &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        post.Name,
			Namespace:   post.Namespace,
			Labels:      post.Labels,
			Annotations: post.Annotations,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &post.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: post.DeploymentTemplate.Labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      post.DeploymentTemplate.Name,
					Namespace: post.DeploymentTemplate.Namespace,
				},
				Spec: corev1.PodSpec{
					Containers: containers,
				},
			},
		},
	}

	_, err := d.Client.AppsV1().Deployments(post.Namespace).Create(context.Background(), dep, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
