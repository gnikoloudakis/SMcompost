package modules

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"compost/models"
	"time"
)

func GetDevices() (maps []orm.Params, success bool) {
	var o orm.Ormer
	o = orm.NewOrm()
	success = false
	num, err := o.QueryTable("devices").Values(&maps)
	if err == nil {
		beego.Debug("Uccessfully got Devices in getDevices num: ", num)
		success = true
	} else {
		beego.Error("Did NOT get devices in getDevices")
	}
	return maps, success
}

func GetDeviceByName(deviceName string) (device models.Devices, success bool) {
	var o orm.Ormer
	o = orm.NewOrm()
	success = false
	device = models.Devices{Name:deviceName}
	err := o.Read(&device, "Name")
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
	} else {
		beego.Debug("Found Device")
		success = true
	}
	return device, success
}

func GetMeasurementsByName(device string) (meas []orm.Params, success bool) {
	success = false
	deviceInstance, res := GetDeviceByName(device)
	if res {
		var o orm.Ormer
		o = orm.NewOrm()
		num, err := o.QueryTable("measurements").Filter("device_id", deviceInstance.Id).OrderBy("-timestamp").Limit(100).Values(&meas, "timestamp", "temperature")
		if err != nil {
			beego.Error("error in getting Initial Measurements :", err)
		} else {
			beego.Debug("got Initial Measurements Number of results:", num)
			success = true
		}
	} else {
		beego.Error("GetMeasurementsByName - not find device")
	}
	return meas, success
}

func GetMeasurementsByDate(device string, startDate string, stopDate string) (meas []orm.Params, success bool) {
	success = false
	dateLayout := "2006-01-02"
	start, err1 := time.Parse(dateLayout, startDate)
	stop, err2 := time.Parse(dateLayout, stopDate)

	beego.Debug(startDate, start, stopDate, stop)
	if err1 == nil && err2 == nil {
		deviceInstance, res := GetDeviceByName(device)
		if res {
			var o orm.Ormer
			o = orm.NewOrm()
			num, err := o.QueryTable("measurements").Filter("device_id", deviceInstance.Id).Filter("timestamp__gte", start).Values(&meas, "timestamp", "temperature")
			if err != nil {
				beego.Error("error in getting Initial Measurements :", err)
			} else {
				beego.Debug("got Initial Measurements Number of results:", num)
				success = true
			}
		} else {
			beego.Error("GetMeasurementsByDate - did not find device")
		}
	}
	return meas, success
}

func GetLatestMeasurement(device string) (meas []models.Measurements, success bool) {
	success = false
	deviceInstance, res := GetDeviceByName(device)
	if res {
		var o orm.Ormer
		o = orm.NewOrm()
		err := o.QueryTable("measurements").Filter("device_id", deviceInstance.Id).OrderBy("-timestamp").One(&meas, "timestamp", "temperature")
		if err == nil {
			beego.Debug("Got latest measurement in GetLatestMeasurement num: ")
			success = true
		} else {
			beego.Error("Could NOT retrieve latest Measurement in GetLatestMeasurement eroor", err)
		}
	} else {
		beego.Error("Did NOT find device in GetLatestMeasurement")
	}
	return meas, success
}