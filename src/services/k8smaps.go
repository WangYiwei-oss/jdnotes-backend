package services

import (
	"fmt"
	"github.com/WangYiwei-oss/jdnotes-backend/src/helper"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1beta1 "k8s.io/api/networking/v1beta1"
	"sort"
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

////////////////////////////////////////////ServiceMap

type ServiceMap struct {
	data sync.Map
}

func NewServiceMap() *ServiceMap {
	return &ServiceMap{}
}

func (s *ServiceMap) Add(service *corev1.Service) {
	if list, ok := s.data.Load(service.Namespace); ok {
		list = append(list.([]*corev1.Service), service)
		s.data.Store(service.Namespace, list)
	} else {
		s.data.Store(service.Namespace, []*corev1.Service{service})
	}
}
func (s *ServiceMap) Update(service *corev1.Service) error {
	if list, ok := s.data.Load(service.Namespace); ok {
		for i, rangePod := range list.([]*corev1.Service) {
			if rangePod.Name == service.Name {
				list.([]*corev1.Service)[i] = service
			}
		}
		return nil
	}
	return fmt.Errorf("service-#{service.Name} not found")
}

func (s *ServiceMap) Delete(service *corev1.Service) {
	if list, ok := s.data.Load(service.Namespace); ok {
		for i, rangePod := range list.([]*corev1.Service) {
			if rangePod.Name == service.Name {
				newList := append(list.([]*corev1.Service)[:i], list.([]*corev1.Service)[i+1:]...)
				s.data.Store(service.Namespace, newList)
			}
		}
	}
}

func (s *ServiceMap) ListByNamespace(ns string) ([]*corev1.Service, error) {
	if list, ok := s.data.Load(ns); ok {
		return list.([]*corev1.Service), nil
	}
	return nil, fmt.Errorf("ServiceMap: namespace %s not found", ns)
}

func (s *ServiceMap) GetServiceByNamespace(ns, serviceName string) (*corev1.Service, error) {
	if list, ok := s.data.Load(ns); ok {
		for _, service := range list.([]*corev1.Service) {
			if service.Name == serviceName {
				return service, nil
			}
		}
		return nil, fmt.Errorf("ServiceMap: record %s.%s not found", ns, serviceName)
	}
	return nil, fmt.Errorf("ServiceMap: namespace %s not found", ns)
}

////////////////////////////////////////////IngressMap

type IngressMap struct {
	data sync.Map
}

func NewIngressMap() *IngressMap {
	return &IngressMap{}
}

func (i *IngressMap) Add(ingress *netv1beta1.Ingress) {
	if list, ok := i.data.Load(ingress.Namespace); ok {
		list = append(list.([]*netv1beta1.Ingress), ingress)
		i.data.Store(ingress.Namespace, list)
	} else {
		i.data.Store(ingress.Namespace, []*netv1beta1.Ingress{ingress})
	}
}
func (i *IngressMap) Update(ingress *netv1beta1.Ingress) error {
	if list, ok := i.data.Load(ingress.Namespace); ok {
		for i, rangeIngress := range list.([]*netv1beta1.Ingress) {
			if rangeIngress.Name == ingress.Name {
				list.([]*netv1beta1.Ingress)[i] = ingress
			}
		}
		return nil
	}
	return fmt.Errorf("IngressMap: Ingress-#{service.Name} not found")
}

func (i *IngressMap) Delete(ingress *netv1beta1.Ingress) {
	if list, ok := i.data.Load(ingress.Namespace); ok {
		for j, rangeIngress := range list.([]*netv1beta1.Ingress) {
			if rangeIngress.Name == ingress.Name {
				newList := append(list.([]*netv1beta1.Ingress)[:j], list.([]*netv1beta1.Ingress)[j+1:]...)
				i.data.Store(ingress.Namespace, newList)
			}
		}
	}
}

func (i *IngressMap) ListByNamespace(ns string) ([]*netv1beta1.Ingress, error) {
	if list, ok := i.data.Load(ns); ok {
		return list.([]*netv1beta1.Ingress), nil
	}
	return nil, fmt.Errorf("IngressMap: namespace %s not found", ns)
}

func (i *IngressMap) GetIngressByNamespace(ns, ingressName string) (*netv1beta1.Ingress, error) {
	if list, ok := i.data.Load(ns); ok {
		for _, ingress := range list.([]*netv1beta1.Ingress) {
			if ingress.Name == ingressName {
				return ingress, nil
			}
		}
		return nil, fmt.Errorf("IngressMap: record %s.%s not found", ns, ingressName)
	}
	return nil, fmt.Errorf("IngressMap: namespace %s not found", ns)
}

////////////////////////////////////////////SecretMap

type SecretMap struct {
	data sync.Map
}

func NewSecretMap() *SecretMap {
	return &SecretMap{}
}

func (s *SecretMap) Add(secret *corev1.Secret) {
	if list, ok := s.data.Load(secret.Namespace); ok {
		list = append(list.([]*corev1.Secret), secret)
		s.data.Store(secret.Namespace, list)
	} else {
		s.data.Store(secret.Namespace, []*corev1.Secret{secret})
	}
}
func (s *SecretMap) Update(secret *corev1.Secret) error {
	if list, ok := s.data.Load(secret.Namespace); ok {
		for i, rangeSecret := range list.([]*corev1.Secret) {
			if rangeSecret.Name == secret.Name {
				list.([]*corev1.Secret)[i] = secret
			}
		}
		return nil
	}
	return fmt.Errorf("SecretMap: Secret-#{secret.Name} not found")
}

func (s *SecretMap) Delete(secret *corev1.Secret) {
	if list, ok := s.data.Load(secret.Namespace); ok {
		for j, rangeIngress := range list.([]*corev1.Secret) {
			if rangeIngress.Name == secret.Name {
				newList := append(list.([]*corev1.Secret)[:j], list.([]*corev1.Secret)[j+1:]...)
				s.data.Store(secret.Namespace, newList)
			}
		}
	}
}

func (s *SecretMap) ListByNamespace(ns string) ([]*corev1.Secret, error) {
	if list, ok := s.data.Load(ns); ok {
		return list.([]*corev1.Secret), nil
	}
	return nil, fmt.Errorf("SecretMap: namespace %s not found", ns)
}

func (s *SecretMap) GetSecretByNamespace(ns, secretName string) (*corev1.Secret, error) {
	if list, ok := s.data.Load(ns); ok {
		for _, secret := range list.([]*corev1.Secret) {
			if secret.Name == secretName {
				return secret, nil
			}
		}
		return nil, fmt.Errorf("SecretMap: record %s.%s not found", ns, secretName)
	}
	return nil, fmt.Errorf("SecretMap: namespace %s not found", ns)
}

////////////////////////////////////////////ConfigMapMap
//由于configmap有一直刷新的问题，所以需要特殊适配

type cm struct {
	cmdata *corev1.ConfigMap
	md5    string
}

func newCm(cmdata *corev1.ConfigMap) *cm {
	return &cm{
		cmdata: cmdata,
		md5:    helper.Md5Data(cmdata.Data),
	}
}

//实现排序接口
type CoreV1ConfigMapMap []*cm

func (c CoreV1ConfigMapMap) Len() int {
	return len(c)
}

func (c CoreV1ConfigMapMap) Less(i, j int) bool {
	return c[i].cmdata.CreationTimestamp.Time.After(c[j].cmdata.CreationTimestamp.Time)
}

func (c CoreV1ConfigMapMap) Swap(i, j int) {
	c[i], c[j] = c[i], c[j]
}

type ConfigMapMap struct {
	data sync.Map
}

func NewConfigMapMap() *ConfigMapMap {
	return &ConfigMapMap{}
}

func (c *ConfigMapMap) Add(configmap *corev1.ConfigMap) {
	if list, ok := c.data.Load(configmap.Namespace); ok {
		list = append(list.([]*cm), newCm(configmap))
		c.data.Store(configmap.Namespace, list)
	} else {
		c.data.Store(configmap.Namespace, []*cm{newCm(configmap)})
	}
}

//true代表有值更新，否则返回false
func (c *ConfigMapMap) Update(configmap *corev1.ConfigMap) (bool, error) {
	if list, ok := c.data.Load(configmap.Namespace); ok {
		for i, item := range list.([]*cm) {
			if item.cmdata.Name == configmap.Name && !helper.CmIsEq(item.cmdata.Data, configmap.Data) {
				list.([]*cm)[i] = newCm(configmap)
				return true, nil
			}
		}
		return false, nil
	}
	return false, fmt.Errorf("ConfigMapMap: ConfigMapMap-#{configmap.Name} not found")
}

func (c *ConfigMapMap) Delete(configmap *corev1.ConfigMap) {
	if list, ok := c.data.Load(configmap.Namespace); ok {
		for j, rangeIngress := range list.([]*cm) {
			if rangeIngress.cmdata.Name == configmap.Name {
				newList := append(list.([]*cm)[:j], list.([]*cm)[j+1:]...)
				c.data.Store(configmap.Namespace, newList)
			}
		}
	}
}

func (c *ConfigMapMap) ListByNamespace(ns string) ([]*corev1.ConfigMap, error) {
	ret := make([]*corev1.ConfigMap, 0)
	if list, ok := c.data.Load(ns); ok {
		cmlist := list.([]*cm)
		sort.Sort(CoreV1ConfigMapMap(cmlist))
		for _, cm := range cmlist {
			ret = append(ret, cm.cmdata)
		}
		return ret, nil
	}
	return nil, fmt.Errorf("ConfigMapMap: namespace %s not found", ns)
}

func (c *ConfigMapMap) GetConfigMapByNamespace(ns, configMapName string) (*corev1.ConfigMap, error) {
	if list, ok := c.data.Load(ns); ok {
		for _, item := range list.([]*cm) {
			if item.cmdata.Name == configMapName {
				return item.cmdata, nil
			}
		}
		return nil, fmt.Errorf("ConfigMapMap: record %s.%s not found", ns, configMapName)
	}
	return nil, fmt.Errorf("ConfigMapMap: namespace %s not found", ns)
}
