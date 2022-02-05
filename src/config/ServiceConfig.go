package config

import (
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
)

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (s *ServiceConfig) JdInitCommonService() *services.CommonService {
	return services.NewCommonService()
}

func (s *ServiceConfig) JdInitDeploymentService() *services.DeploymentService {
	return services.NewDeploymentService()
}

func (s *ServiceConfig) JdInitPodService() *services.PodService {
	return services.NewPodService()
}

func (s *ServiceConfig) JdInitNamespaceService() *services.NamespaceService {
	return services.NewNamespaceService()
}

func (s *ServiceConfig) JdInitServiceService() *services.ServiceService {
	return services.NewServiceService()
}

func (s *ServiceConfig) JdInitIngressService() *services.IngressService {
	return services.NewIngressService()
}
