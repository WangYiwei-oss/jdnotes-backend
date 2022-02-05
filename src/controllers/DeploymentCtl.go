package controllers

import (
	"encoding/json"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/wscore"
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

type wsMessage struct {
	Namespace string `json:"namespace"`
	Url       string `json:"url"`
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

func (d *DeploymentCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "deployments", d.GetList)
	jdft.Handle("GET", "deployments_ws", d.WebSocketConn)
}
