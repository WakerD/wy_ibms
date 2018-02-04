package models

import (
	"errors"
	"time"
	"wy_ibms_demo/db"
	"wy_ibms_demo/forms"

	"github.com/globalsign/mgo/bson"
	"github.com/snluu/uuid"
	//"fmt"
)

//System
type System struct {
	Id            string `db:"id, primarykey, autoincrement" json:"id"`
	System_name       string `db:"system_name" json:"system_name"`
	Icon string `form:"icon" json:"icon" binding:"required"`
	Updated_at    int64  `db:"updated_at" json:"updated_at"`
	Created_at    int64  `db:"created_at" json:"created_at"`
}

//SystemModel ...
type SystemModel struct{}

//All ...
func (m SystemModel) AllSystem() (systems []System, err error) {
	c := db.GetDB().C("system")
	err = c.Find(nil).All(&systems)
	return systems, err
}

//查询权限系统总数
func (m SystemModel) SystemCount() (sum int,err error) {
	c := db.GetDB().C("system")
	n,err := c.Find(nil).Count()
	return n,err
}

//Find One System ...
func (m SystemModel) FindOneSystem(systemId string) (system System, err error) {
	err = db.GetDB().C("system").Find(bson.M{"id": systemId}).One(&system)
	return system, err
}

//insert one system
func (m SystemModel) InsertOneSystem(form forms.SystemForm) (system System, err error) {
	var system_id = uuid.Rand()
	var system_t = time.Now().Unix()
	c := db.GetDB().C("system")

	//var maps = make(map[string]interface{})
	//d, _ := json.Marshal(obj)
	//utils.JSONRaw().MarshalJSON()

	//插入之前判断数据是否已存在
	checkSystem, err := c.Find(bson.M{"system_name": form.System_name}).Count()
	if err != nil {
		return system, err
	}
	if checkSystem > 0 {
		return system, errors.New("System exists")
	}

	err = c.Insert(
		&System{
			Id:      system_id.Hex(),
			System_name:    form.System_name,
			//Authority:    form.Authority,
			Icon:    form.Icon,
			Updated_at: system_t,
			Created_at: system_t})
	if err == nil {
		err = c.Find(bson.M{"system_name": form.System_name}).One(&system)
		if err == nil {
			return system, nil
		}
	}

	return system, errors.New("Not insert system")
}

//根据SystemId更新
func (m SystemModel) UpsertSystemById(form forms.SystemForm) (err error) {
	getDb := db.GetDB().C("system")
	//生成修改时间
	var update_t = time.Now().Unix()

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			"system_name":  form.System_name,
			//"authority":    form.Authority,
			"icon":    form.Icon,
			"updated_at":   update_t}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert system fault")
	}
	return err
}

//DeleteOneSystem
func (m SystemModel) DeleteOneSystem(systemId string) (err error) {
	var system System
	err = db.GetDB().C("system").Find(bson.M{"id": systemId}).One(&system)
	if err != nil {
		return errors.New("Could't find the system")
	}
	err = db.GetDB().C("system").Remove(bson.M{"id": systemId})
	return err
}
