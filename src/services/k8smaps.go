package services

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sync"
)

//////////////////////////////////DeploymentMap

type DeploymentMap struct {
	data sync.Map
}

func NewDeploymentMap() *DeploymentMap {
	return &DeploymentMap{}
}

func (d *DeploymentMap) Add(dep *v1.Deployment) {
	if list, ok := d.data.Load(dep.Namespace); ok {
		list = append(list.([]*v1.Deployment), dep)
		d.data.Store(dep.Namespace, list)
	} else {
		d.data.Store(dep.Namespace, []*v1.Deployment{dep})
	}
}
func (d *DeploymentMap) Update(dep *v1.Deployment) error {
	if list, ok := d.data.Load(dep.Namespace); ok {
		for i, range_dep := range list.([]*v1.Deployment) {
			if range_dep.Name == dep.Name {
				list.([]*v1.Deployment)[i] = dep
			}
		}
		return nil
	}
	return fmt.Errorf("deployment-#{dep.Name} not found")
}

func (d *DeploymentMap) Delete(dep *v1.Deployment) {
	if list, ok := d.data.Load(dep.Namespace); ok {
		for i, range_dep := range list.([]*v1.Deployment) {
			if range_dep.Name == dep.Name {
				newList := append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				d.data.Store(dep.Namespace, newList)
			}
		}
	}
}

func (d *DeploymentMap) DeleteNamespace(ns *corev1.Namespace) {
	if _, ok := d.data.Load(ns.Name); ok {
		d.data.Delete(ns.Namespace)
	}
}

func (d *DeploymentMap) ListByNamespace(ns string) ([]*v1.Deployment, error) {
	if list, ok := d.data.Load(ns); ok {
		return list.([]*v1.Deployment), nil
	}
	return nil, fmt.Errorf("namespace %s not found", ns)
}

func (d *DeploymentMap) GetDeploymentByNamespace(ns, depName string) (*v1.Deployment, error) {
	if list, ok := d.data.Load(ns); ok {
		for _, dep := range list.([]*v1.Deployment) {
			if dep.Name == depName {
				return dep, nil
			}
		}
		return nil, fmt.Errorf("record %s.%s not found", ns, depName)
	}
	return nil, fmt.Errorf("namespace %s not found", ns)
}

////////////////////////////////////////////PodMap

type PodMap struct {
	data sync.Map
}

func NewPodMap() *PodMap {
	return &PodMap{}
}

func (p *PodMap) Add(pod *corev1.Pod) {
	if list, ok := p.data.Load(pod.Namespace); ok {
		list = append(list.([]*corev1.Pod), pod)
		p.data.Store(pod.Namespace, list)
	} else {
		p.data.Store(pod.Namespace, []*corev1.Pod{pod})
	}
}
func (p *PodMap) Update(pod *corev1.Pod) error {
	if list, ok := p.data.Load(pod.Namespace); ok {
		for i, rangePod := range list.([]*corev1.Pod) {
			if rangePod.Name == pod.Name {
				list.([]*corev1.Pod)[i] = pod
			}
		}
		return nil
	}
	return fmt.Errorf("pod-#{pod.Name} not found")
}

func (p *PodMap) Delete(pod *corev1.Pod) {
	if list, ok := p.data.Load(pod.Namespace); ok {
		for i, rangePod := range list.([]*corev1.Pod) {
			if rangePod.Name == pod.Name {
				newList := append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[i+1:]...)
				p.data.Store(pod.Namespace, newList)
			}
		}
	}
}

func (p *PodMap) ListByNamespace(ns string) ([]*corev1.Pod, error) {
	if list, ok := p.data.Load(ns); ok {
		return list.([]*corev1.Pod), nil
	}
	return nil, fmt.Errorf("namespace %s not found", ns)
}

func (p *PodMap) GetPodByNamespace(ns, podName string) (*corev1.Pod, error) {
	if list, ok := p.data.Load(ns); ok {
		for _, pod := range list.([]*corev1.Pod) {
			if pod.Name == podName {
				return pod, nil
			}
		}
		return nil, fmt.Errorf("record %s.%s not found", ns, podName)
	}
	return nil, fmt.Errorf("namespace %s not found", ns)
}

////////////////////////////////////////////NamespaceMap

type NamespaceMap struct {
	data sync.Map
}

func NewNamespaceMap() *NamespaceMap {
	return &NamespaceMap{}
}

func (n *NamespaceMap) Add(ns *corev1.Namespace) {
	n.data.Store(ns.Name, ns)
}
func (n *NamespaceMap) Update(ns *corev1.Namespace) error {
	n.data.Store(ns.Name, ns)
	return nil
}

func (n *NamespaceMap) Delete(ns *corev1.Namespace) {
	n.data.Delete(ns.Name)
}

func (n *NamespaceMap) ListNamespaces() ([]*corev1.Namespace, error) {
	ret := make([]*corev1.Namespace, 0)
	n.data.Range(func(key, value interface{}) bool {
		ret = append(ret, value.(*corev1.Namespace))
		return true
	})
	return ret, nil
}
