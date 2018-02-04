package forms

//InsertDepartmentForm
type InsertDepartmentForm struct {
	Id string `form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
}