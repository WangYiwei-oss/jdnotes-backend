package services

import (
	"context"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

type RoleService struct {
	RoleMap       *RoleMap              `inject:"-"`
	CommonService *CommonService        `inject:"-"`
	RoleHandler   *RoleHandler          `inject:"-"`
	Client        *kubernetes.Clientset `inject:"-"`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

// ListNamespace 返回ns命名空间下所有的secret列表
func (r *RoleService) ListNamespace(ns string) (ret []*models.Role) {
	list, err := r.RoleMap.ListByNamespace(ns)
	if err != nil {
		log.Println("RoleService:", err)
		return
	}
	for _, item := range list {
		ret = append(ret, &models.Role{
			Name:       item.Name,
			Namespace:  item.Namespace,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return ret
}

func (r *RoleService) DelSecret(name, namespace string) error {
	return r.Client.RbacV1().Roles(namespace).Delete(context.Background(), name, v1.DeleteOptions{})
}
