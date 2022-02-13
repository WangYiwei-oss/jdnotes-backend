package services

import (
	"context"
	"fmt"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"log"
)

var PodPhase map[corev1.PodPhase]string

func init() {
	PodPhase = map[corev1.PodPhase]string{
		corev1.PodPending:   "Pending",
		corev1.PodRunning:   "Running",
		corev1.PodSucceeded: "Succeeded",
		corev1.PodFailed:    "Failed",
		corev1.PodUnknown:   "Unknown",
	}
}

type PodService struct {
	PodMap        *PodMap               `inject:"-"`
	CommonService *CommonService        `inject:"-"`
	PodHandler    *PodHandler           `inject:"-"`
	Client        *kubernetes.Clientset `inject:"-"`
	K8sRestConfig *rest.Config          `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (p *PodService) ListNamespace(ns string) (ret []*models.Pod) {
	podlist, err := p.PodMap.ListByNamespace(ns)
	if err != nil {
		log.Println("PodService:", err)
		return
	}
	for _, item := range podlist {
		ret = append(ret, &models.Pod{
			Name:       item.Name,
			Namespace:  item.Namespace,
			Images:     p.CommonService.GetImagesByPod(item.Spec.Containers),
			NodeName:   item.Spec.NodeName,
			IP:         []string{item.Status.PodIP, item.Status.HostIP},
			Phase:      PodPhase[item.Status.Phase],
			IsReady:    p.CommonService.GetPodIsReady(item),
			Message:    item.Status.Message,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return
}

func (p *PodService) DeletePod(name, namespace string) error {
	return p.Client.CoreV1().Pods(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}

func (p *PodService) GetPodContainer(name, namespace string) ([]*models.Container, error) {
	ret := make([]*models.Container, 0)
	pod, err := p.PodMap.GetPodByNamespace(namespace, name)
	if err != nil {
		return nil, err
	}
	for _, c := range pod.Spec.Containers {
		ret = append(ret, &models.Container{
			Name:  c.Name,
			Image: c.Image,
		})
	}
	return ret, nil
}

func (p *PodService) GetPodDetail(name, namespace string) (*models.PodDetail, error) {
	pod, err := p.PodMap.GetPodByNamespace(namespace, name)
	if err != nil {
		return nil, err
	}
	containers := make([]*models.Container, 0)
	for _, c := range pod.Spec.Containers {
		containers = append(containers, &models.Container{
			Name:       c.Name,
			Image:      c.Image,
			Command:    c.Command,
			Args:       c.Args,
			WorkingDir: c.WorkingDir,
		})
		fmt.Println("pppppppppp", c.VolumeMounts)
	}
	fmt.Println("vvvvvvvvvv", pod.Spec.Volumes)
	ret := &models.PodDetail{
		Name:        pod.Name,
		Namespace:   pod.Namespace,
		Images:      p.CommonService.GetImagesByPod(pod.Spec.Containers),
		NodeName:    pod.Spec.NodeName,
		IP:          []string{pod.Status.PodIP, pod.Status.HostIP},
		Phase:       PodPhase[pod.Status.Phase],
		IsReady:     p.CommonService.GetPodIsReady(pod),
		Message:     pod.Status.Message,
		CreateTime:  pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		Labels:      p.CommonService.SimpleMap2String(pod.Labels),
		Annotations: p.CommonService.SimpleMap2String(pod.Annotations),
		Containers:  containers,
	}
	return ret, nil
}

//Exec相关

func (p *PodService) HandleCommand(client *kubernetes.Clientset, config *rest.Config, command []string) (remotecommand.Executor, error) {
	option := &corev1.PodExecOptions{
		Container: "demo",
		Command:   command,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}
	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace("default").
		Name("demo-65797b6745-2lglw").
		SubResource("exec").
		Param("color", "false").
		VersionedParams(option, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return nil, err
	}
	return exec, nil
}
