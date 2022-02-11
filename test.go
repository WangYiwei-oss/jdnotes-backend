package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		log.Fatalln(err)
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	list, err := c.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, pod := range list.Items {
		fmt.Println(pod.Name)
	}
}
