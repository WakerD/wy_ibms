package test

import (
	"testing"
	"fmt"
	"strconv"
	"time"
	"gopkg.in/mgo.v2/bson"
	"wy_ibms_demo/models"
	"wy_ibms_demo/forms"
	"wy_ibms_demo/utils"
	"wy_ibms_demo/db"
	//"log"
	//"os"
	//"github.com/goburrow/modbus"
	//"math"
)

var meterId int
var slaveId string

func Test_string(t *testing.T)  {
	a,v,e:=1,2,3
	fmt.Println("电流",a,"电压",v,"电量",e)
}

func Test_timers(t *testing.T)  {
	ticker:=time.NewTicker(time.Millisecond * 500)
	go func() {
		for t:= range ticker.C{
			fmt.Println("Tick at",t)
		}
	}()
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func Test_channels(t *testing.T)  {
	timeChan:= time.NewTicker(time.Second).C
	tickChan:= time.NewTicker(time.Millisecond * 400).C
	doneChan:= make(chan bool)
	go func() {
		time.Sleep(time.Second*2)
		doneChan<-true
	}()
	for {
		select {
		case <-timeChan:
			fmt.Println("Timer expired")
		case <-tickChan:
			fmt.Println("Ticker ticked")
		case <-doneChan:
			fmt.Println("Done")
			return
		}
	}
}

func Test_channels1(t *testing.T)  {
	for{
		timer:=time.NewTicker(2*time.Second)
		<-timer.C
		fmt.Println("看看时间",time.Now().Format("2006-01-02 15:04:05"))
	}
}

func Test_mongodate(t *testing.T)  {
	var models = new(models.UserModel)
	var VcodeForm forms.VcodeForm
	VcodeForm.Phonenumber="12343214321"
	VcodeForm.Vcode=9876
	models.InsertVcode(VcodeForm)
	//now:=time.Now()
	//fmt.Println("当前时间：",now)
}

func Test_timeUtcToCst(t *testing.T)  {
	//tt:=time.Time.Add(time.Hour * time.Duration(8))
	//const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	//t1, _ := time.Parse(longForm, "Jun 21, 2017 at 0:00am (PST)")
	//fmt.Println("",t1)
	//
	//const shortForm = "2006-Jan-02"
	//t1, _ = time.Parse(shortForm, "2017-Jun-21")
	//fmt.Println(t1)
	//
	//t1, _ = time.Parse("01/02/2006", "06/21/2017")
	//fmt.Println(t1)
	//fmt.Println(t1.Unix())

	timeLayout:="2006-01-02 15:04:05 PM"
	i1, err := strconv.ParseInt("1498003200", 10, 64)
	i2, err := strconv.ParseInt("1516094148", 10, 64)
	//1516093528,1516122580
	//29052
	if err != nil {
		panic(err)
	}
	tm1 := time.Unix(i1, 0)
	tm2 := time.Unix(i2, 0).Format(timeLayout)
	fmt.Println(tm1)
	fmt.Println(tm2)

	//var timestamp int64 = 1498003200
	//tm3 := time.Unix(timestamp, 0)
	//fmt.Println(tm3.Format("2006-01-02 03:04:05 PM"))
	//fmt.Println(tm3.Format("02/01/2006 15:04:05 PM"))
}

func Test_NewTime(t *testing.T)  {

	v:=utils.UnixInt64ToTime(1516947567)
	fmt.Println("日期：",v)
}

//日期格式转化为int64时间戳
func Test_TimeToUnixInt64(t *testing.T) {
	var monitor_t = time.Now().Unix()
	//fmt.Println(monitor_t)
	fmt.Println(utils.UnixInt64ToTime(monitor_t))

	//tm2, _ := time.Parse("2006-01-02 15:04:05", utils.UnixInt64ToTime(monitor_t))
	//fmt.Println(tm2.Unix()-28800)
	//fmt.Println("当天时间：",monitor_t)
}

//测试一段时间查询
func Test_FindDate(t *testing.T)  {
	var electricityMeterModel = new(models.ElectricityMeterModel)
	startstamp := time.Now().Unix()
	endstamp := startstamp + 86400
	//endstamp := startstamp + 45
	electricityLogs, err := electricityMeterModel.FindMeterLogByTypeAndDate(strconv.Itoa(meterId),"every_minute", startstamp, endstamp)
	if err!=nil {
		fmt.Println("FindMeterLogByTypeAndDate有误",electricityLogs)
		return
	}
	//求一个电表每天的电表参数均值
	a_sum, v_sum, e_sum, n := 0.00, 0.00, 0.00, 0.00
	//var SlaveId string
	for _, v := range electricityLogs {
		//insertElectricMeterLogForm.Created_at=utils.UnixInt64ToTime(v.Created_at)
		a_sum = a_sum + v.Current
		v_sum = v_sum + v.Voltage
		e_sum = e_sum + v.TolalActiveEnergy
		n++
	}
	fmt.Println("设备号",meterId)
	fmt.Println("电流",a_sum / n)
	fmt.Println("电压",v_sum / n)
	fmt.Println("总电量",e_sum / n)
}

type ElectricityLog struct {
	Id string	`db:"id" json:"id"`
	SlaveId string	`db:"slaveid" json:"slaveid"`
	Type string		`db:"type" json:"type"`
	Current float64 `db:"current" json:"current"`
	Voltage float64 `db:"voltage" json:"voltage"`
	TolalActiveEnergy float64	`db:"tolalactiveenergy" json:"tolalactiveenergy"`
	//
	Updated_at int64	`db:"updated_at" json:"updated_at"`
	Created_at int64	`db:"created_at" json:"created_at"`
}

func Test_mgo(t *testing.T)  {
	var electricityLogs []ElectricityLog
	startstamp := time.Now().Unix()
	endstamp := startstamp + 86400
	c:=db.GetDB().C("electricityLogMin")
	c.Find(&bson.M{
		"slaveid":meterId,
		"type":"every_minute",
		"created_at":bson.M{
			"$gt": startstamp,
			"$lt": endstamp}}).All(&electricityLogs)
	fmt.Println("查询出的日志",electricityLogs)
}

func Test_nullstr(t *testing.T)  {
	var str string
	if str==""{
		fmt.Println("string默认值",str)
	}else {
		fmt.Println("----")
	}

	var intt int
	if intt==0{
		fmt.Println("intt默认值",intt)
	}else {
		fmt.Println("----")
	}
}

func Test_Map(t *testing.T)  {
	//// 直接创建
	//m2 := make(map[string]interface{})
	//// 然后赋值
	//m2["a"] = "aa"
	//m2["b"] = "bb"
	//// ==========================================
	//// 查找键值是否存在
	//if v, ok := m2["a"]; ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("Key Not Found")
	//}
	//// 遍历map
	//for k, v := range m2 {
	//	fmt.Println(k, v)
	//}

	//i:=bson.M{}
	//if "1"!="" {
	//	i["OrganizationName"]="1"
	//}
	//if "2"!="" {
	//	i["DevName"]="2"
	//}
	//if ""!="" {
	//	i["TcmName"]="3"
	//}
	//if "4"!="" {
	//	i["StaffName"]="4"
	//}
	//if "5"!="" {
	//	i["OpenDate"]="5"
	//}
	//if ""!="" {
	//	i["DoorId"]="6"
	//}
	//
	//for k, v := range i {
	//	fmt.Println(k, v)
	//}

	selector := bson.M{"id": 1,"ic":2,"ie":3}
	for k, v := range selector {
		fmt.Println(k, v)
	}
}

type TestStruct struct {
	Name string
	ID   int32
}

func Test_Bson(t *testing.T)  {
	fmt.Println("start")
	data, err := bson.Marshal(&TestStruct{Name: "Bob"})
	if err != nil {
		panic(err)
	}
	fmt.Println("%q", data)

	value := TestStruct{}
	err2 := bson.Unmarshal(data, &value)
	if err2 != nil {
		panic(err)
	}
	fmt.Println("value:", value)

	mmap := bson.M{}
	err3 := bson.Unmarshal(data, mmap)
	if err3 != nil {
		panic(err)
	}
	fmt.Println("mmap:", mmap)
}

//测试||或，&&与
func Test_YuAndHuo(t *testing.T)  {

	a_avg:=123.0
	v_avg:="NaNo"
	if a_avg==123 || v_avg=="NaN" {
		fmt.Println("electricMeter插入的数值有NAN值,不存入mongo", a_avg, v_avg)
	}
}