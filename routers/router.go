package routers

import (
	"github.com/astaxie/beego"
	"compost/controllers"
)

func init() {
	beego.Router("/dashboard", &controllers.DashboardController{})
	beego.Router("/api/devices/add", &controllers.AddDeviceController{})
	beego.Router("/api/devices/remove", &controllers.DeleteDeviceController{})//////// den exo grapsei ton controller akoma
	beego.Router("/api/measurements/add", &controllers.AddMeasurementsController{})
	beego.Router("/api/measurements/get/:device([a-zA-Z]+[1-9]+)", &controllers.GetMeasurements{}, "*:GetByName")
	beego.Router("/api/measurements/get/:device([a-zA-Z]+[1-9]+)/latest", &controllers.GetMeasurements{}, "*:GetLatestData")
	beego.Router("/api/measurements/get/:device([a-zA-Z]+[1-9]+)/:start([0-9]{2,4}-[0-9]{1,2}-[0-9]{1,2})/:stop([0-9]{2,4}-[0-9]{1,2}-[0-9]{1,2})", &controllers.GetMeasurements{}, "*:GetByDate")
	beego.Router("/charts", &controllers.ChartsController{})
}
