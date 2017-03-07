package modules

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"compost/models"
)

func GetDeviceByName(deviceName string) (device models.Devices, success bool) {
	o := orm.NewOrm()
	success= false
	device = models.Devices{Name:deviceName}
	err := o.Read(&device, "Name")
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
	} else {
		beego.Debug("Found Device for Inserting NEW measurement")
		success = true
	}
	return device, success
}

func GetMeasurementsByName(device string)(meas []orm.Params, success bool){
	success = false
	deviceInstance, res := GetDeviceByName(device)
	if res {
		o := orm.NewOrm()
		num, err := o.QueryTable("measurements").Filter("device_id", deviceInstance.Id).OrderBy("timestamp").Values(&meas)
		if err != nil {
			beego.Error("error in getting Initial Measurements :", err)
		} else {
			beego.Debug("got Initial Measurements Number of results:", num)
			success = true
		}
	}else {
		beego.Error("GetMeasurementsByName - not find device")
	}
	return meas, success
}