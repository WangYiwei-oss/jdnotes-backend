package models

type Ingress struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	CreateTime string `json:"create_time"`
}