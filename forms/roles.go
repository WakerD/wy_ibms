package forms

//import "wy_ibms_demo/utils"

//AuthForm ...
type RoleForm struct {
	Id	string `form:"id" json:"id"`
	Role_name  string `form:"role_name" json:"role_name" binding:"required"`
	Auth map[string]interface{} `form:"auth" json:"auth" binding:"required"`
}
