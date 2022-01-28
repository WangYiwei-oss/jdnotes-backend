package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	var tlsConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	var tansport http.RoundTripper = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
		DisableCompression:    true,
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		server, _ := url.Parse("https://192.168.83.150:6443")
		p := httputil.NewSingleHostReverseProxy(server)
		p.Transport = tansport
		p.ServeHTTP(writer, request)
	})
	log.Println("开始代理 http://0.0.0.0:9090 -> https://192.168.83.150:6443")
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		log.Fatalln("代理服务器启动错误", err)
	}
}
