package main

import (
	"fmt"
	"compost/models"
	_ "compost/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/toolbox"
	"net/http"
	"math/rand"
	"encoding/json"
	"bytes"
	"strconv"
	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/compost?charset=utf8", 10, 10)
	orm.DefaultTimeLoc = time.Local
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

	var arduino1 models.Devices
	arduino1.Name = "Arduino#1"
	arduino1.IP = "192.168.1.10"

	var arduino2 models.Devices
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

	createTasks()
}

func createTasks() {

	addMeasurements := toolbox.NewTask("measurements", "0/5 * * * * *", func() error {//every 5 minutes
		type buffer struct {
			Device      string        `json:"Device"`
			Temperature float32       `json:"Temperature"`
		}
		id := rand.Intn(4)
		if id != 0 {
			meas := buffer{Device:"Arduino"+strconv.Itoa(id), Temperature:rand.Float32()}
			jsondata, err := json.Marshal(meas)
			resp, err := http.Post("http://localhost:8080/api/measurements/add", "application/json", bytes.NewBuffer(jsondata))
			beego.Debug("responce : ", resp, "err :", err)
		}
		return nil
	})

	toolbox.AddTask("measurements", addMeasurements)
	toolbox.StartTask()
	//defer toolbox.StopTask()

}

func main() {

	beego.Run()
}

