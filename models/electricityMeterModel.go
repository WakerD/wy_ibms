package models

import (
	"time"
	"wy_ibms_demo/db"
	"wy_ibms_demo/utils"
	// "gopkg.in/mgo.v2/bson"
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/snluu/uuid"
	"wy_ibms_demo/forms"
	//"strconv"
	"fmt"
)

type ElectricityMeter struct {
	Id         string        `db:"id, primarykey, autoincrement" json:"id"`
	Name       string        `db:"name" json:"name"`
	SubordinateId    string        `db:"subordinateid" json:"subordinateid"`
	Type       string        `db:"type" json:"type"`
	Ip         string        `db:"ip" json:"ip"`
	Channel    string        `db:"channel" json:"channel"`
	Status     string        `db:"status" json:"status"`
	Image_url  string        `db:"image_url" json:"image_url"`
	Location   utils.JSONRaw `db:"location" json:"location"`
	Position   utils.JSONRaw `db:"position" json:"position"`
	Updated_at int64         `db:"updated_at" json:"updated_at"`
	Created_at int64         `db:"created_at" json:"created_at"`
}

type ElectricityLog struct {
	Id                string  `db:"id" json:"id"`
	SubordinateId           string  `db:"subordinateid" json:"subordinateid"`
	Type              string  `db:"type" json:"type"`
	Current           float64 `db:"current" json:"current"`
	Voltage           float64 `db:"voltage" json:"voltage"`
	TolalActiveEnergy float64 `db:"tolalactiveenergy" json:"tolalactiveenergy"`

	Updated_at int64 `db:"updated_at" json:"updated_at"`
	Created_at int64 `db:"created_at" json:"created_at"`
}

//ElectricityMeterModel...
type ElectricityMeterModel struct{}

//Show all electricity...
func (m ElectricityMeterModel) AllElectricity() (electricityMeters []ElectricityMeter, err error) {
	c := db.GetDB().C("electricity")
	c.Find(nil).All(&electricityMeters)
	return electricityMeters, err
}

func (m ElectricityMeterModel) AllMeterLog() (meterlog []ElectricityLog, err error) {
	c := db.GetDB().C("electricityLogMin")
	c.Find(nil).All(&meterlog)
	return meterlog, err
}

//Find one monitor
func (m ElectricityMeterModel) FindOneMeter(electricityMeterId string) (electricityMeter ElectricityMeter, err error) {
	c := db.GetDB().C("electricity")
	c.Find(bson.M{"id": electricityMeterId}).One(&electricityMeter)
	return electricityMeter, err
}

func (m ElectricityMeterModel) InsertOneMeter(form forms.InsertElectricMeterForm) (electricityMeter ElectricityMeter, err error) {
	var meterid = uuid.Rand()
	var electricity_t = time.Now().Unix()

	c := db.GetDB().C("electricity")
	//插入之前判断数据是否已存在
	checkMeter, err := c.Find(bson.M{"name": form.Name}).Count()
	if err != nil {
		return electricityMeter, err
	}
	if checkMeter > 0 {
		return electricityMeter, errors.New("ElectricityMeter exists")
	}

	err = c.Insert(
		&ElectricityMeter{
			Id:         meterid.Hex(),
			Name:       form.Name,
			Type:       form.Type,
			Ip:         form.Ip,
			Channel:    form.Channel,
			Status:     form.Status,
			Image_url:  form.Image_url,
			Location:   form.Location,
			Position:   form.Position,
			Updated_at: electricity_t,
			Created_at: electricity_t})
	if err == nil {
		err = c.Find(bson.M{"name": form.Name}).One(&electricityMeter)
		if err == nil {
			return electricityMeter, nil
		}
	}
	return electricityMeter, errors.New("Not registered")
}

