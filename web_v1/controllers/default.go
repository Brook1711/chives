package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type DanmuRequest struct {
	Token 	string
	Id		string
	Author 	int64
	Time 	float32
	Text 	string
	Color 	int32
	DplayerType int8
}

type MainController struct {
	beego.Controller
}

type DanmuController struct {
	beego.Controller
}

func (c *DanmuController) Post() {
	var data DanmuRequest
	RequestData := c.Ctx.Input.RequestBody
	//fmt.Println(RequestData)
	json.Unmarshal(RequestData, &data)
	//fmt.Println(data)
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *DanmuController) Get() {
	jsoninfo := c.GetString("jsoninfo")
	if jsoninfo == "" {
		c.Ctx.WriteString("jsoninfo is empty")
		return
	}
}

func (c *MainController) Get() {
	c.Data["Website"] = "chives.me"
	c.Data["Email"] = "brook1711@bupt.edu.cn"
	c.TplName = "index.html"
}

