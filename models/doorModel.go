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
	"wy_ibms_demo/utils"
	"fmt"
	//"strconv"
	"strconv"
)

type Door struct {
	Id string `db:"id, primarykey, autoincrement" json:"id"`
	//门牌号
	Door_id string `db:"door_id" json:"door_id" binding:"required"`
	//门名称
	Name    string `db:"name" json:"name"`
	Type    string `db:"type" json:"type"`
	Ip      string `db:"ip" json:"ip"`
	Channel string `db:"channel" json:"channel"`
	//门的（开、关）状态
	Status string `db:"status" json:"status" binding:"required"`
	//门的在线、离线
	Online string `db:"online" json:"online" binding:"required"`
	//人脸识别图片（证件照）
	Image_url string `db:"image_url" json:"image_url" binding:"required"`

	//Location   utils.JSONRaw `db:"location" json:"location"`
	//Position    utils.JSONRaw  `db:"position" json:"position"`

	//门详细位置
	Location_name string `db:"location_name" json:"location_name" binding:"required"`
	Floor         string `db:"floor" json:"floor" binding:"required"`

	Door_position_x int `db:"door_position_x" json:"door_position_x" binding:"required"`
	Door_position_y int `db:"door_position_y" json:"door_position_y" binding:"required"`

	Updated_at int64 `db:"updated_at" json:"updated_at"`
	Created_at int64 `db:"created_at" json:"created_at"`
}

type DoorLog struct {
	Id string `db:"id" json:"id"`
	//门的唯一Id
	Door_id string `db:"door_id" json:"door_id" binding:"required"`
	Door_name string `db:"door_name" json:"door_name" binding:"required"`
	//人员的唯一Id
	Admin_id string `db:"admin_id" json:"admin_id" binding:"required"`
	Admin_name string `db:"admin_name" json:"admin_name" binding:"required"`
	//组织名字
	Organization_name string `db:"organization_name" json:"organization_name" binding:"required"`
	//开门方式
	Tcm_name string `db:"tcm_name" json:"tcm_name" binding:"required"`
	//进出门的时间
	Date string        `db:"date" json:"date" binding:"required"`
	Raw  utils.JSONRaw `db:"raw" json:"raw"`

	Updated_at int64 `db:"updated_at" json:"updated_at"`
	Created_at int64 `db:"created_at" json:"created_at"`
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

type DoorPageForm struct {
	Page     int   `db:"page" json:"page"`
	Pagesize int   `db:"pagesize" json:"pagesize"`
	Nums     int64 `db:"nums" json:"nums"`
}

//DoorModel...
type DoorModel struct{}

var door_t = time.Now().Unix()

//Show all door...
func (m DoorModel) AllDoor() (doors []Door, err error) {
	c := db.GetDB().C("door")
	err = c.Find(nil).All(&doors)
	return doors, err
}

//Find one door
func (m DoorModel) FindOneDoor(doorId string) (door Door, err error) {
	c := db.GetDB().C("door")
	err = c.Find(bson.M{"id": doorId}).One(&door)
	return door, err
}

//insert one door
func (m DoorModel) InsertOneDoor(form forms.InsertDoorForm) (door Door, err error) {
	var doorid = uuid.Rand()
	c := db.GetDB().C("door")

	//var maps = make(map[string]interface{})
	//d, _ := json.Marshal(obj)
	//utils.JSONRaw().MarshalJSON()

	//插入之前判断数据是否已存在
	checkDoor, err := c.Find(bson.M{"name": form.Name}).Count()
	if err != nil {
		return door, err
	}
	if checkDoor > 0 {
		return door, errors.New("Door exists")
	}

	err = c.Insert(
		&Door{
			Id:        doorid.Hex(),
			Door_id:   form.Door_id,
			Name:      form.Name,
			Type:      form.Type,
			Ip:        form.Ip,
			Channel:   form.Channel,
			Status:    form.Status,
			Online:    form.Online,
			Image_url: form.Image_url,
			//Location: form.Location,
			//Door_position: form.Door_position,

			Location_name: form.Location_name,
			Floor:         form.Floor,

			Door_position_x: form.Door_position_x,
			Door_position_y: form.Door_position_y,

			Updated_at: door_t,
			Created_at: door_t})
	if err == nil {
		err = c.Find(bson.M{"name": form.Name}).One(&door)
		if err == nil {
			return door, nil
		}
	}
	return door, errors.New("Not registered")
}

//根据设备id更新
func (m DoorModel) UpsertDoorById(form forms.InsertDoorForm) (err error) {
	getDb := db.GetDB().C("door")
	//生成修改时间
	var door_t = time.Now().Unix()

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			"name":      form.Name,
			"type":      form.Type,
			"ip":        form.Ip,
			"channel":   form.Channel,
			"status":    form.Status,
			"image_url": form.Image_url,
			//"location": form.Location,
			//"door_position": form.Door_position,

			"location_name":   form.Location_name,
			"floor":           form.Floor,
			"door_position_x": form.Door_position_x,
			"door_position_y": form.Door_position_y,

			"updated_at": door_t}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert door fault")
	}
	return err
}

