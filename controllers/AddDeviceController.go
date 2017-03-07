package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"compost/modules"
	"compost/models"
)

type AddDeviceController struct {
	beego.Controller
}

type Device struct {
	Name string        `json:"Name"`
	IP   string        `json:"IP"`
}

func (a *AddDeviceController) Post() {
	p := make([]byte, a.Ctx.Request.ContentLength)
	_, err := a.Ctx.Request.Body.Read(p)
	if err == nil {
		var device Device
		err1 := json.Unmarshal(p, &device)
		if err1 == nil {
			//fmt.Println("Arduino:", measurement.Arduino)
			//fmt.Println("IP:", measurement.IP)
			//fmt.Println("Temperature:", measurement.Temperature)
			success := insertDevice(device)
			if success {
				a.Ctx.Output.Body([]byte("Inserted NEW Device!!"))
			}else {
				a.Ctx.Output.Body([]byte("Device NOT inserted!!"))
			}
		} else {
			beego.Error("Unable to unmarshall the JSON request in insert device \n", err1);
		}
	} else {
		beego.Error("error while reading insert device request \n", err)
	}

}

func insertDevice(device Device) (success bool) {
	success = false
	if modules.FilterDeviceName(device.Name) {
		fmt.Println("Name: ", device.Name, "IP: ", device.IP)
		o := orm.NewOrm()
		o.Using("compost")
		newDevice := models.Devices{Name:device.Name, IP: device.IP}
		id, err := o.Insert(&newDevice)
		if err == nil {
			beego.Debug("Inserted Device with ID: ", id)
			success = true

		} else {
			beego.Error("Error inserting Device - Error: ", err)
		}
	}else {
		beego.Error("Device Name contains illegal characters")
	}
	return success
}