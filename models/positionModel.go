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
)

type Position struct {
	Id         string   `db:"id, primarykey, autoincrement" json:"id"`
	Name       string   `db:"name" json:"name"`
	Updated_at int64    `db:"updated_at" json:"updated_at"`
	Created_at int64    `db:"created_at" json:"created_at"`
}

//PositionModel...
type PositionModel struct{}

var position_t = time.Now().Unix()

//Show all position...
func (m PositionModel) AllPosition() (positions []Position, err error) {
	c := db.GetDB().C("position")
	err = c.Find(nil).All(&positions)
	return positions, err
}

//Find one position
func (m PositionModel) FindOnePosition(positionId string) (position Position, err error) {
	c := db.GetDB().C("position")
	err = c.Find(bson.M{"id": positionId}).One(&position)
	return position, err
}

//insert one position
func (m PositionModel) InsertOnePosition(form forms.InsertPositionForm) (position Position, err error) {
	var positionid = uuid.Rand()
	c := db.GetDB().C("position")

	//var maps = make(map[string]interface{})
	//d, _ := json.Marshal(obj)
	//utils.JSONRaw().MarshalJSON()

	//插入之前判断数据是否已存在
	checkPosition, err := c.Find(bson.M{"name": form.Name}).Count()
	if err != nil {
		return position, err
	}
	if checkPosition > 0 {
		return position, errors.New("Position exists")
	}

	err = c.Insert(
		&Position{
			Id:      positionid.Hex(),
			Name:    form.Name,
			Updated_at: position_t,
			Created_at: position_t})
	if err == nil {
		err = c.Find(bson.M{"name": form.Name}).One(&position)
		if err == nil {
			return position, nil
		}
	}

	return position, errors.New("Not registered")
}

//根据设备id更新
func (m PositionModel) UpsertPositionById(form forms.InsertPositionForm) (err error) {
	getDb := db.GetDB().C("position")
	//生成修改时间
	var position_t = time.Now().Unix()

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			"name":    form.Name,
			"updated_at": position_t}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert position fault")
	}
	return err
}

//DeleteOneAdmin
func (m PositionModel) DeleteOnePosition(positionId string) (err error) {
	var position Position
	err = db.GetDB().C("position").Find(bson.M{"id": positionId}).One(&position)
	if err != nil {
		return errors.New("Could't find the position")
	}
	err = db.GetDB().C("position").Remove(bson.M{"id": positionId})
	return err
}