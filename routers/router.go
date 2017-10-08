package routers

import (
	"snake/controllers"
	"github.com/astaxie/beego"
)
func init() {
	// beego 跨域设置
	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowAllOrigins:  true,
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	AllowCredentials: true,
	//}))

	beego.Router("/w", &controllers.WebSocketController{}, "get:View")
	beego.Router("/ws", &controllers.WebSocketController{}, "get:WS")

}
