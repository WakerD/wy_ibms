package forms

//AdminSigninForm ...
type AdminSigninForm struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//AdminSignupForm ...
type AdminSignupForm struct {
	Id            string `form:"id" json:"id"`
	Account       string `form:"account" json:"account" binding:"required,max=100"`
	Password      string `form:"password" json:"password" binding:"required"`
	Phone_number  string `form:"phone_number" json:"phone_number" binding:"required"`
	Email         string `form:"email" json:"email" binding:"required"`
	Username      string `form:"username" json:"username" binding:"required"`
	Address       string `form:"address" json:"address" binding:"required"`
	Gender          int    `form:"gender" json:"gender" binding:"required"`
	Identity_no   string `form:"identity_no" json:"identity_no" binding:"required"`
	Identity_type string `form:"identity_type" json:"identity_type" binding:"required"`
	Department    string `form:"department" json:"department" binding:"required"`
	Position      string `form:"position" json:"position" binding:"required"`

	//Department_id    string `form:"department_id" json:"department_id" binding:"required"`
	//Position_id      string `form:"position_id" json:"position_id" binding:"required"`
}
