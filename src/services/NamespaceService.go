package services

import (
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"log"
)

type NamespaceService struct {
	NamespaceMap *NamespaceMap `inject:"-"`
}

func NewNamespaceService() *NamespaceService {
	return &NamespaceService{}
}

func (n *NamespaceService) ListNamespaces() (ret []*models.Namespace) {
	namespaces, err := n.NamespaceMap.ListNamespaces()
	if err != nil {
		log.Println(err)
	}
	for _, ns := range namespaces {
		ret = append(ret, &models.Namespace{
			Name:       ns.Name,
			CreateTime: ns.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return
}
