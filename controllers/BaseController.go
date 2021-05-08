package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"runtime"
)

type BaseController struct {
	web.Controller
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
}

type InterfaceArrayResp struct {
	Arr     []*interface{} `json:"arr"`
	HasNext bool           `json:"hasNext"`
}

func SuccessStatus() *Status {
	return &Status{200, ""}
}

func ErrorStatus(msg string) *Status {
	return &Status{400, msg}
}

func (c *BaseController) Prepare() {
	c.EnableRender = false
}

func (c *BaseController) getRequestBody(body interface{}) error {
	bodyStr := string(c.Ctx.Input.RequestBody)
	fmt.Printf(`bodyStr：%s\n`, bodyStr)
	e := json.Unmarshal([]byte(bodyStr), body)
	return e
}

func (c *BaseController) handleResponse(r *Response) {
	c.Data["json"] = r
	err := c.ServeJSON()
	if err != nil {
		logs.Error("handleResponse() error: ->", err)
	}
}

func (c *BaseController) handleJson(json string) {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.WriteString(json)
}

func (c *BaseController) handleErrorResponse(msg string) {
	c.handleResponse(&Response{
		Status: ErrorStatus(msg),
	})
}

func (c *BaseController) handleErrorCodeResponse(code int, msg string) {
	c.handleResponse(&Response{
		Status: &Status{code, msg},
	})
}

func (c *BaseController) handleSuccessResponse(data interface{}) {
	c.handleResponse(&Response{
		Status: SuccessStatus(),
		Data:   data,
	})
}
func (c *BaseController) handleSuccessNilDataResponse() {
	c.handleResponse(&Response{
		Status: SuccessStatus(),
	})
}
func (c *BaseController) handleSuccessStatusMsg(msg string) {
	c.handleResponse(&Response{
		Status: &Status{200, msg},
	})
}

func (c *BaseController) handleSuccessCodeMsg(code int, msg string) {
	c.handleResponse(&Response{
		Status: &Status{code, msg},
	})
}


func MethodTrace() {
	pc := make([]uintptr, 10)  // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}
