package controllers

import (
	"encoding/json"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/wscore"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
	"log"
)

type IngressCtl struct {
	IngressService *services.IngressService `inject:"-"`
}

func NewIngressCtl() *IngressCtl {
	return &IngressCtl{}
}

func (i *IngressCtl) GetList(c *gin.Context) (int, jdft.Json) {
	namespace := c.DefaultQuery("namespace", "default")
	list := i.IngressService.ListNamespace(namespace)
	return 1, list
}

func (i *IngressCtl) _readMessageCallback(client *wscore.WsClient, messageType int, message []byte) {
	m := wsMessage{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		log.Println("IngressCtl: ", err)
	}
	if m.Url == "ingresses" {
		client.Label["namespace"] = m.Namespace
		client.SendMessage(i.IngressService.ListNamespace(m.Namespace))
	}
}

func (i *IngressCtl) _sendStrategy(labelmap wscore.WsClientLabel, vs ...interface{}) bool {
	if labelmap["namespace"] == vs[0].(string) { //一样说明用户正在看当前发生改变的namespace，所以需要通知
		return true
	}
	return false
}

func (i *IngressCtl) WebSocketConn(c *gin.Context) int {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return -103
	} else {
		jdft.WebSocketFactory.Store("ingresses", client, make(map[string]string), i._sendStrategy, i._readMessageCallback)
		return 1
	}
}

// CreateIngress 创建ingress接口
func (i *IngressCtl) CreateIngress(c *gin.Context) (int, string) {
	postModel := &models.IngressPost{}
	err := c.BindJSON(postModel)
	if err != nil {
		return -400, err.Error()
	}
	err = i.IngressService.CreateIngress(postModel)
	if err != nil {
		return -400, err.Error()
	}
	return 1, ""
}

func (i *IngressCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "ingresses", i.GetList)
	jdft.Handle("GET", "ingresses_ws", i.WebSocketConn)
	jdft.Handle("POST", "ingress", i.CreateIngress)
}
