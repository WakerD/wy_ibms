package forms

import "wy_ibms_demo/utils"

//InsertFloorForm
type InsertFloorForm struct {
	Id string `form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
	Type string `form:"type" json:"type" binding:"required"`
	Floor int `form:"floor" json:"floor" binding:"required"`
	Image_url string `form:"image_url" json:"image_url" binding:"required"`
	Grid utils.JSONRaw	`form:"grid" json:"grid" binding:"required"`

	//Floor_grid_x int `form:"floor_grid_x" json:"floor_grid_x" binding:"required"`
	//Floor_grid_y int `form:"Floor_grid_y" json:"Floor_grid_y" binding:"required"`
}

type Grid struct{
	X int	`form:"x" json:"x" binding:"required"`
	Y int	`form:"y" json:"y" binding:"required"`
}
