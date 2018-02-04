package models

import (
	"errors"
	"time"
	"wy_ibms_demo/db"
	"wy_ibms_demo/forms"

	"github.com/globalsign/mgo/bson"
	"github.com/snluu/uuid"
	"wy_ibms_demo/utils"
	"fmt"
)

//role
type Role struct {
	Id            string `db:"id, primarykey, autoincrement" json:"id"`
	Role_name       string `db:"role_name" json:"role_name"`
	//Authority      string `db:"authority" json:"authority"`
	Auth 	map[string]interface{} `db:"auth" json:"auth"`
	Updated_at    int64  `db:"updated_at" json:"updated_at"`
	Created_at    int64  `db:"created_at" json:"created_at"`
}

//RoleModel ...
type RoleModel struct{}

//All ...
func (m RoleModel) AllRole() (roles []Role, err error) {
	c := db.GetDB().C("role")
	err = c.Find(nil).All(&roles)
	return roles, err
}

//查询权限系统总数
func (m RoleModel) RoleCount() (sum int,err error) {
	c := db.GetDB().C("role")
	n,err := c.Find(nil).Count()
	return n,err
}

//FindOneRole ...
func (m RoleModel) FindOneRole(roleId string) (role Role, err error) {
	err = db.GetDB().C("role").Find(bson.M{"id": roleId}).One(&role)
	return role, err
}

//insert one role
func (m RoleModel) InsertOneRole(form forms.RoleForm) (role Role, err error) {
	var role_id = uuid.Rand()
	var role_t = time.Now().Unix()
	c := db.GetDB().C("role")

	//var maps = make(map[string]interface{})
	//d, _ := json.Marshal(obj)
	//utils.JSONRaw().MarshalJSON()

	//插入之前判断数据是否已存在
	checkRole, err := c.Find(bson.M{"role_name": form.Role_name}).Count()
	if err != nil {
		return role, err
	}
	if checkRole > 0 {
		return role, errors.New("Role exists")
	}
	//auth转化为map对象
	//auth,err:=utils.Obj2map(form.Auth)

	//i:=bson.M{}
	//for k,v:=range form.Auth{
	//	i[k]=v
	//}
	
	err = c.Insert(
		&Role{
			Id:      role_id.Hex(),
			Role_name:    form.Role_name,
			//Authority:    form.Authority,
			Auth:    form.Auth,
			Updated_at: role_t,
			Created_at: role_t})
	if err == nil {
		err = c.Find(bson.M{"role_name": form.Role_name}).One(&role)
		if err == nil {
			fmt.Println(role)
			return role, nil
		}
	}

	return role, errors.New("Not insert role")
}

//根据RoleId更新
func (m RoleModel) UpsertRoleById(form forms.RoleForm) (err error) {
	getDb := db.GetDB().C("role")
	//生成修改时间
	var update_t = time.Now().Unix()

	//auth转化为map对象
	auth,err:=utils.Obj2map(form.Auth)
	i:=bson.M{}
	for k,v:=range auth{
		i[k]=v
	}

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			"role_name":  form.Role_name,
			//"authority":    form.Authority,
			"auth":    i,
			"updated_at":   update_t}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert role fault")
	}
	return err
}

//DeleteOneRole
func (m RoleModel) DeleteOneRole(roleId string) (err error) {
	var role Role
	err = db.GetDB().C("role").Find(bson.M{"id": roleId}).One(&role)
	if err != nil {
		return errors.New("Could't find the role")
	}
	err = db.GetDB().C("role").Remove(bson.M{"id": roleId})
	return err
}
