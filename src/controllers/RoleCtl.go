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

type RoleCtl struct {
	K8sClient   *kubernetes.Clientset `inject:"-"`
	RoleService *services.RoleService `inject:"-"`
}

func NewRoleCtl() *RoleCtl {
	return &RoleCtl{}
}

func (r *RoleCtl) GetList(c *gin.Context) (int, jdft.Json) {
	namespace := c.DefaultQuery("namespace", "default")
	list := r.RoleService.ListNamespace(namespace)
	return 1, list
}

func (r *RoleCtl) _readMessageCallback(client *wscore.WsClient, messageType int, message []byte) {
	m := wsMessage{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		log.Println("RoleCtl: ", err)
	}
	if m.Url == "roles" {
		client.Label["namespace"] = m.Namespace
		client.SendMessage(r.RoleService.ListNamespace(m.Namespace))
	}
}

func (r *RoleCtl) _sendStrategy(labelmap wscore.WsClientLabel, vs ...interface{}) bool {
	if labelmap["namespace"] == vs[0].(string) { //一样说明用户正在看当前发生改变的namespace，所以需要通知
		return true
	}
	return false
}

func (r *RoleCtl) WebSocketConn(c *gin.Context) int {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return -103
	} else {
		jdft.WebSocketFactory.Store("roles", client, make(map[string]string), r._sendStrategy, r._readMessageCallback)
		return 1
	}
}

func (r *RoleCtl) DeleteSecret(c *gin.Context) int {
	err := r.RoleService.DelSecret(c.Query("name"), c.Query("namespace"))
	if err != nil {
		return -400
	}
	return 1
}

func (r *RoleCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "roles", r.GetList)
	jdft.Handle("GET", "roles_ws", r.WebSocketConn)
	jdft.Handle("DELETE", "role", r.DeleteSecret)
}
