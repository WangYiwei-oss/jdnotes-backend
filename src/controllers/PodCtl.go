package controllers

import (
	"encoding/json"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/wscore"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
	"log"
)

type PodCtl struct {
	PodService *services.PodService `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (p *PodCtl) GetList(c *gin.Context) (int, jdft.Json) {
	namespace := c.DefaultQuery("namespace", "default")
	list := p.PodService.ListNamespace(namespace)
	return 1, list
}

func (p *PodCtl) _readMessageCallback(client *wscore.WsClient, messageType int, message []byte) {
	m := wsMessage{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		log.Println("PodCtl: ", err)
	}
	if m.Url == "pods" {
		client.Label["namespace"] = m.Namespace
		client.SendMessage(p.PodService.ListNamespace(m.Namespace))
	}
}

func (p *PodCtl) _sendStrategy(labelmap wscore.WsClientLabel, vs ...interface{}) bool {
	if labelmap["namespace"] == vs[0].(string) { //一样说明用户正在看当前发生改变的namespace，所以需要通知
		return true
	}
	return false
}

func (p *PodCtl) WebSocketConn(c *gin.Context) int {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return -103
	} else {
		jdft.WebSocketFactory.Store("pods", client, make(map[string]string), p._sendStrategy, p._readMessageCallback)
		return 1
	}
}

func (p *PodCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "pods", p.GetList)
	jdft.Handle("GET", "pods_ws", p.WebSocketConn)
}
