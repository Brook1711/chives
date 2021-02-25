package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "web_v1/routers"
)

func main() {
	beego.Run()
}

