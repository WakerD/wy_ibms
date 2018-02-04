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

type RoleController struct {}

var roleModel = new(models.RoleModel)

//Find all role
func (ctrl RoleController) AllRole(c *gin.Context) {
	data, err := roleModel.AllRole()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the roles", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//Insert one role
func (ctrl RoleController) InsertOneRole(c *gin.Context)  {
	var roleForm  forms.RoleForm

	if c.BindJSON(&roleForm) != nil {
		c.JSON(406,gin.H{"message":"Invalid form","form":roleForm})
		c.Abort()
		return
	}

	role,err := roleModel.InsertOneRole(roleForm)
	if err != nil {
		c.JSON(406,gin.H{"message":err.Error()})
		c.Abort()
		return
	}

	if role.Id != "" {
		c.JSON(200,gin.H{"message": "Success insert", "role": role})
	} else {
		c.JSON(406,gin.H{"message": "Could not insert this role", "error": err})
		c.Abort()
		return
	}
}

//Find one role by id
func (ctrl RoleController) FindOneRole(c *gin.Context) {
	roleId := c.Query("id")
	fmt.Println(roleId)
	role, err := roleModel.FindOneRole(roleId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one role", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one role", "role": role})
	}
}

//upsert one role by id
func (ctrl RoleController) UpsertRoleById(c *gin.Context) {
	var roleForm forms.RoleForm

	if c.BindJSON(&roleForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": roleForm})
		c.Abort()
		return
	}

	err := roleModel.UpsertRoleById(roleForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one role", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one role"})
	}
}

//delete ont role by id
func (ctrl RoleController) DeleteOneRole(c *gin.Context) {
	roleId := c.Query("id")
	err := roleModel.DeleteOneRole(roleId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one role", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one role", "roleId": roleId})
	}
}
