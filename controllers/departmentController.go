package controllers

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"wy_ibms_demo/forms"
	"wy_ibms_demo/models"
	"wy_ibms_demo/utils"
)

type DepartmentController struct{}

var departmentModel = new(models.DepartmentModel)

func (ctrl DepartmentController) AllDepartment(c *gin.Context) {
	data, err := departmentModel.AllDepartment()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the departments", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//getAdminID ...
func getDepartmentID(c *gin.Context) int64 {
	session := sessions.Default(c)
	departmentId := session.Get("department_id")
	if departmentId != nil {
		return utils.ConvertToInt64(departmentId)
	}
	return 0
}

//getSessionAdminInfo ...
func getSessionDepartmentInfo(c *gin.Context) (departmentSessionInfo utils.DepartmentSessionInfo) {
	session := sessions.Default(c)
	departmentId := session.Get("department_id")
	if departmentId != nil {
		departmentSessionInfo.Id = session.Get("department_id").(string)
		departmentSessionInfo.Name = session.Get("department_name").(string)
	}
	return departmentSessionInfo
}

//insert one department
func (ctrl DepartmentController) InsertOneDepartment(c *gin.Context) {
	var insertDepartmentForm forms.InsertDepartmentForm

	if c.BindJSON(&insertDepartmentForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertDepartmentForm})
		c.Abort()
		return
	}

	department, err := departmentModel.InsertOneDepartment(insertDepartmentForm)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if department.Id != "" {
		session := sessions.Default(c)
		session.Set("department_id", department.Id)
		session.Set("department_name", department.Name)
		session.Save()
		c.JSON(200, gin.H{"message": "Success insert", "department": department})
	} else {
		c.JSON(406, gin.H{"message": "Could not insert this department", "error": err})
		c.Abort()
		return
	}
}

//查询一个用户
func (ctrl DepartmentController) FindOneDepartment(c *gin.Context) {
	departmentId := c.Query("id")
	fmt.Println(departmentId)
	department, err := departmentModel.FindOneDepartment(departmentId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one department", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one department", "department": department})
	}
}

//根据设备id更新
func (ctrl DepartmentController) UpsertDepartmentById(c *gin.Context) {
	var insertDepartmentForm forms.InsertDepartmentForm

	if c.BindJSON(&insertDepartmentForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertDepartmentForm})
		c.Abort()
		return
	}

	err := departmentModel.UpsertDepartmentById(insertDepartmentForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one department", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one department"})
	}
}

//删除一个用户
func (ctrl DepartmentController) DeleteOneDepartment(c *gin.Context) {
	departmentId := c.Query("id")
	err := departmentModel.DeleteOneDepartment(departmentId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one department", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one department", "departmentId": departmentId})
	}
}