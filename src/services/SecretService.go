package services

import (
	"context"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

var SECRET_TYPE_MAP map[string]string

func init() {
	SECRET_TYPE_MAP = map[string]string{
		"Opaque":                              "自定义类型",
		"kubernetes.io/service-account-token": "服务账号令牌",
		"kubernetes.io/dockercfg":             "docker配置",
		"kubernetes.io/dockerconfigjson":      "docker配置(JSON)",
		"kubernetes.io/basic-auth":            "Basic认证凭据",
		"kubernetes.io/ssh-auth":              "SSH凭据",
		"kubernetes.io/tls":                   "TLS凭据",
		"bootstrap.kubernetes.io/token":       "启动引导令牌数据",
	}
}

func GetSecretType(secret *corev1.Secret) string {
	if t, ok := SECRET_TYPE_MAP[string(secret.Type)]; ok {
		return t
	} else {
		return string(secret.Type)
	}
}

type SecretService struct {
	SecretMap     *SecretMap            `inject:"-"`
	CommonService *CommonService        `inject:"-"`
	SecretHandler *SecretHandler        `inject:"-"`
	Client        *kubernetes.Clientset `inject:"-"`
}

func NewSecretService() *SecretService {
	return &SecretService{}
}

// ListNamespace 返回ns命名空间下所有的secret列表
func (s *SecretService) ListNamespace(ns string) (ret []*models.Secret) {
	secretlist, err := s.SecretMap.ListByNamespace(ns)
	if err != nil {
		log.Println("SecretService:", err)
		return
	}
	for _, item := range secretlist {
		ret = append(ret, &models.Secret{
			Name:       item.Name,
			Namespace:  item.Namespace,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Type:       GetSecretType(item),
		})
	}
	return ret
}

func (s *SecretService) DelSecret(name, namespace string) error {
	err := s.Client.CoreV1().Secrets(namespace).Delete(context.Background(), name, v1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (s *SecretService) CreateSecret(post *models.SecretPost) error {
	_, err := s.Client.CoreV1().Secrets(post.Namespace).Create(context.Background(), &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      post.Name,
			Namespace: post.Namespace,
		},
		Type:       corev1.SecretType(post.Type),
		StringData: post.Data,
	}, v1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
