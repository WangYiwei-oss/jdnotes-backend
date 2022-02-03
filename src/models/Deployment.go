package models

type Deployment struct {
	Name       string   `json:"name"`
	Namespace  string   `json:"namespace"`
	Replicas   [3]int32 `json:"replicas"` //总副本数，可用副本数，不可用副本数
	Images     string   `json:"images"`
	IsComplete bool     `json:"is_complete"` //是否完成
	Message    string   `json:"message"`     //显示错误信息
	CreateTime string   `json:"create_time"`
	Pods       []*Pod   `json:"pods"`
}
