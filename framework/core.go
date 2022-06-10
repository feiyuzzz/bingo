package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	// 双层map，一级map存放http请求方式，二级map存放对用的Handler
	router map[string]map[string]ControllerHandler
}

func NewCore() *Core {

	// 定义二级 map
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}

	// 二级 map 存入一级 map
	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter
	return &Core{
		router: router,
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	upperURL := strings.ToUpper(url)
	c.router["GET"][upperURL] = handler
}

func (c *Core) POST(url string, handler ControllerHandler) {
	upperURL := strings.ToUpper(url)
	c.router["POST"][upperURL] = handler
}

func (c *Core) PUT(url string, handler ControllerHandler) {
	upperURL := strings.ToUpper(url)
	c.router["PUT"][upperURL] = handler
}

func (c *Core) DELETE(url string, handler ControllerHandler) {
	upperURL := strings.ToUpper(url)
	c.router["DELETE"][upperURL] = handler
}

func (c *Core) FindRouteByRequest(req *http.Request) ControllerHandler {
	uri := req.URL.Path
	method := req.Method
	upperMethod := strings.ToUpper(method)
	upperURI := strings.ToUpper(uri)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		if handler, ok := methodHandlers[upperURI]; ok {
			return handler
		}
	}
	return nil
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.ServeHTTP")
	//封装定义的context
	ctx := NewContext(request, response)

	//寻找路由
	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(404, "not found")
		return
	}

	//调用路由函数，如果放回err 代表内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
	log.Println("core.ServeHTTP")
}
