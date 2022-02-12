package services

import (
	"context"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

var ServiceTypeMap map[corev1.ServiceType]string
var ProtoMap map[corev1.Protocol]string

func init() {
	ServiceTypeMap = map[corev1.ServiceType]string{
		corev1.ServiceTypeClusterIP:    "ClusterIP",
		corev1.ServiceTypeNodePort:     "NodePort",
		corev1.ServiceTypeLoadBalancer: "LoadBalancer",
		corev1.ServiceTypeExternalName: "ExternalName",
	}
	ProtoMap = map[corev1.Protocol]string{
		corev1.ProtocolTCP:  "TCP",
		corev1.ProtocolUDP:  "UDP",
		corev1.ProtocolSCTP: "SCTP",
	}
}

type ServiceService struct {
	ServiceMap      *ServiceMap           `inject:"-"`
	CommonService   *CommonService        `inject:"-"`
	ServicesHandler *ServiceHandler       `inject:"-"`
	Client          *kubernetes.Clientset `inject:"-"`
}

func NewServiceService() *ServiceService {
	return &ServiceService{}
}

func (s *ServiceService) ListNamespace(ns string) (ret []*models.Service) {
	servicelist, err := s.ServiceMap.ListByNamespace(ns)
	if err != nil {
		log.Println("ServiceService:", err)
		return
	}
	for _, item := range servicelist {
		ret = append(ret, &models.Service{
			Name:        item.Name,
			Namespace:   item.Namespace,
			ClusterIp:   item.Spec.ClusterIP,
			ServiceType: ServiceTypeMap[item.Spec.Type],
			Ports:       s.CommonService.GetServicePorts(item.Spec.Ports),
			CreateTime:  item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return ret
}

func (s *ServiceService) DeleteService(name, namespace string) error {
	return s.Client.CoreV1().Services(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}
