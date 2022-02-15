package config

import (
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type K8sConfig struct {
	K8sRestConfig    *rest.Config
	DepHandler       *services.DepHandler       `inject:"-"`
	PodHandler       *services.PodHandler       `inject:"-"`
	NamespaceHandler *services.NamespaceHandler `inject:"-"`
	ServiceHandler   *services.ServiceHandler   `inject:"-"`
	IngressHandler   *services.IngressHandler   `inject:"-"`
	SecretHandler    *services.SecretHandler    `inject:"-"`
	ConfigMapHandler *services.ConfigMapHandler `inject:"-"`
	NodeHandler      *services.NodeHandler	    `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (k *K8sConfig) JdInitK8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

func (k *K8sConfig) JdInitClient() *kubernetes.Clientset {
	//ip := jdft.GetGlobalSettings()["KUBERNETES_HOST"]
	//config := &rest.Config{
	//	Host:        fmt.Sprintf("http://%s", ip.(string)),
	//	BearerToken: "6c38472af3b00688ab7929c185b56bc6",
	//}
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		log.Fatalln(err)
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

	namespaceInformer := fact.Core().V1().Namespaces()
	namespaceInformer.Informer().AddEventHandler(k.NamespaceHandler)

	serviceInformer := fact.Core().V1().Services()
	serviceInformer.Informer().AddEventHandler(k.ServiceHandler)

	ingressInformer := fact.Networking().V1beta1().Ingresses() //监听Ingress
	ingressInformer.Informer().AddEventHandler(k.IngressHandler)

	secretInformer := fact.Core().V1().Secrets()
	secretInformer.Informer().AddEventHandler(k.SecretHandler)

	configInformer := fact.Core().V1().ConfigMaps()
	configInformer.Informer().AddEventHandler(k.ConfigMapHandler)

	nodeInformer := fact.Core().V1().Nodes()
	nodeInformer.Informer().AddEventHandler(k.NodeHandler)

	fact.Start(wait.NeverStop)
	return fact
}
