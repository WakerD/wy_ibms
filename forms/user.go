package forms

//SigninForm ...
type SigninForm struct {
	// Name        string `form:"name" json:"name" binding:"required,max=100"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
	// Phonenumber string `form:"phonenumber" json:"phonenumber" binding:"required"`
	// Vcode       int    `form:"vcode" json:"vcode" binding:"required"`
}

//SignupForm ...
type SignupForm struct {
	Name        string `form:"name" json:"name" binding:"required,max=100"`
	Email       string `form:"email" json:"email" binding:"required,email"`
	Password    string `form:"password" json:"password" binding:"required"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" binding:"required"`
	Vcode       int    `form:"vcode" json:"vcode" binding:"required"`
}

//VcodeForm ...
type VcodeForm struct {
	Phonenumber string `form:"phoneNumber" json:"phoneNumber" binding:"required"`
	Vcode       int    `form:"vcode" bson:"time" json:"vcode" binding:"required"`
}
