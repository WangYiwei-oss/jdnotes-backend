package services

import (
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"log"
)

type IngressService struct {
	IngressMap     *IngressMap     `inject:"-"`
	CommonService  *CommonService  `inject:"-"`
	IngressHandler *IngressHandler `inject:"-"`
}

func NewIngressService() *IngressService {
	return &IngressService{}
}

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
