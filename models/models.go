package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Devices struct{
	Id 	int			`orm:"auto"`
	Name 	string			`orm:"unique;size(100)"`
	IP 	string			`orm:"unique;size(20)"`
	//Measurements	*Measurements	`orm:"rel(fk)"`
}

type Measurements struct{
	Id 		int		`orm:"auto"`
	Timestamp	time.Time 	`orm:"auto_now_add;type(datetime)"`
	Temperature	float32		`orm:"size(10)"`
	Device 	*Devices 		`orm:"rel(fk);on_delete(cascade)"`
}

func init(){
	orm.RegisterModel(new(Devices), new(Measurements))
}
