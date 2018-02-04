package forms

import "wy_ibms_demo/utils"

//InsertDoorForm
type InsertDoorForm struct {
	Id string `form:"id" json:"id"`
	//门牌号
	Door_id string `form:"door_id" json:"door_id" binding:"required"`
	//门名称
	Name    string `form:"name" json:"name" binding:"required"`
	Type    string `form:"type" json:"type" binding:"required"`
	Ip      string `form:"ip" json:"ip" binding:"required"`
	Channel string `form:"channel" json:"channel" binding:"required"`
	//门的（开、关）状态
	Status string `form:"status" json:"status" binding:"required"`
	//门的在线、离线
	Online string `form:"online" json:"online" binding:"required"`
	//人脸识别图片（证件照）
	Image_url string `form:"image_url" json:"image_url" binding:"required"`

	//Location utils.JSONRaw	`form:"location" json:"location" binding:"required"`
	//Door_position utils.JSONRaw	`form:"door_position" json:"door_position" binding:"required"`

	//门详细位置
	Location_name string `form:"location_name" json:"location_name" binding:"required"`
	Floor         string `form:"floor" json:"floor" binding:"required"`

	Door_position_x int `form:"door_position_x" json:"door_position_x" binding:"required"`
	Door_position_y int `form:"door_position_y" json:"door_position_y" binding:"required"`
}

type InsertDoorLog struct {
	Id string `form:"id" json:"id"`
	//门的唯一Id
	Door_id string `form:"door_id" json:"door_id" binding:"required"`
	Door_name string `form:"door_name" json:"door_name" binding:"required"`
	//人员的唯一Id
	Admin_id string `form:"admin_id" json:"admin_id" binding:"required"`
	Admin_name string `form:"admin_name" json:"admin_name" binding:"required"`
	//组织名字
	Organization_name string `form:"organization_name" json:"organization_name" binding:"required"`
	//开门方式
	Tcm_name string `form:"tcm_name" json:"tcm_name" binding:"required"`
	//进出门的时间
	Date string        `form:"date" json:"date" binding:"required"`
	Raw  utils.JSONRaw `form:"raw" json:"raw"`
}

type DoorPageForm struct {
	Page     int   `form:"page" json:"page"`
	Pagesize int   `form:"pagesize" json:"pagesize"`
	Nums     int64 `form:"nums" json:"nums"`
}
//动态条件查询-专用实体类
type FindDoorConditionForm struct {
	//门的唯一Id
	Door_id string `form:"door_id" json:"door_id"`
	Door_name string `form:"door_name" json:"door_name"`
	//人员的唯一Id
	Admin_id string `form:"admin_id" json:"admin_id"`
	Admin_name string `form:"admin_name" json:"admin_name"`
	//组织名字
	Organization_name string `form:"organization_name" json:"organization_name"`
	//开门方式
	Tcm_name string `form:"tcm_name" json:"tcm_name"`
	//开始时间
	StartDate     string `form:"startDate" json:"startDate"`
	//结束时间
	EndDate     string `form:"endDate" json:"endDate"`
	//分页参数
	Page     int   `form:"page" json:"page"`
	Pagesize int   `form:"pagesize" json:"pagesize"`
	Nums     int64 `form:"nums" json:"nums"`
}