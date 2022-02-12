package services

import (
	"context"
	"fmt"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapService struct {
	ConfigMapMap  *ConfigMapMap         `inject:"-"`
	CommonService *CommonService        `inject:"-"`
	ConfigHandler *ConfigMapHandler     `inject:"-"`
	Client        *kubernetes.Clientset `inject:"-"`
}

func NewConfigMapService() *ConfigMapService {
	return &ConfigMapService{}
}

func (c *ConfigMapService) ListNamespace(namespace string) (ret []*models.ConfigMap) {
	configList, err := c.ConfigMapMap.ListByNamespace(namespace)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range configList {
		ret = append(ret, &models.ConfigMap{
			Name:       item.Name,
			Namespace:  item.Namespace,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Data:       item.Data,
		})
	}
	return
}

func (c *ConfigMapService) CreateConfigMap(post *models.ConfigMapPost) error {
	config := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      post.Name,
			Namespace: post.Namespace,
		},
		Data: post.Data,
	}
	_, err := c.Client.CoreV1().ConfigMaps(post.Namespace).Create(context.Background(), config, metav1.CreateOptions{})
	return err
}

func (c *ConfigMapService) DeleteConfigMap(name, namespace string) error {
	return c.Client.CoreV1().ConfigMaps(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}

func (c *ConfigMapService) UpdateConfigMap(post *models.ConfigMapPost) error {
	config := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      post.Name,
			Namespace: post.Namespace,
		},
		Data: post.Data,
	}
	_, err := c.Client.CoreV1().ConfigMaps(post.Namespace).Update(context.Background(), config, metav1.UpdateOptions{})
	return err
}
