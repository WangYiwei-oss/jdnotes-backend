package models

type Pod struct {
	Name       string
	Namespace  string
	Images     string
	NodeName   string
	IP         []string //第一个是POD IP 第二个是node ip
	Phase      string   //状态
	IsReady    bool     //pod是否就绪
	Message    string
	CreateTime string
}
