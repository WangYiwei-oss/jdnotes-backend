package models

type NodeCapacity struct{
	Cpu int64
	Memory int64
	Pods int64
}

type NodeUsage struct{
	Cpu int64
	Memory int64
	Pods int64
}

type Node struct {
	Name string	`json:"name"`
	IP string	`json:"ip"`
	HostName string	`json:"hostname"`
	CreateTime string	`json:"create_time"`
	Phase string	`json:"phase"`
	Labels []string    `json:"labels"`
	Taints []string		`json:"taints"`
	Capacity *NodeCapacity	`json:"capacity"`
	Usage *NodeUsage 	`json:"usage"`
}
