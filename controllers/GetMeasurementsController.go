package controllers

import (
	"github.com/astaxie/beego"
	"compost/modules"
	"encoding/json"
)

type GetMeasurements struct {
	beego.Controller
}

func (m *GetMeasurements) Get() {
	//measurements := getInitialMeasurements()
	//m.Data["Measurements"] = &measurements
	//m.Ctx.Output.Body([]byte("Got Initial Measurements!!"))
	//beego.Debug("Device Name : ", m.Ctx.Input.Param(":device"))
	measurements, err := modules.GetMeasurementsByName(m.Ctx.Input.Param(":device"))
	if err !=true {
		beego.Error("Error while querying measurements by name")
	}else {
		//beego.Debug(" Got measuremetns by name : \n", measurements)
		meas,e := json.Marshal(measurements)
		if e != nil{
			beego.Error("error while marshaling measurements")
		}else {
			m.Ctx.Output.Body(meas)
		}
	}

}

