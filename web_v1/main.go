package main

import (
	_ "web_v1/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

