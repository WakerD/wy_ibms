package forms

//InsertPositionForm
type InsertPositionForm struct {
	Id string `form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
	Department_id string `form:"department_id" json:"department_id"`
}
