package forms

//SystemForm ...
type SystemForm struct {
	Id	string `form:"id" json:"id"`
	System_name  string `form:"system_name" json:"system_name" binding:"required"`
	//Authority  string `form:"authority" json:"authority" binding:"required"`
	Icon string `form:"icon" json:"icon" binding:"required"`
}

