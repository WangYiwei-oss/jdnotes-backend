package services

import (
	"context"
	"fmt"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"golang.org/x/crypto/ssh"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"net"
)

type NodeService struct {
	NodeMap         *NodeMap              `inject:"-"`
	CommonService   *CommonService        `inject:"-"`
	NodeHandler     *NodeHandler          `inject:"-"`
	Client          *kubernetes.Clientset `inject:"-"`
	VersionedClient *versioned.Clientset  `inject:"-"`
	PodMap          *PodMap               `inject:"-"`
}

func NewNodeService() *NodeService {
	return &NodeService{}
}

func (n *NodeService) GetNodeStatus(node *corev1.Node) string {
	for _, item := range node.Status.Conditions {
		if item.Type == "Ready" && item.Status == "True" {
			return "Ready"
		}
	}
	return "NotReady"
}

func (n *NodeService) List() (ret []*models.Node) {
	nodeList := n.NodeMap.List()
	for _, item := range nodeList {
		ret = append(ret, &models.Node{
			Name:       item.Name,
			IP:         item.Status.Addresses[0].Address,
			HostName:   item.Status.Addresses[1].Address,
			Phase:      n.GetNodeStatus(item),
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return
}

func (n *NodeService) TaintsToSlice(taints []corev1.Taint) []string {
	ret := make([]string, 0)
	for _, taint := range taints {
		s := string(taint.Effect) + ": " + taint.Key + "=" + taint.Value
		ret = append(ret, s)
	}
	return ret
}

func (n *NodeService) GetNodeUsage(c *versioned.Clientset, node *corev1.Node) []float64 {
	nodeMetric, err := c.MetricsV1beta1().NodeMetricses().Get(context.Background(), node.Name, metav1.GetOptions{})
	if err != nil {
		return nil
	}
	cpu := float64(nodeMetric.Usage.Cpu().MilliValue())
	memory := float64(nodeMetric.Usage.Memory().MilliValue())
	return []float64{
		cpu, memory,
	}
}

func (n *NodeService) GetNode(name string) (*models.NodeDetail, error) {
	node, err := n.NodeMap.Get(name)
	if err != nil {
		return nil, err
	}
	if n.GetNodeStatus(node) == "NotReady" {
		return &models.NodeDetail{
			Name:       node.Name,
			IP:         node.Status.Addresses[0].Address,
			HostName:   node.Status.Addresses[1].Address,
			CreateTime: node.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Phase:      "NotReady",
			Labels:     n.CommonService.SimpleMap2Slice(node.Labels),
			Taints:     n.TaintsToSlice(node.Spec.Taints),
			Capacity:   models.NewNodeCapacity(0, 0, 0),
			Usage:      models.NewNodeUsage(0, 0, 0),
		}, nil
	}

	usages := n.GetNodeUsage(n.VersionedClient, node)
	ret := &models.NodeDetail{
		Name:       node.Name,
		IP:         node.Status.Addresses[0].Address,
		HostName:   node.Status.Addresses[1].Address,
		CreateTime: node.CreationTimestamp.Format("2006-01-02 15:04:05"),
		Phase:      n.GetNodeStatus(node),
		Labels:     n.CommonService.SimpleMap2Slice(node.Labels),
		Taints:     n.TaintsToSlice(node.Spec.Taints),
		Capacity:   models.NewNodeCapacity(node.Status.Capacity.Cpu().Value(), node.Status.Capacity.Memory().Value(), node.Status.Capacity.Pods().Value()),
		Usage:      models.NewNodeUsage(usages[0], usages[1], n.PodMap.GetPodNumberByNode(node.Name)),
	}
	return ret, nil
}

//ssh相关
func (n *NodeService) ConnectNodeShell(post *models.NodeShellPost) (*ssh.Session, error) {
	auth := []ssh.AuthMethod{
		ssh.Password(post.Password),
	}
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig := &ssh.ClientConfig{
		User:            post.User,
		Auth:            auth,
		HostKeyCallback: hostKeyCallbk,
	}
	addr := fmt.Sprintf("%s:%d", post.Ip, 22)
	var client *ssh.Client
	var err error
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	var session *ssh.Session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}
