package main

import (
	"github.com/feiyuzzz/bingo/framework"
	"log"
)

func UserLoginController(c *framework.Context) error {
	// 打印控制器名字
	log.Println("===== UserLoginController =====")
	c.Json(200, "ok, UserLoginController")
	return nil
}
