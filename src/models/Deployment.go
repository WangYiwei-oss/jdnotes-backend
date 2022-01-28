package models

type Deployment struct {
	Name       string
	Namespace  string
	Replicas   [3]int32 //总副本数，可用副本数，不可用副本数
	Images     string
	IsComplete bool   //是否完成
	Message    string //显示错误信息
	CreateTime string
	Pods       []*Pod
}
