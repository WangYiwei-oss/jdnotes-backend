package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/wscore"
	"github.com/WangYiwei-oss/jdnotes-backend/src/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"net/http"
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

func (p *PodCtl) DeletePod(c *gin.Context) int {
	err := p.PodService.DeletePod(c.Query("name"), c.Query("namespace"))
	if err != nil {
		return -400
	}
	return 1
}

func (p *PodCtl) GetContainers(c *gin.Context) (int, jdft.Json) {
	ns, name := c.Query("namespace"), c.Query("name")
	containers, err := p.PodService.GetPodContainer(name, ns)
	if err != nil {
		log.Println("PodCtl: Get container failed")
		return -400, err.Error()
	}
	return 1, containers
}

func (p *PodCtl) GetContainerLog(c *gin.Context) (int, string) {
	ns, name, containerName := c.Query("namespace"), c.Query("name"), c.Query("container_name")
	req := p.PodService.Client.CoreV1().Pods(ns).GetLogs(name, &corev1.PodLogOptions{
		Container: containerName,
	})
	ret := req.Do(c)
	b, err := ret.Raw()
	if err != nil {
		log.Println("PodCtl: Get log error", err)
		return -400, err.Error()
	}
	return 1, string(b)
}

func (p *PodCtl) GetContainerLogStream(c *gin.Context) int {
	fmt.Println("获取日志")
	ns, name, containerName := c.Query("namespace"), c.Query("name"), c.Query("container_name")
	req := p.PodService.Client.CoreV1().Pods(ns).GetLogs(name, &corev1.PodLogOptions{
		Container: containerName,
		Follow:    true,
	})
	reader, err := req.Stream(c)
	if err != nil {
		log.Println("PodCtl: Get log error", err)
		return -400
	}
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Writer.Header().Set("Content-Type", "text/html")
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		if n > 0 {
			fmt.Println("读了", string(buf[0:n]))
			c.Writer.Write([]byte(string(buf[0:n])))
			c.Writer.(http.Flusher).Flush()
		}
	}
	return 1
}

//func (p *PodCtl)ExecWsConnect(c *gin.Context)int{
//	client,err := wscore.Upgrader.Upgrade(c.Writer,c.Request,nil)
//	if err != nil {
//		return -103
//	}else{
//		jdft.WebSocketFactory.Store("pod_shell",client,
//			make(map[string]string), nil,
//			func(client *wscore.WsClient, i int, bytes []byte) {})
//	}
//	return 1
//}

type WsShellClient struct {
	client *websocket.Conn
}

func NewWsShellClient(client *websocket.Conn) *WsShellClient {
	return &WsShellClient{client: client}
}

func (w *WsShellClient) Write(p []byte) (n int, err error) {
	err = w.client.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (w *WsShellClient) Read(p []byte) (n int, err error) {
	_, b, err := w.client.ReadMessage()
	if err != nil {
		return 0, err
	}
	return copy(p, string(b)), nil
}

func (p *PodCtl) ExecWsConnect(c *gin.Context) int {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("PodCtl: 升级失败")
		return -103
	}
	client.WriteMessage(websocket.TextMessage, []byte("连接成功\n"))
	shellClient := NewWsShellClient(client)
	excutor, err := p.PodService.HandleCommand(p.PodService.Client, p.PodService.K8sRestConfig, []string{"sh"})
	if err != nil {
		log.Println("PodCtl: 创建Excutor失败")
		return -103
	}
	err = excutor.Stream(remotecommand.StreamOptions{
		Stdin:  shellClient,
		Stdout: shellClient,
		Stderr: shellClient,
		Tty:    true,
	})
	if err != nil {
		log.Println("PodCtl: 创建流失败")
		return -103
	}
	return 1
}

func (p *PodCtl) GetPodDetail(c *gin.Context) (int, jdft.Json) {
	detail, err := p.PodService.GetPodDetail(c.Query("name"), c.Query("namespace"))
	if err != nil {
		return -400, err.Error()
	}
	return 1, detail
}

func (p *PodCtl) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "pods", p.GetList)
	jdft.Handle("GET", "pods_ws", p.WebSocketConn)
	jdft.Handle("GET", "pod", p.GetPodDetail)
	jdft.Handle("DELETE", "pod", p.DeletePod)
	jdft.Handle("GET", "pod/containers", p.GetContainers)
	jdft.Handle("GET", "pod/container/log", p.GetContainerLogStream)
	jdft.Handle("GET", "pod/exec/ws", p.ExecWsConnect)
}
