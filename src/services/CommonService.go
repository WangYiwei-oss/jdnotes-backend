package services

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"strconv"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (c *CommonService) GetImagesFromDeployment(dep v1.Deployment) string {
	return c.GetImagesByPod(dep.Spec.Template.Spec.Containers)
}

func (c *CommonService) GetImagesByPod(containers []corev1.Container) string {
	if len(containers) < 1 {
		return ""
	}
	images := ""
	for _, container := range containers {
		images = images + container.Image + "/"
	}
	return images[:len(images)]
}

func (c *CommonService) GetPodIsReady(pod *corev1.Pod) bool {
	if pod.Status.Phase != "Running" {
		return false
	}
	/*
		Pod的PodConditions共有四种状态:
		1. PodScheduled: Pod已经被调度到某节点
		2. ContainersReady: Pod中所有容器都准备就绪了
		3. Initialized: 所有的Init容器都已经成功启动
		4. Ready: Pod可以为请求提供服务

		这里取所有状态都为true才行
	*/
	for _, condition := range pod.Status.Conditions {
		if condition.Status != "True" {
			return false
		}
	}
	//pod.Spec.ReadinessGates是额外的自定义的Conditions，也要进行判断
	for _, rg := range pod.Spec.ReadinessGates {
		for _, condition := range pod.Status.Conditions {
			if condition.Type == rg.ConditionType && condition.Status != "True" {
				return false
			}
		}
	}
	return true
}

func (c *CommonService)GetNodeIsReady(node *corev1.Node)bool{
	for _,condition := range node.Status.Conditions{
		if condition.Status!=corev1.ConditionTrue{
			return false
		}
	}
	return true
}

func (c *CommonService) GetServicePorts(servicePorts []corev1.ServicePort) string {
	ret := ""
	if servicePorts == nil || len(servicePorts) == 0 {
		return ret
	}
	for _, port := range servicePorts {
		nodeport := ""
		if port.NodePort != 0 {
			nodeport = ":" + strconv.Itoa(int(port.NodePort))
		}
		ret = ret + port.Name + "/" + strconv.Itoa(int(port.Port)) + nodeport + "/" + ProtoMap[port.Protocol] + ","
	}
	return ret[:len(ret)-1]
}

func (c *CommonService) GetPodMounts(container *corev1.Pod) string {
	ret := ""
	return ret
}

func (c *CommonService) SimpleMap2String(labels map[string]string) (ret string) {
	for k, v := range labels {
		if ret != "" {
			ret += ","
		}
		ret += fmt.Sprintf("%s=%s", string(k), string(v))
	}
	return
}

func (c *CommonService) SimpleMap2Slice(labels map[string]string) []string {
	ret:=make([]string,0)
	for k, v := range labels {
		ret = append(ret,fmt.Sprintf("%s=%s", string(k), string(v)))
	}
	return ret
}
