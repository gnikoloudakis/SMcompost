package routers

import (
	"compost/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/measurements", &controllers.MeasurementsController{})
}
