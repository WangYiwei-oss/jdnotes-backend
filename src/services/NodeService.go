package services

import (
	"fmt"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"k8s.io/client-go/kubernetes"
)

type NodeService struct {
	NodeMap *NodeMap	`inject:"-"`
	CommonService *CommonService  `inject:"-"`
	NodeHandler *NodeHandler `inject:"-"`
	Client *kubernetes.Clientset `inject:"-"`
}

func NewNodeService() *NodeService {
	return &NodeService{}
}

func (n *NodeService) List() (ret []*models.Node) {
	nodeList := n.NodeMap.List()
	for _, item := range nodeList {
		fmt.Println("------",item.Status.Phase)

		ret = append(ret, &models.Node{
			Name:       item.Name,
			IP: item.Status.Addresses[0].Address,
			HostName: item.Status.Addresses[1].Address,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Phase: "",
			Labels: n.CommonService.SimpleMap2Slice(item.Labels),
		})
	}
	return
}

