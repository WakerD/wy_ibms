package models

import (
	"errors"
	// "fmt"
	"time"

	"wy_ibms_demo/db"
	"wy_ibms_demo/forms"

	"github.com/globalsign/mgo/bson"
	"github.com/snluu/uuid"
	"golang.org/x/crypto/bcrypt"
)

//admin
type Admin struct {
	Id            string `db:"id, primarykey, autoincrement" json:"id"`
	Account       string `db:"account" json:"account"`
	Password      string `db:"password" json:"password"`
	Phone_number  string `db:"phone_number" json:"phone_number"`
	Email         string `db:"email" json:"email"`
	Username      string `db:"username" json:"username"`
	Address       string `db:"address" json:"address"`
	Gender        int    `db:"gender" json:"gender"`
	Identity_no   string `db:"identity_no" json:"identity_no"`
	Identity_type string `db:"identity_type" json:"identity_type"`
	Department    string `db:"department" json:"department"`
	Position      string `db:"position" json:"position"`
	Updated_at    int64  `db:"updated_at" json:"updated_at"`
	Created_at    int64  `db:"created_at" json:"created_at"`
}

//AdminModel ...
type AdminModel struct{}

var admin_t = time.Now().Unix()

//All ...
func (m AdminModel) AllAdmin() (admins []Admin, err error) {
	c := db.GetDB().C("admin")
	err = c.Find(nil).All(&admins)
	return admins, err
}

//AdminSignin ...
func (m AdminModel) AdminSignin(form forms.AdminSigninForm) (admin Admin, err error) {

	err = db.GetDB().C("admin").Find(bson.M{"account": form.Account}).One(&admin)

	if err != nil {
		return admin, err
	}

	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(admin.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return admin, errors.New("Invalid password")
	}
	return admin, nil
}

//AdminSignup ...
func (m AdminModel) AdminSignup(form forms.AdminSignupForm) (admin Admin, err error) {
	var (
		adminid = uuid.Rand()
		// data = time.Unix(1514291163, 0).Format("2006-01-02 13:04:05")
	)
	getDb := db.GetDB().C("admin")

	checkAdmin, err := getDb.Find(bson.M{"account": form.Account}).Count()
	if err != nil {
		return admin, err
	}

	if checkAdmin > 0 {
		return admin, errors.New("Admin exists")
	}
	//用户密码加密
	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	//用户id用uuid
	err = getDb.Insert(
		&Admin{
			Id:            adminid.Hex(),
			Account:       form.Account,
			Password:      string(hashedPassword),
			Phone_number:  form.Phone_number,
			Email:         form.Email,
			Username:      form.Username,
			Address:       form.Address,
			Gender:        form.Gender,
			Identity_no:   form.Identity_no,
			Identity_type: form.Identity_type,
			Department:    form.Department,
			Position:      form.Position,
			Updated_at:    admin_t,
			Created_at:    admin_t})

	if err == nil {
		err = getDb.Find(bson.M{"account": form.Account}).One(&admin)
		if err == nil {
			return admin, nil
		}
	}

	return admin, errors.New("Not registered")
}

//FindOneAdmin ...
func (m AdminModel) FindOneAdmin(adminId string) (admin Admin, err error) {
	err = db.GetDB().C("admin").Find(bson.M{"id": adminId}).One(&admin)
	return admin, err
}

//根据用户id更新
func (m AdminModel) UpsertAdminById(form forms.AdminSignupForm) (err error) {
	getDb := db.GetDB().C("admin")
	//用户密码加密
	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	//生成修改时间
	var update_t = time.Now().Unix()

	selector := bson.M{"id": form.Id}
	data := bson.M{
		"$set": bson.M{
			// "account":       form.Account,
			"password":      string(hashedPassword),
			"phone_number":  form.Phone_number,
			"email":         form.Email,
			"username":      form.Username,
			"address":       form.Address,
			"gender":        form.Gender,
			"identity_no":   form.Identity_no,
			"identity_type": form.Identity_type,
			"department":    form.Department,
			"position":      form.Position,
			"updated_at":    update_t}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert admin fault")
	}
	return err
}

//DeleteOneAdmin
func (m AdminModel) DeleteOneAdmin(adminId string) (err error) {
	var admin Admin
	err = db.GetDB().C("admin").Find(bson.M{"id": adminId}).One(&admin)
	if err != nil {
		return errors.New("Could't find the admin")
	}
	err = db.GetDB().C("admin").Remove(bson.M{"id": adminId})
	return err
}
