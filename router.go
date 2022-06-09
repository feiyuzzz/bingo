package main

import "github.com/feiyuzzz/bingo/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