//根据设备Id，修改开、关状态
func (m DoorModel) UpsertDoorStatusById(form forms.InsertDoorForm) (err error) {
	getDb := db.GetDB().C("door")

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			"status": form.Status,
		}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert door fault")
	}
	return err
}

//DeleteOneAdmin
func (m DoorModel) DeleteOneDoor(doorId string) (err error) {
	var door Door
	err = db.GetDB().C("door").Find(bson.M{"id": doorId}).One(&door)
	if err != nil {
		return errors.New("Could't find the door")
	}
	err = db.GetDB().C("door").Remove(bson.M{"id": doorId})
	return err
}

//insert one doorLog日志的插入不需要管，这个接口后期再维护
func (m DoorModel) InsertOneDoorLog(form forms.InsertDoorLog) (doorLog DoorLog, err error) {
	var doorid = uuid.Rand()
	c := db.GetDB().C("doorLog")

	//var maps = make(map[string]interface{})
	//d, _ := json.Marshal(obj)
	//utils.JSONRaw().MarshalJSON()

	//插入之前判断数据是否已存在
	checkDoor, err := c.Find(bson.M{"name": form.Id}).Count()
	if err != nil {
		return doorLog, err
	}
	if checkDoor > 0 {
		return doorLog, errors.New("DoorLog exists")
	}

	err = c.Insert(
		&DoorLog{
			Id:         doorid.Hex(),
			Door_id:    form.Door_id,
			Door_name:    form.Door_name,
			Admin_id:   form.Admin_id,
			Admin_name:   form.Admin_name,
			Organization_name:   form.Organization_name,
			Tcm_name:   form.Tcm_name,
			Date:       form.Date,
			Raw:	form.Raw,
			Updated_at: door_t,
			Created_at: door_t})
	if err == nil {
		err = c.Find(bson.M{"name": form.Id}).One(&doorLog)
		if err == nil {
			return doorLog, nil
		}
	}
	return doorLog, errors.New("Not registered")
}

//Show all doorLog...
func (m DoorModel) AllDoorLog() (doorLogs []DoorLog, err error) {
	c := db.GetDB().C("doorLog")
	err = c.Find(nil).All(&doorLogs)
	return doorLogs, err
}

//Find doorLog By 富士xx字段
func (m DoorModel) FindDoorLogByOrganizationNo(organizationNo string) (doorLogs []DoorLog, err error) {
	c := db.GetDB().C("doorLog")
	//err = c.Find(bson.M{}).Select(bson.M{"OrganizationNo":organizationNo}).One(&doorLogs)
	err = c.Find(bson.M{"raw.OrganizationNo": organizationNo}).One(&doorLogs)
	return doorLogs, err
}

//Find doorLog By 富士开门方式字段
func (m DoorModel) FindDoorLogByTcmId(tcmId string) (doorLogs []DoorLog, err error) {
	c := db.GetDB().C("doorLog")
	//err = c.Find(bson.M{}).Select(bson.M{"TcmId":tcmId}).One(&doorLogs)
	err = c.Find(bson.M{"raw.TcmId": tcmId}).One(&doorLogs)
	return doorLogs, err
}

//Find doorLog By 富士部门字段
func (m DoorModel) FindDoorLogByDevNo(devNo string) (doorLogs []DoorLog, err error) {
	c := db.GetDB().C("doorLog")
	//err = c.Find(bson.M{}).Select(bson.M{"DevNo":devNo}).One(&doorLogs)
	err = c.Find(bson.M{"raw.DevNo": devNo}).One(&doorLogs)

	return doorLogs, err
}

