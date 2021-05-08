package controllers

import (
	"github.com/beego/beego/v2/core/logs"
)

type JO struct {
	Msg     string `json:"msg"`
	Name string `json:"name"`
	Desc       string `json:"desc"`
}

type MainController struct {
	//beego.Controller
    BaseController
}

func (c *MainController) GetDocs() {

	logs.Info("/api/doc")

	MethodTrace()

	jo := &JO{
		Msg:  "Hello Get Docs Invoked.",
		Name: "This is Name.",
		Desc: "This Is Desc.",
	}
	c.handleSuccessResponse(jo)
}

func (c *MainController) GetInit() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.handleSuccessCodeMsg(200,"")
}




