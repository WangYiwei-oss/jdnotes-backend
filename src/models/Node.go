package models

type NodeCapacity struct {
	Cpu    int64
	Memory int64
	Pods   int64
}

func NewNodeCapacity(cpu int64, memory int64, pods int64) *NodeCapacity {
	return &NodeCapacity{Cpu: cpu, Memory: memory, Pods: pods}
}

type NodeUsage struct {
	Cpu    float64
	Memory float64
	Pods   int
}

func NewNodeUsage(cpu float64, memory float64, pods int) *NodeUsage {
	return &NodeUsage{Cpu: cpu, Memory: memory, Pods: pods}
}

type Node struct {
	Name       string `json:"name"`
	IP         string `json:"ip"`
	HostName   string `json:"hostname"`
	Phase      string `json:"phase"`
	CreateTime string `json:"create_time"`
}

type NodeDetail struct {
	Name       string        `json:"name"`
	IP         string        `json:"ip"`
	HostName   string        `json:"hostname"`
	CreateTime string        `json:"create_time"`
	Phase      string        `json:"phase"`
	Labels     []string      `json:"labels"`
	Taints     []string      `json:"taints"`
	Capacity   *NodeCapacity `json:"capacity"`
	Usage      *NodeUsage    `json:"usage"`
}

type NodeShellPost struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}
