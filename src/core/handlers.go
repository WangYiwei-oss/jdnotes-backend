package core

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"log"
)

///////////////////DepHandler

type DepHandler struct {
	DepMap *DeploymentMap `inject:"-"`
}

func (d *DepHandler) OnAdd(obj interface{}) {
	d.DepMap.Add(obj.(*v1.Deployment))
}
func (d *DepHandler) OnUpdate(oldObj, newObj interface{}) {
	err := d.DepMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	}
}

func (d *DepHandler) OnDelete(obj interface{}) {
	d.DepMap.Delete(obj.(*v1.Deployment))
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
