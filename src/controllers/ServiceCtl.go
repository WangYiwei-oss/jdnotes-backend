package controllers

import (
	"encoding/json"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/wscore"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
	"log"
)

type ServiceCtl struct {
	ServiceService *services.ServiceService `inject:"-"`
}

func NewServiceCtl() *ServiceCtl {
	return &ServiceCtl{}
}

func (s *ServiceCtl) GetList(c *gin.Context) (int, jdft.Json) {
	namespace := c.DefaultQuery("namespace", "default")
	list := s.ServiceService.ListNamespace(namespace)
	return 1, list
}

func (s *ServiceCtl) _readMessageCallback(client *wscore.WsClient, messageType int, message []byte) {
	m := wsMessage{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		log.Println("ServiceCtl: ", err)
	}
	if m.Url == "services" {
		client.Label["namespace"] = m.Namespace
		client.SendMessage(s.ServiceService.ListNamespace(m.Namespace))
	}
}

func (s *ServiceCtl) _sendStrategy(labelmap wscore.WsClientLabel, vs ...interface{}) bool {
	if labelmap["namespace"] == vs[0].(string) { //一样说明用户正在看当前发生改变的namespace，所以需要通知
		return true
	}
	return false
}

func (s *ServiceCtl) WebSocketConn(c *gin.Context) int {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return -103
	} else {
		jdft.WebSocketFactory.Store("pods", client, make(map[string]string), s._sendStrategy, s._readMessageCallback)
		return 1
	}
}

func (s *ServiceCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "services", s.GetList)
	jdft.Handle("GET", "services_ws", s.WebSocketConn)
}
