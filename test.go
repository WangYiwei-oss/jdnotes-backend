package main

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"os"
)

func main() {
	//创建客户端
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		log.Fatalln(err)
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	//配置option
	option := &corev1.PodExecOptions{
		Container: "demo",
		Command:   []string{"sh"},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}

	//创建req
	req := c.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace("default").
		Name("demo-65797b6745-2lglw").
		SubResource("exec").
		Param("color", "false").
		VersionedParams(option, scheme.ParameterCodec)
	fmt.Println("请求路径为:", req.URL())

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		log.Fatalln(err)
	}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    true,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
