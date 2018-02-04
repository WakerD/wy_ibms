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

//User ...
type User struct {
	ID          string `db:"id, primarykey, autoincrement" json:"id"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"-"`
	Name        string `db:"name" json:"name"`
	PhoneNumber string `db:"phoneNumber" json:"phoneNumber" binding:"required"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
}

//Vcode ...
type Vcode struct {
	PhoneNumber string `db:"phoneNumber" json:"phoneNumber" binding:"required"`
	Vcode       int    `db:"vcode" json:"vcode"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
}

// //admin
// type Admin struct {
// 	Id            string `db:"id, primarykey, autoincrement" json:"id"`
// 	Username      string `db:"username" json:"username"`
// 	Password      string `db:"password" json:"password"`
// 	Phone_number  string `db:"phone_number" json:"name"`
// 	Email         string `db:"email" json:"email"`
// 	Address       string `db:"address" json:"address"`
// 	Gender        bool   `db:"gender" json:"gender"`
// 	Identity_no   string `db:"identity_no" json:"identity_no"`
// 	Identity_type string `db:"identity_type" json:"identity_type"`
// 	Department    string `db:"department" json:"department"`
// 	Position      string `db:"position" json:"position"`
// 	Updated_at    int64  `db:"updated_at" json:"updated_at"`
// 	Created_at    int64  `db:"updated_at" json:"updated_at"`
// }

//UserModel ...
type UserModel struct{}

var user_t = time.Now().Unix()

//All ...
func (m UserModel) All() (users []User, err error) {
	c := db.GetDB().C("user")
	err = c.Find(nil).All(&users)
	return users, err
}

//Signin ...
func (m UserModel) Signin(form forms.SigninForm) (user User, err error) {

	err = db.GetDB().C("user").Find(bson.M{"email": form.Email}).One(&user)

	if err != nil {
		return user, err
	}

	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, errors.New("Invalid password")
	}

	return user, nil
}

//Signup ...
func (m UserModel) Signup(form forms.SignupForm) (user User, err error) {
	var (
		id = uuid.Rand()

		// data = time.Unix(1514291163, 0).Format("2006-01-02 13:04:05")
	)
	// fmt.Println(t)
	// fmt.Println(data)

	getDb := db.GetDB().C("user")

	checkUser, err := getDb.Find(bson.M{"email": form.Email}).Count()
	if err != nil {
		return user, err
	}

	if checkUser > 0 {
		return user, errors.New("User exists")
	}
	//用户密码加密
	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	//用户id用uuid
	err = getDb.Insert(
		&User{
			ID:          id.Hex(),
			Email:       form.Email,
			Password:    string(hashedPassword),
			Name:        form.Name,
			PhoneNumber: form.PhoneNumber,
			UpdatedAt:   user_t,
			CreatedAt:   user_t})

	if err == nil {
		err = getDb.Find(bson.M{"email": form.Email}).One(&user)
		if err == nil {
			return user, nil
		}
	}

	return user, errors.New("Not registered")
}

 //生成验证码
 func (m UserModel) InsertVcode(form forms.VcodeForm) (err error) {
	 now:=time.Now()
 	err = db.GetDB().C("vcode").Insert(&Vcode{
		PhoneNumber: form.Phonenumber,
 		Vcode: form.Vcode,
 		UpdatedAt: &now,
 		CreatedAt: &now})
 	return err
 }

// selector := bson.M{"_id": bson.ObjectIdHex("571de968a99cff2c68264807")}
// data := bson.M{"$set": bson.M{"age": 21}}
// err := getDB().C("user").Update(selector, data)
// if err != nil {
//     panic(err)
// }
//用户多次点击后，更新验证码
// func (m UserModel) UpdateVcode(form forms.VcodeForm) (err error) {
// 	err = db.GetDB().C("vcode").Update(bson.M{"phonenumber": form.Phonenumber}, bson.M{"$set": bson.M{"vcode": form.Vcode, "UpdatedAt": time.Now().Unix()}})
// 	return err
// }

// selector := bson.M{"key": "max"}
// data := bson.M{"$set": bson.M{"value": 30}}
// changeInfo, err := getDB().C("config").Upsert(selector, data)
// if err != nil {
//     panic(err)
// }
// fmt.Printf("%+v\n", changeInfo)
// // 首次执行output: &{Updated:0 Removed:0 UpsertedId:ObjectIdHex("571df02ea99cff2c6826480a")}
// // 再次执行output: &{Updated:1 Removed:0 UpsertedId:<nil>}

//用户多次点击后，更新验证码
func (m UserModel) UpsertVcode(form forms.VcodeForm) (err error) {

	getDb := db.GetDB().C("vcode")

	selector := bson.M{"phonenumber": form.Phonenumber}
	data := bson.M{
		"$set": bson.M{
			"vcode":     form.Vcode,
			"updatedAt": user_t}}

	_, err = getDb.Upsert(selector, data)
	if err != nil {
		return errors.New("Upsert fault")
	}

	return err
}

// // 根据用户号码查询是否生成验证码
// func (m UserModel) FindVcodeByPhonenumber(userPhonenumber string) (vcode Vcode, err error) {
// 	err = db.GetDB().C("vcode").Find(bson.M{"phonenumber": userPhonenumber}).One(&vcode)
// 	return vcode, err
// }

//One ...
func (m UserModel) FindOne(userID int64) (user User, err error) {
	err = db.GetDB().C("user").Find(bson.M{"id": userID}).One(&user)
	return user, err
}

func (m UserModel) DeleteOne(userID int64) (err error) {
	err = db.GetDB().C("user").Remove(bson.M{"id": userID})
	return err
}
