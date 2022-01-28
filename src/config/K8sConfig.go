package config

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdnotes-backend/src/core"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

type K8sConfig struct {
	DepHandler *core.DepHandler `inject:"-"`
	PodHandler *core.PodHandler `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (*K8sConfig) JdInitClient() *kubernetes.Clientset {
	ip := jdft.GetGlobalSettings()["KUBERNETES_HOST"]
	config := &rest.Config{
		Host:        fmt.Sprintf("http://%s", ip.(string)),
		BearerToken: "6c38472af3b00688ab7929c185b56bc6",
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("init k8s client error", err)
	}
	return c
}

func (k *K8sConfig) JdInitInformer() informers.SharedInformerFactory {
	fact := informers.NewSharedInformerFactory(k.JdInitClient(), 0)

	depInformer := fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(k.DepHandler)

	podInformer := fact.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(k.PodHandler)

	fact.Start(wait.NeverStop)
	return fact
}
