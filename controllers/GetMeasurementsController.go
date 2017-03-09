package controllers

import (
	"github.com/astaxie/beego"
	"compost/modules"
	"encoding/json"
	_ "fmt"
	_ "time"
)

type GetMeasurements struct {
	beego.Controller
}

func (m *GetMeasurements) GetByName() {

	//beego.Debug("Device Name : ", m.Ctx.Input.Param(":device"))
	//beego.Debug("Start Date : ", m.Ctx.Input.Param(":start"))
	//beego.Debug("Stop Date : ", m.Ctx.Input.Param(":stop"))
	measurements, err := modules.GetMeasurementsByName(m.Ctx.Input.Param(":device"))
	//qq := make(map[time.Time]float64)
	//for _, element := range measurements {
	//	qq[element["Timestamp"].(time.Time)] = element["Temperature"].(float64)
	//}
	//beego.Debug(qq)
	if err != true {
		beego.Error("Error while querying measurements by name")
	} else {
		//beego.Debug(" Got measuremetns by name : \n", measurements)
		meas, e := json.Marshal(measurements)
		if e != nil {
			beego.Error("error while marshaling measurements")
		} else {
			//beego.Debug(meas)
			m.Ctx.Output.Body(meas)
		}
	}
}

func (m *GetMeasurements) GetByDate() {
	measurements, success := modules.GetMeasurementsByDate(m.Ctx.Input.Param(":device"), m.Ctx.Input.Param(":start"), m.Ctx.Input.Param(":stop"))
	//fmt.Println(measurements[2])
	if success != true {
		beego.Error("Error while querying measurements by Date")
	} else {
		meas, e := json.Marshal(measurements)
		if e != nil {
			beego.Error("error while marshaling measurements")
		} else {
			m.Ctx.Output.Body(meas)
		}
	}
}

func (m *GetMeasurements) GetLatestData(){
	measurement, success := modules.GetLatestMeasurement(m.Ctx.Input.Param(":device"))
	if success != true {
		beego.Error("Error while querying Latest measurement")
	} else {
		meas, e := json.Marshal(measurement)
		if e != nil {
			beego.Error("error while marshaling Latest measurement")
		} else {
			m.Ctx.Output.Body(meas)
		}
	}
	m.Ctx.Output.Body([]byte("ok"))
}