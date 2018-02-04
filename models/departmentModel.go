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

type Department struct {
	Id         string   `db:"id, primarykey, autoincrement" json:"id"`
	Name       string   `db:"name" json:"name"`
	Updated_at int64    `db:"updated_at" json:"updated_at"`
	Created_at int64    `db:"created_at" json:"created_at"`
}

//DepartmentModel...
type DepartmentModel struct{}

var department_t = time.Now().Unix()

//Show all department...
func (m DepartmentModel) AllDepartment() (departments []Department, err error) {
	c := db.GetDB().C("department")
	err = c.Find(nil).All(&departments)
	return departments, err
}

//Find one department
func (m DepartmentModel) FindOneDepartment(departmentId string) (department Department, err error) {
	c := db.GetDB().C("department")
	err = c.Find(bson.M{"id": departmentId}).One(&department)
	return department, err
}

//insert one department
func (m DepartmentModel) InsertOneDepartment(form forms.InsertDepartmentForm) (department Department, err error) {
	var departmentid = uuid.Rand()
	c := db.GetDB().C("department")

	//var maps = make(map[string]interface{})
	//d, _ := json.Marshal(obj)
	//utils.JSONRaw().MarshalJSON()

	//插入之前判断数据是否已存在
	checkDepartment, err := c.Find(bson.M{"name": form.Name}).Count()
	if err != nil {
		return department, err
	}
	if checkDepartment > 0 {
		return department, errors.New("Department exists")
	}

	err = c.Insert(
		&Department{
			Id:      departmentid.Hex(),
			Name:    form.Name,
			Updated_at: department_t,
			Created_at: department_t})
	if err == nil {
		err = c.Find(bson.M{"name": form.Name}).One(&department)
		if err == nil {
			return department, nil
		}
	}

	return department, errors.New("Not registered")
}

//根据设备id更新
func (m DepartmentModel) UpsertDepartmentById(form forms.InsertDepartmentForm) (err error) {
	getDb := db.GetDB().C("department")
	//生成修改时间
	var department_t = time.Now().Unix()

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			"name":    form.Name,
			"updated_at": department_t}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert department fault")
	}
	return err
}

//DeleteOneAdmin
func (m DepartmentModel) DeleteOneDepartment(departmentId string) (err error) {
	var department Department
	err = db.GetDB().C("department").Find(bson.M{"id": departmentId}).One(&department)
	if err != nil {
		return errors.New("Could't find the department")
	}
	err = db.GetDB().C("department").Remove(bson.M{"id": departmentId})
	return err
}