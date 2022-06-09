package framework

import "net/http"

type Core struct {
}

func NewCode() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}
