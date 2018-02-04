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
)

type Floor struct {
	Id         string   `db:"id, primarykey, autoincrement" json:"id"`
	Name       string   `db:"name" json:"name"`
	Type       string   `db:"type" json:"type"`
	Floor         int   `db:"floor" json:"floor"`
	Image_url    string   `db:"image_url" json:"image_url"`
	Grid     utils.JSONRaw      `db:"grid" json:"grid"`

	//Floor_grid_x int `db:"floor_grid_x" json:"floor_grid_x"`
	//Floor_grid_y int `db:"Floor_grid_y" json:"Floor_grid_y"`

	Updated_at int64    `db:"updated_at" json:"updated_at"`
	Created_at int64    `db:"created_at" json:"created_at"`
}

type Grid struct{
	X int	`db:"x" json:"x"`
	Y int	`db:"y" json:"y"`
}

//FloorModel...
type FloorModel struct{}

var floor_t = time.Now().Unix()

//Show all floor...
func (m FloorModel) AllFloor() (floors []Floor, err error) {
	c := db.GetDB().C("floor")
	err = c.Find(nil).All(&floors)
	return floors, err
}

//Find one floor
func (m FloorModel) FindOneFloor(floorId string) (floor Floor, err error) {
	c := db.GetDB().C("floor")
	err = c.Find(bson.M{"id": floorId}).One(&floor)
	return floor, err
}

//insert one floor
func (m FloorModel) InsertOneFloor(form forms.InsertFloorForm) (floor Floor, err error) {
	var floorid = uuid.Rand()
	c := db.GetDB().C("floor")

	//var maps = make(map[string]interface{})
	//d, _ := json.Marshal(obj)
	//utils.JSONRaw().MarshalJSON()

	//插入之前判断数据是否已存在
	checkFloor, err := c.Find(bson.M{"name": form.Name}).Count()
	if err != nil {
		return floor, err
	}
	if checkFloor > 0 {
		return floor, errors.New("Floor exists")
	}

	err = c.Insert(
		&Floor{
			Id:      floorid.Hex(),
			Name:    form.Name,
			Type:    form.Type,
			Floor:      form.Floor,
			Image_url: form.Image_url,
			Grid: form.Grid,

			//Floor_grid_x:	form.Floor_grid_x,
			//Floor_grid_y:	form.Floor_grid_y,

			Updated_at: floor_t,
			Created_at: floor_t})
	if err == nil {
		err = c.Find(bson.M{"name": form.Name}).One(&floor)
		if err == nil {
			return floor, nil
		}
	}

	return floor, errors.New("Not registered")
}

//根据设备id更新
func (m FloorModel) UpsertFloorById(form forms.InsertFloorForm) (err error) {
	getDb := db.GetDB().C("floor")
	//生成修改时间
	var floor_t = time.Now().Unix()

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			"name":    form.Name,
			"type":    form.Type,
			"floor":      form.Floor,
			"image_url": form.Image_url,
			"grid": form.Grid,

			//"Floor_grid_x":	form.Floor_grid_x,
			//"Floor_grid_y":	form.Floor_grid_y,

			"updated_at": floor_t}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert floor fault")
	}
	return err
}

//DeleteOneAdmin
func (m FloorModel) DeleteOneFloor(floorId string) (err error) {
	var floor Floor
	err = db.GetDB().C("floor").Find(bson.M{"id": floorId}).One(&floor)
	if err != nil {
		return errors.New("Could't find the floor")
	}
	err = db.GetDB().C("floor").Remove(bson.M{"id": floorId})
	return err
}