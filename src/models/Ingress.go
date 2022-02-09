package models

type Ingress struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	CreateTime string `json:"create_time"`
}

type IngressPath struct {
	Path    string `json:"path"`
	SvcName string `json:"svc_name"`
	Port    int    `json:"port"`
}

type IngressRules struct {
	Host  string         `json:"host"`
	Paths []*IngressPath `json:"paths"`
}
type IngressPost struct {
	Name      string          `json:"name"`
	Namespace string          `json:"namespace"`
	Rules     []*IngressRules `json:"rules"`
}
