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

type SecretCtl struct {
	K8sClient     *kubernetes.Clientset   `inject:"-"`
	SecretService *services.SecretService `inject:"-"`
}

func NewSecretCtl() *SecretCtl {
	return &SecretCtl{}
}

func (s *SecretCtl) GetList(c *gin.Context) (int, jdft.Json) {
	namespace := c.DefaultQuery("namespace", "default")
	list := s.SecretService.ListNamespace(namespace)
	return 1, list
}

func (s *SecretCtl) _readMessageCallback(client *wscore.WsClient, messageType int, message []byte) {
	m := wsMessage{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		log.Println("DeploymentsCtl: ", err)
	}
	if m.Url == "secrets" {
		client.Label["namespace"] = m.Namespace
		client.SendMessage(s.SecretService.ListNamespace(m.Namespace))
	}
}

func (s *SecretCtl) _sendStrategy(labelmap wscore.WsClientLabel, vs ...interface{}) bool {
	if labelmap["namespace"] == vs[0].(string) { //一样说明用户正在看当前发生改变的namespace，所以需要通知
		return true
	}
	return false
}

func (s *SecretCtl) WebSocketConn(c *gin.Context) int {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return -103
	} else {
		jdft.WebSocketFactory.Store("secrets", client, make(map[string]string), s._sendStrategy, s._readMessageCallback)
		return 1
	}
}

func (s *SecretCtl) CreateSecret(c *gin.Context) int {
	post := &models.SecretPost{}
	err := c.ShouldBind(post)
	if err != nil {
		return -401
	}
	err = s.SecretService.CreateSecret(post)
	if err != nil {
		return -400
	}
	return 1
}

func (s *SecretCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "secrets", s.GetList)
	jdft.Handle("GET", "secrets_ws", s.WebSocketConn)
}
