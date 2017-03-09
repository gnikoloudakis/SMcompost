package controllers

import (
	"github.com/astaxie/beego"
	"compost/modules"
)

type ChartsController struct {
	beego.Controller
}

//type ChartData struct {
//	Y  int             `json:"Y"`
//	X  time.Time       `json:"X"`
//	Id int             `json:"Id"`
//}
//type Arduino struct {
//	Id 	int			`orm:"auto"`
//	Name 	string			`orm:"unique;size(100)"`
//	IP 	string			`orm:"size(20)"`
//}


func (d *ChartsController) Get() {
	devices, _ := modules.GetDevices()
	////a :=fmt.Sprintf("{{'id':%d, 'X':%s, 'Y':%d}}", 2, time.Now().UTC(), rand.Intn(100))
	////p := fmt.Println
	////p("id = ", d.Ctx.Input.Param(":id"))///den exei doulepsei akom afto
	////p(a)
	//t := time.Now().UTC()
	////p("time", t)
	//dt := chartData{}
	//dt.X = t
	//dt.Y = rand.Intn(100)
	//dt.Id = rand.Intn(5)

	//d.Data["json"] = &dt
	//d.ServeJSON()
	d.Data["Devices"] = &devices
	d.TplName = "charts.html"

	//d.Ctx.Output.Body([]byte("OK!!"))
}



