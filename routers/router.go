package routers

import (
	"HelloBee/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{},"Get:GetInit")
	beego.Router("/api/doc", &controllers.MainController{},"Get:GetDocs")
}
