package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"github.com/astaxie/beego"
	"compost/models"
	"compost/modules"
)

type AddMeasurementsController struct {
	beego.Controller
}

type Measurement struct {
	Device	    string        `json:"Device"`
	Temperature float32       `json:"Temperature"`
}

func (r *AddMeasurementsController) Post() {

	p := make([]byte, r.Ctx.Request.ContentLength)
	_, err := r.Ctx.Request.Body.Read(p)
	if err == nil {
		var measurement Measurement
		err1 := json.Unmarshal(p, &measurement)
		if err1 == nil {
			//fmt.Println("Arduino:", measurement.Arduino)
			//fmt.Println("IP:", measurement.IP)
			//fmt.Println("Temperature:", measurement.Temperature)
			insertMeasurement(measurement)

		} else {
			beego.Error("Unable to unmarshall the JSON request", err1);
		}
	}else{
		beego.Error("error in reading request :", err)
	}
	r.Ctx.Output.Body([]byte("Measurements POST went OK!!"))
}



func insertMeasurement(measurement Measurement) {
	fmt.Println("Device: ", measurement.Device, "Temperature: ", measurement.Temperature)
	deviceInstance, success := modules.GetDeviceByName(measurement.Device)

	if success {
		o := orm.NewOrm()
		o.Using("compost")
		newMeasurement := models.Measurements{Temperature:measurement.Temperature, Device: &deviceInstance}
		id1, err1 := o.Insert(&newMeasurement)
		if err1 == nil {
			beego.Debug("Inserted measurement with ID: ", id1)

		} else {
			beego.Error("Error inserting measurement - Error: ", err1)
		}
	}else {
		beego.Error("Did NOT find Device before adding measurement")
	}
}