//Find doorLog By date-currentday
func (m DoorModel) FindDoorLogByDate(date string) (doorLogs []DoorLog, err error) {
	currentday:=utils.TimeToUnixInt64(date)
	nextday:=currentday-86400
	c := db.GetDB().C("doorLog")
	err = c.Find(bson.M{"created_at":bson.M{"$lte": currentday,"$gte": nextday}}).All(&doorLogs)
	return doorLogs, err
}

//Find one doorLog By doorId
func (m DoorModel) FindDoorLogByDoorId(doorId string) (doorLog DoorLog, err error) {
	c := db.GetDB().C("doorLog")
	err = c.Find(bson.M{"id": doorId}).One(&doorLog)
	return doorLog, err
}

//Find one doorLog By AdminId
func (m DoorModel) FindDoorLogByAdminId(adminId string) (doorLog DoorLog, err error) {
	c := db.GetDB().C("doorLog")
	err = c.Find(bson.M{"id": adminId}).One(&doorLog)
	return doorLog, err
}

//Find doorLog By condition
func (m DoorModel) FindDoorLogByCondition(form forms.FindDoorConditionForm) (doorLogs []DoorLog, paginatorMap map[string]int,err error,) {
	c := db.GetDB().C("doorLog")
	i:=bson.M{}
	if form.Door_id!="" {
		i["door_id"]=&form.Door_id
		//i["raw"]=bson.M{"OrganizationName": form.Department_name}
	}
	if form.Door_name!="" {
		i["door_name"]=form.Door_name
		//i["raw"]=bson.M{"DevName": form.DevName}
	}
	if form.Admin_id!="" {
		i["admin_id"]=form.Admin_id
		//i["raw"]=bson.M{"TcmName": form.TcmName}
	}
	if form.Admin_name!="" {
		i["admin_name"]=form.Admin_name
		//i["raw"]=bson.M{"StaffName": form.StaffName}
	}
	if form.Organization_name!="" {
		i["organization_name"]=form.Organization_name
	}
	if form.Tcm_name!="" {
		i["tcm_name"]=form.Tcm_name
	}
	if form.StartDate!=""&&form.EndDate!="" {
		////把string日期转化为int64
		//startDate,err:=strconv.ParseInt(form.StartDate,10,64)
		//endDate,err:=strconv.ParseInt(form.EndDate,10,64)
		//if err != nil {
		//	fmt.Println("日期格式转换有误！",startDate,endDate)
		//}
		fmt.Println("日期",form.StartDate,form.EndDate)
		i["date"]=&bson.M{"$gte": form.StartDate,"$lte": form.EndDate}
		//utils.JSONRaw()
		//i["raw"]=bson.M{"OpenDate": bson.M{"$lte": startDate,"$gte": endDate}}
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

	err = c.Find(i).Skip((page-1)*form.Pagesize).Limit(form.Pagesize).Sort("$natural").All(&doorLogs)
	//fmt.Println("条件查询数据doorLogs：",doorLogs)

	//怎么查询嵌套结构中的字段thinking
	//for k, v := range doorLogs {
	//	fmt.Println("key值",k)
	//	//var kkk bson.M
	//	ppp:=bson.M{}
	//	err:=bson.Unmarshal(v.Raw,ppp)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	//fmt.Println("err值",err)
	//	fmt.Println("value值",ppp)
	//}

	return doorLogs,paginatorMap, err
}

//查询日志总数
func (m DoorModel) LogCount() (sum int,err error){
	c := db.GetDB().C("doorLog")
	n,err := c.Find(nil).Count()
	return n,err
}

// Find one doorLog By date 分页
func (m DoorModel) FindDoorLogByPage(paginatorMap map[string]int,pagesize int) (doorLogs []DoorLog, err error) {
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
		fmt.Println("Capital of United States is not present")
	}

	c := db.GetDB().C("doorLog")
	err = c.Find(nil).Skip((page-1)*pagesize).Limit(pagesize).Sort("$natural").All(&doorLogs)
	//err = c.Find(nil).All(&doorLogs)

	return doorLogs, err
}

//DeleteOneDoorLog
func (m DoorModel) DeleteOneDoorLog(doorId string) (err error) {
	var door Door
	err = db.GetDB().C("door").Find(bson.M{"id": doorId}).One(&door)
	if err != nil {
		return errors.New("Could't find the door")
	}
	err = db.GetDB().C("door").Remove(bson.M{"id": doorId})
	return err
}
