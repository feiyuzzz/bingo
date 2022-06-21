package main

import (
	"github.com/feiyuzzz/bingo/framework"
	"log"
)

func registerRouter(core *framework.Core) {
	log.Println("===== register =====")
	core.Get("/user/login", UserLoginController)
	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}

}
