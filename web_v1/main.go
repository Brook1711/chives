package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "web_v1/routers"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	//InsertFilter是提供一个过滤函数
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods:   []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: 	[]string{"Content-Type", "Content-Length", "Authorization", "Accept", "X-Requested-With" , "yourHeaderFeild"},
		//公开的HTTP标头列表
		//ExposeHeaders:	[]string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:	[]string{"*"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	beego.Run()
}