//按minute插入电表日志记录
func (m ElectricityMeterModel) InsertOneMeterLogMin(form forms.InsertElectricMeterLogForm) (electricityLog ElectricityLog, err error) {
	var meterid = uuid.Rand()
	var electricity_t = time.Now().Unix()

	form.Id = meterid.Hex()
	form.Type = "every_minute"

	c := db.GetDB().C("electricityLogMin")
	//插入之前判断数据是否已存在
	checkMeterLog, err := c.Find(bson.M{"id": form.Id}).Count()
	if err != nil {
		return electricityLog, err
	}
	if checkMeterLog > 0 {
		return electricityLog, errors.New("ElectricityLog exists")
	}

	err = c.Insert(
		&ElectricityLog{
			Id:                form.Id,
			SubordinateId:           form.SubordinateId,
			Type:              form.Type,
			Current:           form.Current,
			Voltage:           form.Voltage,
			TolalActiveEnergy: form.TolalActiveEnergy,
			Updated_at:        electricity_t,
			Created_at:        electricity_t})
	if err == nil {
		err = c.Find(bson.M{"id": form.Id}).One(&electricityLog)
		if err == nil {
			return electricityLog, nil
		}
	}
	return electricityLog, errors.New("Not insert per minute")
}

//按Hour插入电表日志记录
func (m ElectricityMeterModel) InsertOneMeterLogHour(form forms.InsertElectricMeterLogForm) (electricityLog ElectricityLog, err error) {
	var meterid = uuid.Rand()
	var electricity_t = time.Now().Unix()

	form.Id = meterid.Hex()
	form.Type = "every_hour"

	c := db.GetDB().C("electricityLogHour")
	//插入之前判断数据是否已存在
	checkMeterLog, err := c.Find(bson.M{"id": form.Id}).Count()
	if err != nil {
		return electricityLog, err
	}
	if checkMeterLog > 0 {
		return electricityLog, errors.New("ElectricityLog exists")
	}

	err = c.Insert(
		&ElectricityLog{
			Id:                form.Id,
			SubordinateId:           form.SubordinateId,
			Type:              form.Type,
			Current:           form.Current,
			Voltage:           form.Voltage,
			TolalActiveEnergy: form.TolalActiveEnergy,
			Updated_at:        electricity_t,
			Created_at:        electricity_t})
	if err == nil {
		err = c.Find(bson.M{"id": form.Id}).One(&electricityLog)
		if err == nil {
			return electricityLog, nil
		}
	}
	return electricityLog, errors.New("Not insert per hour")
}

//按day插入电表日志记录
func (m ElectricityMeterModel) InsertOneMeterLogDay(form forms.InsertElectricMeterLogForm) (electricityLog ElectricityLog, err error) {
	var meterid = uuid.Rand()
	var electricity_t = time.Now().Unix()

	form.Id = meterid.Hex()
	form.Type = "every_day"

	c := db.GetDB().C("electricityLogDay")
	//插入之前判断数据是否已存在
	checkMeterLog, err := c.Find(bson.M{"id": form.Id}).Count()
	if err != nil {
		return electricityLog, err
	}
	if checkMeterLog > 0 {
		return electricityLog, errors.New("ElectricityLog exists")
	}

	err = c.Insert(
		&ElectricityLog{
			Id:                form.Id,
			SubordinateId:           form.SubordinateId,
			Type:              form.Type,
			Current:           form.Current,
			Voltage:           form.Voltage,
			TolalActiveEnergy: form.TolalActiveEnergy,
			Updated_at:        electricity_t,
			Created_at:        electricity_t})
	if err == nil {
		err = c.Find(bson.M{"id": form.Id}).One(&electricityLog)
		if err == nil {
			return electricityLog, nil
		}
	}
	return electricityLog, errors.New("Not insert per day")
}

//按month插入电表日志记录
func (m ElectricityMeterModel) InsertOneMeterLogMonth(form forms.InsertElectricMeterLogForm) (electricityLog ElectricityLog, err error) {
	var meterid = uuid.Rand()
	var electricity_t = time.Now().Unix()

	form.Id = meterid.Hex()
	form.Type = "every_month"

	c := db.GetDB().C("electricityLogMth")
	//插入之前判断数据是否已存在
	checkMeterLog, err := c.Find(bson.M{"id": form.Id}).Count()
	if err != nil {
		return electricityLog, err
	}
	if checkMeterLog > 0 {
		return electricityLog, errors.New("ElectricityLog exists")
	}

	err = c.Insert(
		&ElectricityLog{
			Id:                form.Id,
			SubordinateId:           form.SubordinateId,
			Type:              form.Type,
			Current:           form.Current,
			Voltage:           form.Voltage,
			TolalActiveEnergy: form.TolalActiveEnergy,
			Updated_at:        electricity_t,
			Created_at:        electricity_t})
	if err == nil {
		err = c.Find(bson.M{"id": form.Id}).One(&electricityLog)
		if err == nil {
			return electricityLog, nil
		}
	}
	return electricityLog, errors.New("Not insert per month")
}

