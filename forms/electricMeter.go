package forms

import "wy_ibms_demo/utils"

type InsertElectricMeterForm struct {
	Id string	`form:"id" json:"id"`
	Name string	`form:"name" json:"name" binding:"required"`
	Type string `form:"type" json:"type" binding:"required"`
	Ip string	`form:"ip" json:"ip" binding:"required"`
	Channel string	`form:"channel" json:"channel" binding:"required"`
	Status string	`form:"status" json:"status" binding:"required"`
	Image_url string	`form:"image_url" json:"image_url" binding:"required"`
	Location utils.JSONRaw	`form:"location" json:"location" binding:"required"`
	Position utils.JSONRaw	`form:"position" json:"position" binding:"required"`
}

type Location struct {
	Location_name string	`form:"location_name" json:"location_name" binding:"required"`
	Floor int	`form:"floor" json:"floor" binding:"required"`
}

type Position struct {
	X int	`form:"x" json:"x" binding:"required"`
	Y int	`form:"y" json:"y" binding:"required"`
}

type InsertElectricMeterLogForm struct {
	Id string	`form:"id" json:"id"`
	SlaveId string	`form:"slaveId" json:"slaveId"`
	Type string		`form:"type" json:"type"`
	Current float64 `form:"current" json:"current"`
	Voltage float64 `form:"voltage" json:"voltage"`
	TolalActiveEnergy float64 `form:"tolalActiveEnergy" json:"tolalActiveEnergy"`
	Updated_at string    `form:"updated_at" json:"updated_at"`
	Created_at string    `form:"created_at" json:"created_at"`
}