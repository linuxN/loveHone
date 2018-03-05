package routers

import (
	"github.com/astaxie/beego"
	"loveHome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetAreaInfo")
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionName;delete:DelSessionName")
	beego.Router("/api/v1.0/houses/index", &controllers.HousesIndexController{}, "get:GetHousesIndex")

	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:Rec")
	beego.Router("/api/v1.0/sessions", &controllers.UserController{}, "post:Login")
}
