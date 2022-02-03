package services

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"log"
)

///////////////////DepHandler

type DepHandler struct {
	DepMap     *DeploymentMap     `inject:"-"`
	DepService *DeploymentService `inject:"-"`
}

func (d *DepHandler) OnAdd(obj interface{}) {
	d.DepMap.Add(obj.(*v1.Deployment))
	jdft.WebSocketFactory.SendAllClass("deployments", d.DepService.ListNamespace(obj.(*v1.Deployment).Namespace))
}
func (d *DepHandler) OnUpdate(oldObj, newObj interface{}) {
	err := d.DepMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	} else {
		jdft.WebSocketFactory.SendAllClass("deployments", d.DepService.ListNamespace(newObj.(*v1.Deployment).Namespace))
	}
}

func (d *DepHandler) OnDelete(obj interface{}) {
	d.DepMap.Delete(obj.(*v1.Deployment))
	jdft.WebSocketFactory.SendAllClass("deployments", d.DepService.ListNamespace(obj.(*v1.Deployment).Namespace))
}

///////////////////PodHandler

type PodHandler struct {
	PodMap *PodMap `inject:"-"`
}

func (p *PodHandler) OnAdd(obj interface{}) {
	p.PodMap.Add(obj.(*corev1.Pod))
}
func (p *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	err := p.PodMap.Update(newObj.(*corev1.Pod))
	if err != nil {
		log.Println(err)
	}
}

func (p *PodHandler) OnDelete(obj interface{}) {
	p.PodMap.Delete(obj.(*corev1.Pod))
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
