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

type ConfigMapCtl struct {
	K8sClient        *kubernetes.Clientset      `inject:"-"`
	ConfigMapService *services.ConfigMapService `inject:"-"`
}

func NewConfigMapCtl() *ConfigMapCtl {
	return &ConfigMapCtl{}
}

func (c *ConfigMapCtl) GetList(ctx *gin.Context) (int, jdft.Json) {
	namespace := ctx.DefaultQuery("namespace", "default")
	list := c.ConfigMapService.ListNamespace(namespace)
	return 1, list
}

func (c *ConfigMapCtl) _readMessageCallback(client *wscore.WsClient, messageType int, message []byte) {
	m := wsMessage{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		log.Println("ConfigMapCtl: ", err)
	}
	if m.Url == "configmaps" {
		client.Label["namespace"] = m.Namespace
		client.SendMessage(c.ConfigMapService.ListNamespace(m.Namespace))
	}
}

func (c *ConfigMapCtl) _sendStrategy(labelmap wscore.WsClientLabel, vs ...interface{}) bool {
	if labelmap["namespace"] == vs[0].(string) { //一样说明用户正在看当前发生改变的namespace，所以需要通知
		return true
	}
	return false
}

func (c *ConfigMapCtl) WebSocketConn(ctx *gin.Context) int {
	client, err := wscore.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return -103
	} else {
		jdft.WebSocketFactory.Store("configmaps", client, make(map[string]string), c._sendStrategy, c._readMessageCallback)
		return 1
	}
}

//删除
func (c *ConfigMapCtl) DeleteConfigmap(ctx *gin.Context) int {
	err := c.ConfigMapService.DeleteConfigMap(ctx.Query("name"), ctx.Query("namespace"))
	if err != nil {
		return -400
	}
	return 1
}

//创建
func (c *ConfigMapCtl) CreateConfigMap(ctx *gin.Context) int {
	configPost := &models.ConfigMapPost{}
	err := ctx.BindJSON(configPost)
	if err != nil {
		return -400
	}
	err = c.ConfigMapService.CreateConfigMap(configPost)
	if err != nil {
		log.Println(err)
		return -400
	}
	return 1
}

//更新
func (c *ConfigMapCtl) UpdateConfigMap(ctx *gin.Context) int {
	configPost := &models.ConfigMapPost{}
	err := ctx.BindJSON(configPost)
	if err != nil {
		return -400
	}
	err = c.ConfigMapService.UpdateConfigMap(configPost)
	if err != nil {
		return -400
	}
	return 1
}

func (c *ConfigMapCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "configmaps", c.GetList)
	jdft.Handle("GET", "configmaps_ws", c.WebSocketConn)
	jdft.Handle("DELETE", "configmap", c.DeleteConfigmap)
	jdft.Handle("POST", "configmap", c.CreateConfigMap)
	jdft.Handle("PUT", "configmap", c.UpdateConfigMap)
}
