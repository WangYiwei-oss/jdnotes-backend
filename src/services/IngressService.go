package services

import (
	"context"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	netv1beta1 "k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"log"
)

type IngressService struct {
	IngressMap     *IngressMap           `inject:"-"`
	CommonService  *CommonService        `inject:"-"`
	IngressHandler *IngressHandler       `inject:"-"`
	Client         *kubernetes.Clientset `inject:"-"`
}

func NewIngressService() *IngressService {
	return &IngressService{}
}

// ListNamespace 返回ns命名空间下所有的service列表
func (i *IngressService) ListNamespace(ns string) (ret []*models.Ingress) {
	ingresslist, err := i.IngressMap.ListByNamespace(ns)
	if err != nil {
		log.Println("IngressService:", err)
		return
	}
	for _, item := range ingresslist {
		ret = append(ret, &models.Ingress{
			Name:       item.Name,
			Namespace:  item.Namespace,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return ret
}

func (i *IngressService) CreateIngress(post *models.IngressPost) error {
	className := "nginx"
	ingressRules := []netv1beta1.IngressRule{}
	for _, r := range post.Rules {
		httpRuleValue := &netv1beta1.HTTPIngressRuleValue{}
		rulePaths := make([]netv1beta1.HTTPIngressPath, 0)
		for _, pathCfg := range r.Paths {
			rulePaths = append(rulePaths, netv1beta1.HTTPIngressPath{
				Path: pathCfg.Path,
				Backend: netv1beta1.IngressBackend{
					ServiceName: pathCfg.SvcName,
					ServicePort: intstr.FromInt(pathCfg.Port),
				},
			})
		}
		httpRuleValue.Paths = rulePaths
		rule := netv1beta1.IngressRule{
			Host: r.Host,
			IngressRuleValue: netv1beta1.IngressRuleValue{
				HTTP: httpRuleValue,
			},
		}
		ingressRules = append(ingressRules, rule)
	}

	ingress := &netv1beta1.Ingress{
		TypeMeta: v1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "networking.k8s.io/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      post.Name,
			Namespace: post.Namespace,
		},
		Spec: netv1beta1.IngressSpec{
			IngressClassName: &className,
			Rules:            ingressRules,
		},
	}

	_, err := i.Client.NetworkingV1beta1().Ingresses(post.Namespace).Create(context.Background(), ingress, v1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
