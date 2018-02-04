package forms

//import "wy_ibms_demo/utils"

//InsertMonitorForm
type InsertMonitorForm struct {
	Id string `form:"id" json:"id"`
	Name string	`form:"name" json:"name" binding:"required"`
	Type string	`form:"type" json:"type" binding:"required"`
	Ip string	`form:"ip" json:"ip" binding:"required"`
	Channel string	`form:"channel" json:"channel" binding:"required"`
	Status string	`form:"status" json:"status" binding:"required"`
	Image_url string `form:"image_url" json:"image_url" binding:"required"`
	//Location utils.JSONRaw	`form:"location" json:"location" binding:"required"`
	//Monitor_position utils.JSONRaw	`form:"monitor_position" json:"monitor_position" binding:"required"`

	Location_name string	`form:"location_name" json:"location_name" binding:"required"`
	Floor string	`form:"floor" json:"floor" binding:"required"`

	Monitor_position_x int	`form:"monitor_position_x" json:"monitor_position_x" binding:"required"`
	Monitor_position_y int	`form:"monitor_position_y" json:"monitor_position_y" binding:"required"`
}

//type Location struct {
//	Location_name string	`form:"location_name" json:"location_name" binding:"required"`
//	Floor int	`form:"floor" json:"floor" binding:"required"`
//}
//
//type Monitor_position struct {
//	X int	`form:"x" json:"x" binding:"required"`
//	Y int	`form:"y" json:"y" binding:"required"`
//}

type FindMonitorConditionForm struct {
	Name     string   `form:"name" json:"name"`
	Type		string   `form:"type" json:"type"`
	Ip     string `form:"ip" json:"ip"`
	Status     string `form:"status" json:"status"`
	Floor     string `form:"floor" json:"floor"`
	//分页参数
	Page     int   `form:"page" json:"page"`
	Pagesize int   `form:"pagesize" json:"pagesize"`
	Nums     int64 `form:"nums" json:"nums"`
}