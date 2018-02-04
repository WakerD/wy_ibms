package common

//page struct
type PageForm struct {
	Page     int   `form:"page" json:"page"`
	Pagesize int   `form:"pagesize" json:"pagesize"`
	Nums     int64 `form:"nums" json:"nums"`
}