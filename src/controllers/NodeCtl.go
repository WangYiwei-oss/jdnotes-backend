package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/wscore"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"log"
)

type NodeCtl struct {
	K8sClient        *kubernetes.Clientset      `inject:"-"`
	NodeService      *services.NodeService      `inject:"-"`
}

func NewNodeCtl() *NodeCtl {
	return &NodeCtl{}
}


func (n *NodeCtl) GetList(ctx *gin.Context) (int, jdft.Json) {
	list := n.NodeService.List()
	return 1, list
}

func (n *NodeCtl) _readMessageCallback(client *wscore.WsClient, messageType int, message []byte) {
	m := wsMessage{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		log.Println("NodeCtl: ", err)
	}
	if m.Url == "nodes" {
		client.SendMessage(n.NodeService.List())
	}
}

func (n *NodeCtl) _sendStrategy(labelmap wscore.WsClientLabel, vs ...interface{}) bool {
	return true
}

func (n *NodeCtl) WebSocketConn(ctx *gin.Context) int {
	client, err := wscore.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	fmt.Println("okkkk")
	if err != nil {
		return -103
	} else {
		jdft.WebSocketFactory.Store("nodes", client, make(map[string]string), n._sendStrategy, n._readMessageCallback)
		return 1
	}
}

func (n *NodeCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "nodes", n.GetList)
	jdft.Handle("GET", "nodes_ws", n.WebSocketConn)
}
