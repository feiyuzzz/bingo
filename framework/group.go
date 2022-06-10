package framework

// IGroup 代表前缀分组
type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

type Group struct {
	core   *Core
	prefix string
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Get(uri, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.POST(uri, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.PUT(uri, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.DELETE(uri, handler)
}

// core中初始化
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
