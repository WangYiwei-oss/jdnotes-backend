package services

import (
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"k8s.io/client-go/kubernetes"
	"log"
	"strings"
)

type ResourceService struct {
	Client *kubernetes.Clientset `inject:"-"`
}

func NewResourceService() *ResourceService {
	return &ResourceService{}
}

func (r *ResourceService) getGroupVersion(str string) (group, version string) {
	list := strings.Split(str, "/")
	if len(list) == 1 {
		return "core", list[0]
	} else if len(list) == 2 {
		return list[0], list[1]
	}
	log.Fatalln("error GroupVersion")
	return
}

func (r *ResourceService) ListResources() ([]*models.GroupResources, error) {
	_, resources, err := r.Client.ServerGroupsAndResources()
	if err != nil {
		return nil, err
	}
	gRes := make([]*models.GroupResources, 0)
	for _, resource := range resources {
		group, version := r.getGroupVersion(resource.GroupVersion)
		gr := &models.GroupResources{
			Group:     group,
			Version:   version,
			Resources: make([]*models.Resources, 0),
		}
		for _, rr := range resource.APIResources {
			res := &models.Resources{
				Name:  rr.Name,
				Verbs: rr.Verbs,
			}
			gr.Resources = append(gr.Resources, res)
		}
		gRes = append(gRes, gr)
	}
	return gRes, nil
}
