package models

type Deployment struct {
	Name       string   `json:"name"`
	Namespace  string   `json:"namespace"`
	Replicas   [3]int32 `json:"replicas"` //总副本数，可用副本数，不可用副本数
	Images     string   `json:"images"`
	IsComplete bool     `json:"is_complete"` //是否完成
	Message    string   `json:"message"`     //显示错误信息
	CreateTime string   `json:"create_time"`
}

type DeploymentDetail struct {
	Name        string   `json:"name"`
	Namespace   string   `json:"namespace"`
	Labels      string   `json:"labels"`
	Annotations string   `json:"annotations"`
	Replicas    [3]int32 `json:"replicas"` //总副本数，可用副本数，不可用副本数
	Images      string   `json:"images"`
	IsComplete  bool     `json:"is_complete"` //是否完成
	Message     string   `json:"message"`     //显示错误信息
	CreateTime  string   `json:"create_time"`
	Pods        []*Pod   `json:"pods"`
}

type DeploymentPost struct {
	Name               string            `json:"name" binding:"required"`
	Namespace          string            `json:"namespace" binding:"required"`
	Labels             map[string]string `json:"labels"`
	Annotations        map[string]string `json:"annotations"`
	Replicas           int32             `json:"replicas" binding:"required"`
	DeploymentTemplate *PodPost          `json:"deployment_template" binding:"required"`
}
