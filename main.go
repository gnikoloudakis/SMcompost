package main

import (
	"fmt"
	"compost/models"
	_ "compost/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/compost?charset=utf8", 10, 10)
	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Database alias, Drop table and re-create, Print log
	err := orm.RunSyncdb("default", force, verbose)
	if err != nil {
		beego.Error(err)
	}

	o := orm.NewOrm()
	o.Using("compost")

	var arduino1 models.Arduino
	arduino1.Name = "Arduino#1"
	arduino1.IP = "192.168.1.10"

	var arduino2 models.Arduino
	arduino2.Name = "Arduino#2"
	arduino2.IP = "192.168.1.20"

	id1, err1 := o.Insert(&arduino1)
	if err1 == nil {
		fmt.Println("id of arduino is ", id1)
	} else {
		fmt.Println("name already exists")
	}
	id2, err2 := o.Insert(&arduino2)
	if err2 == nil {
		fmt.Println("id of arduino is ", id2)
	} else {
		fmt.Println("name already exists")
	}

}

func main() {
	beego.Run()
}

