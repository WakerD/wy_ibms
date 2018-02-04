package controllers

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"wy_ibms_demo/forms"
	"wy_ibms_demo/models"
	"wy_ibms_demo/utils"
)

type AdminController struct{}

var adminModel = new(models.AdminModel)

func (ctrl AdminController) AllAdmin(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"message": "user pong",
	// })
	data, err := adminModel.AllAdmin()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the admins", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//getAdminID ...
func getAdminID(c *gin.Context) int64 {
	session := sessions.Default(c)
	adminId := session.Get("admin_id")
	if adminId != nil {
		return utils.ConvertToInt64(adminId)
	}
	return 0
}

//getSessionAdminInfo ...
func getSessionAdminInfo(c *gin.Context) (adminSessionInfo utils.AdminSessionInfo) {
	session := sessions.Default(c)
	adminId := session.Get("admin_id")
	if adminId != nil {
		adminSessionInfo.Id = session.Get("admin_id").(string)
		adminSessionInfo.Username = session.Get("admin_name").(string)
		adminSessionInfo.Email = session.Get("admin_email").(string)
	}
	return adminSessionInfo
}

//AdminSignin ...
func (ctrl AdminController) AdminSignin(c *gin.Context) {
	var adminSigninForm forms.AdminSigninForm

	if c.BindJSON(&adminSigninForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": adminSigninForm})
		c.Abort()
		return
	}

	admin, err := adminModel.AdminSignin(adminSigninForm)
	if err == nil {
		session := sessions.Default(c)
		session.Set("admin_id", admin.Id)
		session.Set("admin_email", admin.Email)
		session.Set("admin_name", admin.Username)
		session.Save()
		c.JSON(200, gin.H{"message": "Admin signed in", "admin": admin})
	} else {
		c.JSON(406, gin.H{"message": "Invalid signin details", "error": err.Error()})
		c.Abort()
		return
	}
}

//Signup ...
func (ctrl AdminController) AdminSignup(c *gin.Context) {
	var adminSignupForm forms.AdminSignupForm

	if c.BindJSON(&adminSignupForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": adminSignupForm})
		c.Abort()
		return
	}

	admin, err := adminModel.AdminSignup(adminSignupForm)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if admin.Id != "" {
		session := sessions.Default(c)
		session.Set("admin_id", admin.Id)
		session.Set("admin_email", admin.Email)
		session.Set("admin_name", admin.Username)
		session.Set("admin_password", admin.Password)
		session.Save()
		c.JSON(200, gin.H{"message": "Success signup", "admin": admin})
	} else {
		c.JSON(406, gin.H{"message": "Could not signup this admin", "error": err})
		c.Abort()
		return
	}
}

//Signout ...
func (ctrl AdminController) AdminSignout(c *gin.Context) {
	session := sessions.Default(c)
	fmt.Println("-------注销用户session---------")
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{"message": "Signed out..."})
}

//查询一个用户
func (ctrl AdminController) FindOneAdmin(c *gin.Context) {
	adminId := c.Query("id")
	fmt.Println(adminId)
	admin, err := adminModel.FindOneAdmin(adminId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one admin", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one admin", "admin": admin})
	}
}

//修改用户信息
func (ctrl AdminController) UpsertAdminById(c *gin.Context) {
	var adminSignupForm forms.AdminSignupForm

	if c.BindJSON(&adminSignupForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": adminSignupForm})
		c.Abort()
		return
	}

	err := adminModel.UpsertAdminById(adminSignupForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one admin", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one admin"})
	}
}

//删除一个用户
func (ctrl AdminController) DeleteOneAdmin(c *gin.Context) {
	adminId := c.Query("id")
	err := adminModel.DeleteOneAdmin(adminId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one admin", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one admin", "adminId": adminId})
	}
}
