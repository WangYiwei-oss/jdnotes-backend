package models

type Service struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	ClusterIp   string `json:"cluster_ip"`
	ServiceType string `json:"service_type"`
	Ports       string `json:"ports"`
	CreateTime  string `json:"create_time"`
}
