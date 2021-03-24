package routers

import (
	"web_v1/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/v3", &controllers.DanmuController{})
    beego.Router("/", &controllers.MainController{})
}
