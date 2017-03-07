package controllers

import (
	"github.com/astaxie/beego"
)

type testController struct {
	beego.Controller
}

func (c *testController) Get() {
	c.TplName = "index.tpl"

}
