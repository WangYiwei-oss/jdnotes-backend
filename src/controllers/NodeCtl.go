package controllers

import (
	"encoding/json"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/wscore"
	"github.com/WangYiwei-oss/jdnotes-backend/src/models"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"k8s.io/client-go/kubernetes"
	"log"
)

type NodeCtl struct {
	K8sClient   *kubernetes.Clientset `inject:"-"`
	NodeService *services.NodeService `inject:"-"`
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
	if err != nil {
		return -103
	} else {
		jdft.WebSocketFactory.Store("nodes", client, make(map[string]string), n._sendStrategy, n._readMessageCallback)
		return 1
	}
}

func (n *NodeCtl) GetNode(ctx *gin.Context) (int, jdft.Json) {
	node, err := n.NodeService.GetNode(ctx.Query("name"))
	if err != nil {
		log.Println("NodeCtl:", err.Error())
		return -400, err.Error()
	}
	return 1, node
}

func (n *NodeCtl) ShellWsConnect(c *gin.Context) int {
	postData := &models.NodeShellPost{
		Name:     c.Query("name"),
		User:     c.Query("user"),
		Password: c.Query("password"),
		Ip:       c.Query("ip"),
	}
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("PodCtl: 升级失败")
		return -103
	}
	session, err := n.NodeService.ConnectNodeShell(postData)
	if err != nil {
		client.WriteMessage(websocket.TextMessage, []byte("服务器连接失败,用户名或密码错误\n"))
		log.Println("NodeCtl:", err)
		return -400
	}
	defer session.Close()

	shellClient := NewWsShellClient(client)
	session.Stdout = shellClient
	session.Stdin = shellClient
	session.Stderr = shellClient
	nodeShellModes := ssh.TerminalModes{
		ssh.ECHO:          1, //enable echoing
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm-256color", 300, 500, nodeShellModes)
	if err != nil {
		client.WriteMessage(websocket.TextMessage, []byte("服务器连接失败，网络错误\n"))
		log.Println("NodeCtl:", err)
		return -400
	}
	err = session.Run("sh")
	if err != nil {
		client.WriteMessage(websocket.TextMessage, []byte("运行sh错误\n"))
		log.Println("NodeCtl:", err)
		return -400
	}
	client.WriteMessage(websocket.TextMessage, []byte("连接成功\n"))
	return 1
}

func (n *NodeCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "nodes", n.GetList)
	jdft.Handle("GET", "node", n.GetNode)
	jdft.Handle("GET", "nodes_ws", n.WebSocketConn)
	jdft.Handle("GET", "node/exec/ws", n.ShellWsConnect)
}
