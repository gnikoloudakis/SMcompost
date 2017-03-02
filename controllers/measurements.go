package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"github.com/astaxie/beego"
	"compost/models"
)

type MeasurementsController struct {
	beego.Controller
}

type Measurement struct {
	Arduino     string        `json:"arduino"`
	IP          string        `json:"ip"`
	Temperature float32        `json:"temperature"`
}

func (r *MeasurementsController) Post() {

	p := make([]byte, r.Ctx.Request.ContentLength)
	_, err := r.Ctx.Request.Body.Read(p)
	fmt.Println("err", err)
	if err == nil {
		var measurement Measurement
		err1 := json.Unmarshal(p, &measurement)
		if err1 == nil {
			//fmt.Println("Arduino:", measurement.Arduino)
			//fmt.Println("IP:", measurement.IP)
			//fmt.Println("Temperature:", measurement.Temperature)
			insertMeasurement(measurement)

		} else {
			fmt.Println("Unable to unmarshall the JSON request", err1);
		}
	}
	r.Ctx.Output.Body([]byte("Measurements POST went OK!!"))
}

func getArduino(arduinoName string) (arduino models.Arduino) {
	o := orm.NewOrm()
	arduino = models.Arduino{Name:arduinoName}
	err := o.Read(&arduino, "Name")
	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println("got Arduino")
	}
	return arduino
}

func insertMeasurement(measurement Measurement) {
	fmt.Println("Arduino: ", measurement.Arduino, "Temperature: ", measurement.Temperature)
	arduinoInstance := getArduino(measurement.Arduino)
	o := orm.NewOrm()
	o.Using("compost")
	newMeasurement := models.Measurements{Tempereature:measurement.Temperature, Arduino: &arduinoInstance}
	id1, err1 := o.Insert(&newMeasurement)
	if err1 == nil {
		fmt.Println("Inserted measurement with ID: ", id1)

	} else {
		fmt.Println("Error inserting measurement - Error: ", err1)
	}
}