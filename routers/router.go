package routers

import (
	"github.com/astaxie/beego"
	"compost/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/devices/add", &controllers.AddDeviceController{})
	beego.Router("/api/devices/remove", &controllers.DeleteDeviceController{})//////// den exo grapsei ton controller akoma
	beego.Router("/api/measurements/add", &controllers.AddMeasurementsController{})
	beego.Router("/api/measurements/get/:device([a-zA-Z]+[1-9]+)", &controllers.GetMeasurements{})
	beego.Router("/charts", &controllers.ChartsController{})
}
