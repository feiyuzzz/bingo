package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	// 双层map，一级map存放http请求方式，二级map存放对用的Handler
	router map[string]*Tree
}

func NewCore() *Core {
	// init router
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

func (c *Core) FindRouteByRequest(req *http.Request) ControllerHandler {
	uri := req.URL.Path
	method := req.Method
	upperMethod := strings.ToUpper(method)
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) POST(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) PUT(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) DELETE(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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
}
