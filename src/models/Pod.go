package models

type Pod struct {
	Name       string   `json:"name"`
	Namespace  string   `json:"namespace"`
	Images     string   `json:"images"`
	NodeName   string   `json:"node_name"`
	IP         []string `json:"ip"`       //第一个是POD IP 第二个是node ip
	Phase      string   `json:"phase"`    //状态
	IsReady    bool     `json:"is_ready"` //pod是否就绪
	Message    string   `json:"message"`
	CreateTime string   `json:"create_time"`
}

type PodShellPost struct {
	PodName       string `json:"pod_name"`
	Namespace     string `json:"namespace"`
	ContainerName string `json:"container_name"`
}

type PodDetail struct {
	Name        string       `json:"name"`
	Namespace   string       `json:"namespace"`
	Images      string       `json:"images"`
	NodeName    string       `json:"node_name"`
	IP          []string     `json:"ip"`       //第一个是POD IP 第二个是node ip
	Phase       string       `json:"phase"`    //状态
	IsReady     bool         `json:"is_ready"` //pod是否就绪
	Message     string       `json:"message"`
	CreateTime  string       `json:"create_time"`
	Labels      []string     `json:"labels"`
	Annotations []string     `json:"annotations"`
	Containers  []*Container `json:"containers"`
}

type PodPost struct {
	Name        string            `json:"name" binding:"required"`
	Namespace   string            `json:"namespace" binding:"required"`
	Labels      map[string]string `json:"Labels"`
	Annotations map[string]string `json:"annotations"`
	Containers  []*ContainerPost  `json:"containers" binding:"required"`
}
