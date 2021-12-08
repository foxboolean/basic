package main

import (
	"fmt"
	"net/http"
)

type Routable interface {
	Route(method string, pattern string, handlerFunc func(c *Context))
}
type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

type HandlerBaseOnMap struct {
	handlers map[string]func(c *Context)
}

// Route 修改路由方法，支持 map
func (h HandlerBaseOnMap) Route(method string, pattern string, handlerFunc func(c *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc
}
func (h HandlerBaseOnMap) ServeHTTP(c *Context) {
	request, writer := c.R, c.W
	// 分发路由
	key := h.key(request.Method, request.URL.Path)

	// 匹配到路由处理
	if handler, ok := h.handlers[key]; ok {
		c := NewContext(writer, request)
		handler(c)
	} else {
		// 404
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte("not any router mathc"))
	}
}

func (h *HandlerBaseOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}
var _ Handler = NewHandlerBaseOnMap()

func NewHandlerBaseOnMap() Handler {
	return HandlerBaseOnMap{
		handlers: make(map[string]func(c *Context),128),
	}
}