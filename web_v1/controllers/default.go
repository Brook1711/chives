package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "chives.me"
	c.Data["Email"] = "brook1711@bupt.edu.cn"
	c.TplName = "dplayer.html"
}
