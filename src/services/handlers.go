package services

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1beta1 "k8s.io/api/networking/v1beta1"
	"log"
)

///////////////////DepHandler

type DepHandler struct {
	DepMap     *DeploymentMap     `inject:"-"`
	DepService *DeploymentService `inject:"-"`
}

func (d *DepHandler) OnAdd(obj interface{}) {
	fmt.Println("创建了dep")
	d.DepMap.Add(obj.(*v1.Deployment))
	jdft.WebSocketFactory.SendAllClass("deployments",
		d.DepService.ListNamespace(obj.(*v1.Deployment).Namespace),
		obj.(*v1.Deployment).Namespace)
}
func (d *DepHandler) OnUpdate(oldObj, newObj interface{}) {
	err := d.DepMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	} else {
		jdft.WebSocketFactory.SendAllClass("deployments",
			d.DepService.ListNamespace(newObj.(*v1.Deployment).Namespace),
			newObj.(*v1.Deployment).Namespace)
	}
}

func (d *DepHandler) OnDelete(obj interface{}) {
	fmt.Println("删除了dep") //？？？？？？？？？？？？？为什么不删除
	d.DepMap.Delete(obj.(*v1.Deployment))
	jdft.WebSocketFactory.SendAllClass("deployments",
		d.DepService.ListNamespace(obj.(*v1.Deployment).Namespace),
		obj.(*v1.Deployment).Namespace)
}

///////////////////PodHandler

type PodHandler struct {
	PodMap     *PodMap     `inject:"-"`
	PodService *PodService `inject:"-"`
}

func (p *PodHandler) OnAdd(obj interface{}) {
	p.PodMap.Add(obj.(*corev1.Pod))
	jdft.WebSocketFactory.SendAllClass("pods",
		p.PodService.ListNamespace(obj.(*corev1.Pod).Namespace),
		obj.(*corev1.Pod).Namespace)
}
func (p *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	err := p.PodMap.Update(newObj.(*corev1.Pod))
	if err != nil {
		log.Println(err)
	}
	jdft.WebSocketFactory.SendAllClass("pods",
		p.PodService.ListNamespace(newObj.(*corev1.Pod).Namespace),
		newObj.(*corev1.Pod).Namespace)
}

func (p *PodHandler) OnDelete(obj interface{}) {
	p.PodMap.Delete(obj.(*corev1.Pod))
	jdft.WebSocketFactory.SendAllClass("pods",
		p.PodService.ListNamespace(obj.(*corev1.Pod).Namespace),
		obj.(*corev1.Pod).Namespace)
}

///////////////////NamespaceHandler

type NamespaceHandler struct {
	NamespaceMap *NamespaceMap `inject:"-"`
}

func (n *NamespaceHandler) OnAdd(obj interface{}) {
	n.NamespaceMap.Add(obj.(*corev1.Namespace))
}
func (n *NamespaceHandler) OnUpdate(oldObj, newObj interface{}) {
	err := n.NamespaceMap.Update(newObj.(*corev1.Namespace))
	if err != nil {
		log.Println(err)
	}
}

func (n *NamespaceHandler) OnDelete(obj interface{}) {
	n.NamespaceMap.Delete(obj.(*corev1.Namespace))
}

///////////////////ServiceHandler

type ServiceHandler struct {
	ServiceMap     *ServiceMap     `inject:"-"`
	ServiceService *ServiceService `inject:"-"`
}

func (s *ServiceHandler) OnAdd(obj interface{}) {
	s.ServiceMap.Add(obj.(*corev1.Service))
	jdft.WebSocketFactory.SendAllClass("services",
		s.ServiceService.ListNamespace(obj.(*corev1.Service).Namespace),
		obj.(*corev1.Service).Namespace)
}
func (s *ServiceHandler) OnUpdate(oldObj, newObj interface{}) {
	err := s.ServiceMap.Update(newObj.(*corev1.Service))
	if err != nil {
		log.Println(err)
	}
	jdft.WebSocketFactory.SendAllClass("services",
		s.ServiceService.ListNamespace(newObj.(*corev1.Service).Namespace),
		newObj.(*corev1.Service).Namespace)
}

func (s *ServiceHandler) OnDelete(obj interface{}) {
	s.ServiceMap.Delete(obj.(*corev1.Service))
	jdft.WebSocketFactory.SendAllClass("services",
		s.ServiceService.ListNamespace(obj.(*corev1.Service).Namespace),
		obj.(*corev1.Service).Namespace)
}

///////////////////IngressHandler

type IngressHandler struct {
	IngressMap     *IngressMap     `inject:"-"`
	IngressService *IngressService `inject:"-"`
}

