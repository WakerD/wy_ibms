package controllers

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"wy_ibms_demo/forms"
	"wy_ibms_demo/models"
	"wy_ibms_demo/utils"
)

type FloorController struct{}

var floorModel = new(models.FloorModel)

func (ctrl FloorController) AllFloor(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"message": "user pong",
	// })
	data, err := floorModel.AllFloor()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the floors", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//getAdminID ...
func getFloorID(c *gin.Context) int64 {
	session := sessions.Default(c)
	floorId := session.Get("floor_id")
	if floorId != nil {
		return utils.ConvertToInt64(floorId)
	}
	return 0
}

//getSessionAdminInfo ...
func getSessionFloorInfo(c *gin.Context) (floorSessionInfo utils.FloorSessionInfo) {
	session := sessions.Default(c)
	floorId := session.Get("floor_id")
	if floorId != nil {
		floorSessionInfo.Id = session.Get("floor_id").(string)
		floorSessionInfo.Name = session.Get("floor_name").(string)
		floorSessionInfo.Type = session.Get("floor_email").(string)
	}
	return floorSessionInfo
}

//insert one floor
func (ctrl FloorController) InsertOneFloor(c *gin.Context) {
	var insertFloorForm forms.InsertFloorForm

	if c.BindJSON(&insertFloorForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertFloorForm})
		c.Abort()
		return
	}

	floor, err := floorModel.InsertOneFloor(insertFloorForm)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if floor.Id != "" {
		session := sessions.Default(c)
		session.Set("floor_id", floor.Id)
		session.Set("floor_name", floor.Name)
		session.Set("floor_type", floor.Type)
		session.Save()
		c.JSON(200, gin.H{"message": "Success insert", "floor": floor})
	} else {
		c.JSON(406, gin.H{"message": "Could not insert this floor", "error": err})
		c.Abort()
		return
	}
}

//查询一个用户
func (ctrl FloorController) FindOneFloor(c *gin.Context) {
	floorId := c.Query("id")
	fmt.Println(floorId)
	floor, err := floorModel.FindOneFloor(floorId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one floor", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one floor", "floor": floor})
	}
}

//根据设备id更新
func (ctrl FloorController) UpsertFloorById(c *gin.Context) {
	var insertFloorForm forms.InsertFloorForm

	if c.BindJSON(&insertFloorForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertFloorForm})
		c.Abort()
		return
	}

	err := floorModel.UpsertFloorById(insertFloorForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one floor", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one floor"})
	}
}

//删除一个用户
func (ctrl FloorController) DeleteOneFloor(c *gin.Context) {
	floorId := c.Query("id")
	err := floorModel.DeleteOneFloor(floorId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one floor", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one floor", "floorId": floorId})
	}
}