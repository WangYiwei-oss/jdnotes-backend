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
