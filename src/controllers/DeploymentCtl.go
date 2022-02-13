package controllers

import (
	"encoding/json"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/wscore"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"log"
)

type DeploymentCtl struct {
	K8sClient  *kubernetes.Clientset       `inject:"-"`
	DepService *services.DeploymentService `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (d *DeploymentCtl) GetList(c *gin.Context) (int, jdft.Json) {
	namespace := c.DefaultQuery("namespace", "default")
	list := d.DepService.ListNamespace(namespace)
	return 1, list
}

func (d *DeploymentCtl) _readMessageCallback(client *wscore.WsClient, messageType int, message []byte) {
	m := wsMessage{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		log.Println("DeploymentsCtl: ", err)
	}
	if m.Url == "deployments" {
		client.Label["namespace"] = m.Namespace
		client.SendMessage(d.DepService.ListNamespace(m.Namespace))
	}
}

func (d *DeploymentCtl) _sendStrategy(labelmap wscore.WsClientLabel, vs ...interface{}) bool {
	if labelmap["namespace"] == vs[0].(string) { //一样说明用户正在看当前发生改变的namespace，所以需要通知
		return true
	}
	return false
}

func (d *DeploymentCtl) WebSocketConn(c *gin.Context) int {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return -103
	} else {
		jdft.WebSocketFactory.Store("deployments", client, make(map[string]string), d._sendStrategy, d._readMessageCallback)
		return 1
	}
}

// DeleteDeployment 删除deployment
func (d *DeploymentCtl) DeleteDeployment(c *gin.Context) int {
	err := d.DepService.DeleteDeployment(c.Query("name"), c.Query("namespace"))
	if err != nil {
		return -400
	}
	return 1
}

// GetDeploymentDetail 获取deployment详情
func (d *DeploymentCtl) GetDeploymentDetail(c *gin.Context) (int, jdft.Json) {
	detail, err := d.DepService.GetDeploymentDetail(c.Query("name"), c.Query("namespace"))
	if err != nil {
		return -400, err.Error()
	}
	return 1, detail
}

func (d *DeploymentCtl) CreateDeployment(c *gin.Context) (int, string) {
	postModel := &models.DeploymentPost{}
	err := c.BindJSON(postModel)
	if err != nil {
		return -400, err.Error()
	}
	err = d.DepService.CreateDeployment(postModel)
	if err != nil {
		return -400, err.Error()
	}
	return 1, ""
}

func (d *DeploymentCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "deployments", d.GetList)
	jdft.Handle("GET", "deployment", d.GetDeploymentDetail)
	jdft.Handle("GET", "deployments_ws", d.WebSocketConn)
	jdft.Handle("DELETE", "deployment", d.DeleteDeployment)
	jdft.Handle("POST", "deployment", d.CreateDeployment)
}
