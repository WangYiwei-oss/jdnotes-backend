package models

type Container struct {
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	Command    []string `json:"command"`
	Args       []string `json:"args"`
	WorkingDir string   `json:"working_dir"`
	Mounts     string   `json:"mounts"`
	Ports      string   `json:"ports"`
}

type ContainerPost struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image" binding:"required"`
}
