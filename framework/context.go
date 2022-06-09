package framework

import (
	"context"
	"net/http"
	"sync"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter

	ctx     context.Context
	handler ControllerHandler

	// 是否超时
	hashTimeout bool
	// 写保护
	writerMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
	}
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeOut() {
	ctx.hashTimeout = true
}

func (ctx *Context) HasTimeOut() bool {
	return ctx.hashTimeout
}
