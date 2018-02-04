package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wy_ibms_demo/forms"
	//"wy_ibms_demo/forms/common"
	"wy_ibms_demo/models"
	//"wy_ibms_demo/utils"
	//"strconv"
)

type SystemController struct {}

var systemModel = new(models.SystemModel)

//Find all system
func (ctrl SystemController) AllSystem(c *gin.Context) {
	data, err := systemModel.AllSystem()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the systems", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//Insert one system
func (ctrl SystemController) InsertOneSystem(c *gin.Context)  {
	var systemForm  forms.SystemForm

	if c.BindJSON(&systemForm) != nil {
		c.JSON(406,gin.H{"message":"Invalid form","form":systemForm})
		c.Abort()
		return
	}

	system,err := systemModel.InsertOneSystem(systemForm)
	if err != nil {
		c.JSON(406,gin.H{"message":err.Error()})
		c.Abort()
		return
	}

	if system.Id != "" {
		c.JSON(200,gin.H{"message": "Success insert", "system": system})
	} else {
		c.JSON(406,gin.H{"message": "Could not insert this system", "error": err})
		c.Abort()
		return
	}
}

//Find one system by id
func (ctrl SystemController) FindOneSystem(c *gin.Context) {
	systemId := c.Query("id")
	fmt.Println(systemId)
	system, err := systemModel.FindOneSystem(systemId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one system", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one system", "system": system})
	}
}

//upsert one system by id
func (ctrl SystemController) UpsertSystemById(c *gin.Context) {
	var systemForm forms.SystemForm

	if c.BindJSON(&systemForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": systemForm})
		c.Abort()
		return
	}

	err := systemModel.UpsertSystemById(systemForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one system", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one system"})
	}
}

//delete ont system by id
func (ctrl SystemController) DeleteOneSystem(c *gin.Context) {
	systemId := c.Query("id")
	err := systemModel.DeleteOneSystem(systemId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one system", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one system", "systemId": systemId})
	}
}

