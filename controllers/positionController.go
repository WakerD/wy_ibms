package controllers

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"wy_ibms_demo/forms"
	"wy_ibms_demo/models"
	"wy_ibms_demo/utils"
)

type PositionController struct{}

var positionModel = new(models.PositionModel)

func (ctrl PositionController) AllPosition(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"message": "user pong",
	// })
	data, err := positionModel.AllPosition()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the positions", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//getAdminID ...
func getPositionID(c *gin.Context) int64 {
	session := sessions.Default(c)
	positionId := session.Get("position_id")
	if positionId != nil {
		return utils.ConvertToInt64(positionId)
	}
	return 0
}

//getSessionAdminInfo ...
func getSessionPositionInfo(c *gin.Context) (positionSessionInfo utils.PositionSessionInfo) {
	session := sessions.Default(c)
	positionId := session.Get("position_id")
	if positionId != nil {
		positionSessionInfo.Id = session.Get("position_id").(string)
		positionSessionInfo.Name = session.Get("position_name").(string)
	}
	return positionSessionInfo
}

//insert one position
func (ctrl PositionController) InsertOnePosition(c *gin.Context) {
	var insertPositionForm forms.InsertPositionForm

	if c.BindJSON(&insertPositionForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertPositionForm})
		c.Abort()
		return
	}

	position, err := positionModel.InsertOnePosition(insertPositionForm)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if position.Id != "" {
		session := sessions.Default(c)
		session.Set("position_id", position.Id)
		session.Set("position_name", position.Name)
		session.Save()
		c.JSON(200, gin.H{"message": "Success insert", "position": position})
	} else {
		c.JSON(406, gin.H{"message": "Could not insert this position", "error": err})
		c.Abort()
		return
	}
}

//查询一个用户
func (ctrl PositionController) FindOnePosition(c *gin.Context) {
	positionId := c.Query("id")
	fmt.Println(positionId)
	position, err := positionModel.FindOnePosition(positionId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one position", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one position", "position": position})
	}
}

//根据设备id更新
func (ctrl PositionController) UpsertPositionById(c *gin.Context) {
	var insertPositionForm forms.InsertPositionForm

	if c.BindJSON(&insertPositionForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertPositionForm})
		c.Abort()
		return
	}

	err := positionModel.UpsertPositionById(insertPositionForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one position", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one position"})
	}
}

//删除一个用户
func (ctrl PositionController) DeleteOnePosition(c *gin.Context) {
	positionId := c.Query("id")
	err := positionModel.DeleteOnePosition(positionId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one position", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one position", "positionId": positionId})
	}
}