func (i *IngressHandler) OnAdd(obj interface{}) {
	i.IngressMap.Add(obj.(*netv1beta1.Ingress))
	jdft.WebSocketFactory.SendAllClass("ingresses",
		i.IngressService.ListNamespace(obj.(*netv1beta1.Ingress).Namespace),
		obj.(*netv1beta1.Ingress).Namespace)
}
func (i *IngressHandler) OnUpdate(oldObj, newObj interface{}) {
	err := i.IngressMap.Update(newObj.(*netv1beta1.Ingress))
	if err != nil {
		log.Println(err)
	}
	jdft.WebSocketFactory.SendAllClass("ingresses",
		i.IngressService.ListNamespace(newObj.(*netv1beta1.Ingress).Namespace),
		newObj.(*netv1beta1.Ingress).Namespace)
}

func (i *IngressHandler) OnDelete(obj interface{}) {
	i.IngressMap.Delete(obj.(*netv1beta1.Ingress))
	jdft.WebSocketFactory.SendAllClass("ingresses",
		i.IngressService.ListNamespace(obj.(*netv1beta1.Ingress).Namespace),
		obj.(*netv1beta1.Ingress).Namespace)
}

///////////////////SecretHandler

type SecretHandler struct {
	SecretMap     *SecretMap     `inject:"-"`
	SecretService *SecretService `inject:"-"`
}

func (s *SecretHandler) OnAdd(obj interface{}) {
	s.SecretMap.Add(obj.(*corev1.Secret))
	jdft.WebSocketFactory.SendAllClass("secrets",
		s.SecretService.ListNamespace(obj.(*corev1.Secret).Namespace),
		obj.(*corev1.Secret).Namespace)
}
func (s *SecretHandler) OnUpdate(oldObj, newObj interface{}) {
	err := s.SecretMap.Update(newObj.(*corev1.Secret))
	if err != nil {
		log.Println(err)
	}
	jdft.WebSocketFactory.SendAllClass("secrets",
		s.SecretService.ListNamespace(newObj.(*corev1.Secret).Namespace),
		newObj.(*corev1.Secret).Namespace)
}

func (s *SecretHandler) OnDelete(obj interface{}) {
	s.SecretMap.Delete(obj.(*corev1.Secret))
	jdft.WebSocketFactory.SendAllClass("secrets",
		s.SecretService.ListNamespace(obj.(*corev1.Secret).Namespace),
		obj.(*corev1.Secret).Namespace)
}

///////////////////ConfigMapHandler

type ConfigMapHandler struct {
	ConfigMapMap     *ConfigMapMap     `inject:"-"`
	ConfigMapService *ConfigMapService `inject:"-"`
}

func (c *ConfigMapHandler) OnAdd(obj interface{}) {
	c.ConfigMapMap.Add(obj.(*corev1.ConfigMap))
	jdft.WebSocketFactory.SendAllClass("configmaps",
		c.ConfigMapService.ListNamespace(obj.(*corev1.ConfigMap).Namespace),
		obj.(*corev1.ConfigMap).Namespace)
}
func (c *ConfigMapHandler) OnUpdate(oldObj, newObj interface{}) {
	isupdate, err := c.ConfigMapMap.Update(newObj.(*corev1.ConfigMap))
	if err != nil {
		log.Println(err)
	}
	if isupdate {
		fmt.Println("config更新并通知")
		jdft.WebSocketFactory.SendAllClass("configmaps",
			c.ConfigMapService.ListNamespace(newObj.(*corev1.ConfigMap).Namespace),
			newObj.(*corev1.ConfigMap).Namespace)
	}
}

func (c *ConfigMapHandler) OnDelete(obj interface{}) {
	c.ConfigMapMap.Delete(obj.(*corev1.ConfigMap))
	jdft.WebSocketFactory.SendAllClass("configmaps",
		c.ConfigMapService.ListNamespace(obj.(*corev1.ConfigMap).Namespace),
		obj.(*corev1.ConfigMap).Namespace)
}

///////////////////NodeHandler

type NodeHandler struct {
	NodeMap     *NodeMap     `inject:"-"`
	NodeService *NodeService `inject:"-"`
}

func (n *NodeHandler) OnAdd(obj interface{}) {
	n.NodeMap.Add(obj.(*corev1.Node))
	jdft.WebSocketFactory.SendAllClass("nodes",
		n.NodeService.List(),
		nil)
}
func (n *NodeHandler) OnUpdate(oldObj, newObj interface{}) {
	err := n.NodeMap.Update(newObj.(*corev1.Node))
	if err != nil {
		log.Println(err)
	}
	jdft.WebSocketFactory.SendAllClass("nodes",
		n.NodeService.List(),
		nil)
}

func (n *NodeHandler) OnDelete(obj interface{}) {
	n.NodeMap.Delete(obj.(*corev1.Node))
	jdft.WebSocketFactory.SendAllClass("nodes",
		n.NodeService.List(),
		nil)
}
