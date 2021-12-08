package main

import (
	"fmt"
	"net/http"
)

func main() {
	server:= NewSdkHttpServer("lee")
	server.Route("POST", "/user", SignUpWithContext)
	err := server.Start(":8080")
	if err != nil {
		fmt.Printf("err is %v \n", err)
	}
}

// sdkHttpServer 基于 net/http 实现 http server
type sdkHttpServer struct {
	// Name server 的名字
	Name string
	// handlers 接口实现基础功能
	handlers Handler
	// root 根节点
	root Filter
}
type Server interface {
	Routable
	// Start 启动服务器
	Start(address string) error
}

// Route 修改路由方法，支持 map
func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(c *Context)) {
	s.handlers.Route(method, pattern, handlerFunc)
}

// Start 传入 handler
func (s *sdkHttpServer) Start(address string) error {
	// Filter 绑定到路由
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

// NewSdkHttpServer 返回类型 Server
func NewSdkHttpServer(name string, buliders ...FilterBuilder) Server {
	handler := NewHandlerBaseOnMap()
	var root = func(c *Context) {
		handler.ServeHTTP(c)
	}
	// root 原本是最内层的，包装到最外层
	for i := len(buliders) -1; i >= 0; i-- {
		b := buliders[i]
		root = b(root)
	}
	return &sdkHttpServer{
		Name: name,
		handlers: handler,
		root: root,
	}
}