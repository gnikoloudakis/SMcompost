package controllers

import (
	"github.com/astaxie/beego"
	"compost/modules"
)

type DashboardController struct {
	beego.Controller
}

func (d *DashboardController) Get(){
	devices, _ := modules.GetDevices()
	d.Data["Devices"] = &devices
	d.TplName = "dashboard.html"
}