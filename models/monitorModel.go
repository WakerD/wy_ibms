package models

import (
	"errors"
	//"fmt"
	//"github.com/gin-gonic/gin/json"
	"github.com/globalsign/mgo/bson"
	"github.com/snluu/uuid"
	//"golang.org/x/crypto/bcrypt"
	//"net"
	"time"
	"wy_ibms_demo/db"
	"wy_ibms_demo/forms"
	//"wy_ibms_demo/utils"
	"fmt"
	"strconv"
	"wy_ibms_demo/utils"
)

type Monitor struct {
	Id         string   `db:"id, primarykey, autoincrement" json:"id"`
	Name       string   `db:"name" json:"name"`
	Type       string   `db:"type" json:"type"`
	Ip         string   `db:"ip" json:"ip"`
	Channel    string   `db:"channel" json:"channel"`
	Status     string      `db:"status" json:"status"`
	Image_url string `db:"image_url" json:"image_url" binding:"required"`
	//Location   utils.JSONRaw `db:"location" json:"location"`
	//Position    utils.JSONRaw  `db:"position" json:"position"`

	Location_name string	`db:"location_name" json:"location_name" binding:"required"`
	Floor string	`db:"floor" json:"floor" binding:"required"`

	Monitor_position_x int	`db:"monitor_position_x" json:"monitor_position_x" binding:"required"`
	Monitor_position_y int	`db:"monitor_position_y" json:"monitor_position_y" binding:"required"`

	Updated_at int64    `db:"updated_at" json:"updated_at"`
	Created_at int64    `db:"created_at" json:"created_at"`
}

//type Location struct {
//	Location_name string `db:"location_name" json:"location_name"`
//	Floor         int    `db:"floor" json:"floor"`
//}
//
//type Position struct {
//	X int	`db:"x" json:"x"`
//	Y int	`db:"y" json:"y"`
//}

//MonitorModel...
type MonitorModel struct{}

var monitor_t = time.Now().Unix()

//Show all monitor...
func (m MonitorModel) AllMonitor() (monitors []Monitor, err error) {
	c := db.GetDB().C("monitor")
	err = c.Find(nil).All(&monitors)
	return monitors, err
}

//Find one monitor
func (m MonitorModel) FindOneMonitor(monitorId string) (monitor Monitor, err error) {
	c := db.GetDB().C("monitor")
	err = c.Find(bson.M{"id": monitorId}).One(&monitor)
	return monitor, err
}

//insert one monitor
func (m MonitorModel) InsertOneMonitor(form forms.InsertMonitorForm) (monitor Monitor, err error) {
	var monitorid = uuid.Rand()
	c := db.GetDB().C("monitor")

	//var maps = make(map[string]interface{})
	//d, _ := json.Marshal(obj)
	//utils.JSONRaw().MarshalJSON()

	//插入之前判断数据是否已存在
	checkMonitor, err := c.Find(bson.M{"name": form.Name}).Count()
	if err != nil {
		return monitor, err
	}
	if checkMonitor > 0 {
		return monitor, errors.New("Monitor exists")
	}

	err = c.Insert(
		&Monitor{
			Id:      monitorid.Hex(),
			Name:    form.Name,
			Type:    form.Type,
			Ip:      form.Ip,
			Channel: form.Channel,
			Status:  form.Status,
			Image_url: form.Image_url,
			//Location: form.Location,
			//Monitor_position: form.Monitor_position,

			Location_name: form.Location_name,
			Floor: form.Floor,

			Monitor_position_x: form.Monitor_position_x,
			Monitor_position_y: form.Monitor_position_y,

			Updated_at: monitor_t,
			Created_at: monitor_t})
	if err == nil {
		err = c.Find(bson.M{"name": form.Name}).One(&monitor)
		if err == nil {
			return monitor, nil
		}
	}
	return monitor, errors.New("Not registered")
}

//根据设备id更新
func (m MonitorModel) UpsertMonitorById(form forms.InsertMonitorForm) (err error) {
	getDb := db.GetDB().C("monitor")
	//生成修改时间
	var monitor_t = time.Now().Unix()

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			"name":    form.Name,
			"type":    form.Type,
			"ip":      form.Ip,
			"channel": form.Channel,
			"status":  form.Status,
			"image_url": form.Image_url,
			//"location": form.Location,
			//"monitor_position": form.Monitor_position,

			"location_name":	form.Location_name,
			"floor":	form.Floor,
			"monitor_position_x":	form.Monitor_position_x,
			"monitor_position_y":	form.Monitor_position_y,

			"updated_at": monitor_t}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert monitor fault")
	}
	return err
}

//DeleteOneAdmin
func (m MonitorModel) DeleteOneMonitor(monitorId string) (err error) {
	var monitor Monitor
	err = db.GetDB().C("monitor").Find(bson.M{"id": monitorId}).One(&monitor)
	if err != nil {
		return errors.New("Could't find the monitor")
	}
	err = db.GetDB().C("monitor").Remove(bson.M{"id": monitorId})
	return err
}

//查询日志总数
func (m MonitorModel) MonitorCount() (sum int,err error){
	c := db.GetDB().C("monitor")
	n,err := c.Find(nil).Count()
	return n,err
}

//Find doorLog By condition
func (m MonitorModel) FindMonitorByCondition(form forms.FindMonitorConditionForm) (monitors []Monitor,paginatorMap map[string]int, err error) {
	c := db.GetDB().C("monitor")
	i:=bson.M{}
	if form.Name!="" {
		i["name"]=form.Name
	}
	if form.Type!="" {
		i["type"]=form.Type
	}
	if form.Ip!="" {
		i["ip"]=form.Ip
	}
	if form.Status!="" {
		i["status"]=form.Status
	}
	if form.Floor!="" {
		i["floor"]=form.Floor
	}
	//求相应条件查询的总数sum
	sum,err:=c.Find(i).Count()
	if err != nil {
		fmt.Println("获取条件查询的总数sum有误",err)
	}
	sumstr:=strconv.Itoa(sum)
	sum64, err := strconv.ParseInt(sumstr, 10, 64)
	form.Nums=sum64
	//分页数据插入
	paginatorMap = utils.Paginator(form.Page,form.Pagesize,form.Nums)

	//pages, ok := paginatorMap["pages"]
	totalpages, ok := paginatorMap["totalpages"]
	firstpage, ok := paginatorMap["firstpage"]
	lastpage, ok := paginatorMap["lastpage"]
	page, ok := paginatorMap["currpage"]
	if ok{
		//fmt.Println("Capital of",pages,"is",paginatorMap["pages"])
		fmt.Println("Capital of",totalpages,"is",paginatorMap["totalpages"])
		fmt.Println("Capital of",firstpage,"is",paginatorMap["firstpage"])
		fmt.Println("Capital of",lastpage,"is",paginatorMap["lastpage"])
		fmt.Println("Capital of",page,"is",paginatorMap["currpage"])
	}else{
		fmt.Println("Capital of paginatorMap is not present")
	}
	fmt.Println(i)

	err = c.Find(i).Skip((page-1)*form.Pagesize).Limit(form.Pagesize).Sort("$natural").All(&monitors)

	return monitors,paginatorMap, err
}