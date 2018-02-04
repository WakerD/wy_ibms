package controllers

import (
	// "fmt"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"wy_ibms_demo/forms"
	"wy_ibms_demo/models"
	"wy_ibms_demo/utils"
)

type UserController struct{}

var userModel = new(models.UserModel)

func (ctrl UserController) All(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"message": "user pong",
	// })
	data, err := userModel.All()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the users", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//getUserID ...
func getUserID(c *gin.Context) int64 {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID != nil {
		return utils.ConvertToInt64(userID)
	}
	return 0
}

//getSessionUserInfo ...
func getSessionUserInfo(c *gin.Context) (userSessionInfo utils.UserSessionInfo) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	if userId != nil {
		userSessionInfo.Id = utils.ConvertToInt64(userId)
		userSessionInfo.Name = session.Get("user_name").(string)
		userSessionInfo.Email = session.Get("user_email").(string)
	}
	return userSessionInfo
}

//Signin ...
func (ctrl UserController) Signin(c *gin.Context) {
	var signinForm forms.SigninForm

	if c.BindJSON(&signinForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signinForm})
		c.Abort()
		return
	}

	user, err := userModel.Signin(signinForm)
	if err == nil {
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("user_email", user.Email)
		session.Set("user_name", user.Name)
		// session.Set("code", Vcode)
		session.Save()

		c.JSON(200, gin.H{"message": "User signed in", "user": user})
	} else {
		c.JSON(406, gin.H{"message": "Invalid signin details", "error": err.Error()})
		c.Abort()
		return
	}
}

//Signup ...
func (ctrl UserController) Signup(c *gin.Context) {
	var signupForm forms.SignupForm
	fmt.Println("-------form参数校验------------")
	if c.BindJSON(&signupForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signupForm})
		c.Abort()
		return
	}
	fmt.Println("-------model层调用------------")
	user, err := userModel.Signup(signupForm)
	fmt.Println(user)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	// fmt.Println("-------校验验证码开始------------")
	// //校验验证码
	// vcode, err := userModel.FindVcodeByPhonenumber(signupForm.Phonenumber)
	// fmt.Println(vcode)
	// if signupForm.Vcode != vcode.Vcode {
	// 	c.JSON(406, gin.H{"message": "Invalid Vcode", "form": signupForm.Vcode})
	// 	c.Abort()
	// 	return
	// }
	// if user_exist.Vcode != user_exist.Vcode {
	// 	c.JSON(406, gin.H{"message": "Invalid Vcode", "form": Vcode})
	// 	c.Abort()
	// 	return
	// }
	fmt.Println("---------判断id是否大于0----------")
	if user.ID != "" {
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("user_email", user.Email)
		session.Set("user_name", user.Name)
		session.Set("user_password", user.Password)
		// session.Set("user_vcode", Vcode)
		session.Save()
		c.JSON(200, gin.H{"message": "Success signup", "user": user})
	} else {
		c.JSON(406, gin.H{"message": "Could not signup this user", "error": err})
		c.Abort()
		return
	}
}

//Signout ...
func (ctrl UserController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{"message": "Signed out..."})
}

//生成验证码vcode...
func (ctrl UserController) InsertVcode(c *gin.Context) {
	var vcodeForm forms.VcodeForm

	phonenumber := c.PostForm("Phonenumber")
	vcode := utils.VcodeTool()

	vcodeForm.Phonenumber = phonenumber
	vcodeForm.Vcode = vcode

	// if c.BindJSON(&vcodeForm) != nil {
	// 	c.JSON(406, gin.H{"message": "Invalid form", "form": vcodeForm})
	// 	c.Abort()
	// 	return
	// }

	err := userModel.UpsertVcode(vcodeForm)

	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	// if user.ID != "" {
	// 	session := sessions.Default(c)
	// 	session.Set("user_id", user.ID)
	// 	session.Set("user_email", user.Email)
	// 	session.Set("user_name", user.Name)
	// 	session.Set("user_password", user.Password)
	// 	// session.Set("user_vcode", Vcode)
	// 	session.Save()
	// }

	c.JSON(200, gin.H{"message": "create vcode...", "form": vcodeForm})
}

//查询一个用户
func (ctrl UserController) FindOne(c *gin.Context) {
	user, err := userModel.FindOne(getUserID(c))

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one user", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one user", "user": user})
	}

}

//删除一个用户
func (ctrl UserController) DeleteOne(c *gin.Context) {
	err := userModel.DeleteOne(getUserID(c))

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one user", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one user", "userID": getUserID(c)})
	}

}