//根据日志类型查询一段时间
func (m ElectricityMeterModel) FindMeterLogByTypeAndDate(meterId string, logtype string, starttime int64, endtime int64) (electricityLogs []ElectricityLog, err error) {
	if logtype == "every_minute" {
		fmt.Println("starttime %d - endtime %d", starttime, endtime)
		c := db.GetDB().C("electricityLogMin")
		c.Find(&bson.M{
			"subordinateid": meterId,
			"type":    logtype,
			"created_at": bson.M{
				"$lte": endtime,
				"$gte": starttime}}).All(&electricityLogs)
		return electricityLogs, err
	} else if logtype == "every_hour" {
		c := db.GetDB().C("electricityLogHour")
		c.Find(&bson.M{
			"subordinateid": meterId,
			"type":    logtype,
			"created_at": bson.M{
				"$gte": starttime,
				"$lte": endtime}}).All(&electricityLogs)
		return electricityLogs, err
	} else if logtype == "every_day" {
		c := db.GetDB().C("electricityLogDay")
		c.Find(&bson.M{
			"subordinateid": meterId,
			"type":    logtype,
			"created_at": bson.M{
				"$gte": starttime,
				"$lte": endtime}}).All(&electricityLogs)
		return electricityLogs, err
	} else if logtype == "every_month" {
		c := db.GetDB().C("electricityLogMth")
		c.Find(&bson.M{
			"subordinateid": meterId,
			"type":    logtype,
			"created_at": bson.M{
				"$gte": starttime,
				"$lte": endtime}}).All(&electricityLogs)
		return electricityLogs, err
	}

	return electricityLogs,errors.New("查询一段时间电表日志出错")
}

// //根据日志类型查询一段时间
// func (m ElectricityMeterModel) FindMeterLogByTypeAndDate(meterId string, logtype string, starttime int64, endtime int64) (electricityLogs []ElectricityLog, err error) {
// 	if logtype == "every_minute" {
// 		fmt.Println("starttime %d - endtime %d", starttime, endtime)
// 		c := db.GetDB().C("electricityLogMin")
// 		c.Find(&bson.M{
// 			"subordinateid": meterId,
// 			"type":    logtype,
// 			"created_at": bson.M{
// 				"$lte": endtime,
// 				"$gte": starttime}}).All(&electricityLogs)
// 		return electricityLogs, err
// 	} else if logtype == "every_day" {
// 		c := db.GetDB().C("electricityLogDay")
// 		c.Find(&bson.M{
// 			"subordinateid": meterId,
// 			"type":    logtype,
// 			"created_at": bson.M{
// 				"$gte": starttime,
// 				"$lte": endtime}}).All(&electricityLogs)
// 		return electricityLogs, err
// 	} else if logtype == "every_month" {
// 		c := db.GetDB().C("electricityLogMth")
// 		c.Find(&bson.M{"subordinateid": meterId, "type": logtype, "created_at": bson.M{"$gte": starttime, "$lte": endtime}}).All(&electricityLogs)
// 		return electricityLogs, err
// 	} else {
// 		fmt.Println("查询一段时间电表日志出错")
// 		return
// 	}
// }

//根据设备id更新
func (m ElectricityMeterModel) UpsertMeterById(form forms.InsertElectricMeterForm) (err error) {
	var electricity_t = time.Now().Unix()
	c := db.GetDB().C("electricity")

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			"name":       form.Name,
			"type":       form.Type,
			"ip":         form.Ip,
			"channel":    form.Channel,
			"status":     form.Status,
			"image_url":  form.Image_url,
			"location":   form.Location,
			"position":   form.Position,
			"updated_at": electricity_t}}
	_, err = c.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert electricity fault")
	}
	return err
}

func (m ElectricityMeterModel) DeleteOneMeter(electricityMeterId string) (err error) {
	var electricityMeter ElectricityMeter
	err = db.GetDB().C("electricity").Find(bson.M{"id": electricityMeterId}).One(&electricityMeter)
	if err != nil {
		return errors.New("Could't find the electricityMeter")
	}
	err = db.GetDB().C("electricity").Remove(bson.M{"id": electricityMeterId})
	return err
}
