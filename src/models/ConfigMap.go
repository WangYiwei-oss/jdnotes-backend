package models

type ConfigMap struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	CreateTime string            `json:"create_time"`
	Data       map[string]string `json:"data"`
}

type ConfigMapPost struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Data      map[string]string `json:"data"`
}
