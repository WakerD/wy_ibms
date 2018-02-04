package controllers

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"wy_ibms_demo/forms"
	"wy_ibms_demo/models"
	"wy_ibms_demo/utils"
	//"github.com/bitly/go-simplejson"
	//"go/types"
	"strconv"
)

type DoorController struct{}

var doorModel = new(models.DoorModel)

func (ctrl DoorController) AllDoor(c *gin.Context) {
	data, err := doorModel.AllDoor()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doors", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//getAdminID ...
func getDoorID(c *gin.Context) int64 {
	session := sessions.Default(c)
	doorId := session.Get("door_id")
	if doorId != nil {
		return utils.ConvertToInt64(doorId)
	}
	return 0
}

//getSessionAdminInfo ...
func getSessionDoorInfo(c *gin.Context) (doorSessionInfo utils.DoorSessionInfo) {
	session := sessions.Default(c)
	doorId := session.Get("door_id")
	if doorId != nil {
		doorSessionInfo.Id = session.Get("door_id").(string)
		doorSessionInfo.Name = session.Get("door_name").(string)
		doorSessionInfo.Type = session.Get("door_email").(string)
	}
	return doorSessionInfo
}

//insert one door
func (ctrl DoorController) InsertOneDoor(c *gin.Context) {
	var insertDoorForm forms.InsertDoorForm

	if c.BindJSON(&insertDoorForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertDoorForm})
		c.Abort()
		return
	}

	//var Location forms.Location
	//
	//if c.BindJSON(&Location) != nil {
	//	c.JSON(406, gin.H{"message": "Invalid form", "form": Location})
	//	c.Abort()
	//	return
	//}
	//
	//var Door_position forms.Door_position
	//
	//if c.BindJSON(&Door_position) != nil {
	//	c.JSON(406, gin.H{"message": "Invalid form", "form": Door_position})
	//	c.Abort()
	//	return
	//}

	door, err := doorModel.InsertOneDoor(insertDoorForm)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if door.Id != "" {
		session := sessions.Default(c)
		session.Set("door_id", door.Id)
		session.Set("door_name", door.Name)
		session.Set("door_type", door.Type)
		session.Save()
		c.JSON(200, gin.H{"message": "Success insert", "door": door})
	} else {
		c.JSON(406, gin.H{"message": "Could not insert this door", "error": err})
		c.Abort()
		return
	}
}

//查询一个用户
func (ctrl DoorController) FindOneDoor(c *gin.Context) {
	doorId := c.Query("id")
	fmt.Println(doorId)
	door, err := doorModel.FindOneDoor(doorId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one door", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one door", "door": door})
	}
}

//根据设备id更新
func (ctrl DoorController) UpsertDoorById(c *gin.Context) {
	var insertDoorForm forms.InsertDoorForm

	if c.BindJSON(&insertDoorForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertDoorForm})
		c.Abort()
		return
	}

	err := doorModel.UpsertDoorById(insertDoorForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one door", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one door"})
	}
}

//根据设备id更新状态
func (ctrl DoorController) UpsertDoorStatusById(c *gin.Context) {
	var insertDoorForm forms.InsertDoorForm

	if c.BindJSON(&insertDoorForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertDoorForm})
		c.Abort()
		return
	}

	err := doorModel.UpsertDoorStatusById(insertDoorForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one door status", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one door status"})
	}
}

//删除一个门禁
func (ctrl DoorController) DeleteOneDoor(c *gin.Context) {
	Id := c.Query("id")
	err := doorModel.DeleteOneDoor(Id)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one door", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one door", "doorId": Id})
	}
}

//insert one doorLog
func (ctrl DoorController) InsertOneDoorLog(c *gin.Context) {
	var insertDoorLog forms.InsertDoorLog

	if c.BindJSON(&insertDoorLog) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertDoorLog})
		c.Abort()
		return
	}

	//var Location forms.Location
	//
	//if c.BindJSON(&Location) != nil {
	//	c.JSON(406, gin.H{"message": "Invalid form", "form": Location})
	//	c.Abort()
	//	return
	//}
	//
	//var Door_position forms.Door_position
	//
	//if c.BindJSON(&Door_position) != nil {
	//	c.JSON(406, gin.H{"message": "Invalid form", "form": Door_position})
	//	c.Abort()
	//	return
	//}

	doorLog, err := doorModel.InsertOneDoorLog(insertDoorLog)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if doorLog.Id != "" {
		session := sessions.Default(c)
		session.Set("id", doorLog.Id)
		session.Set("door_id", doorLog.Door_id)
		session.Set("admin_id", doorLog.Admin_id)
		session.Set("date", doorLog.Date)
		session.Save()
		c.JSON(200, gin.H{"message": "Success insert", "doorLog": doorLog})
	} else {
		c.JSON(406, gin.H{"message": "Could not insert this doorLog", "error": err})
		c.Abort()
		return
	}
}

//查询所有门禁日志
func (ctrl DoorController) AllDoorLog(c *gin.Context) {
	data, err := doorModel.AllDoorLog()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doorLogs", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//根据富士字段查询门禁记录-富士xx字段
func (ctrl DoorController) FindDoorLogByOrganizationNo(c *gin.Context) {
	organizationNo := c.Query("organizationNo")
	fmt.Println(organizationNo)
	log, err := doorModel.FindDoorLogByOrganizationNo(organizationNo)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doorLogs", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"log": log})
	}
}

//根据富士字段查询门禁记录-富士开门方式字段
func (ctrl DoorController) FindDoorLogByTcmId(c *gin.Context) {
	tcmId := c.Query("tcmId")
	fmt.Println(tcmId)
	log, err := doorModel.FindDoorLogByTcmId(tcmId)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doorLogs", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"log": log})
	}
}

//根据富士字段查询门禁记录-富士部门字段
func (ctrl DoorController) FindDoorLogByDevNo(c *gin.Context) {
	devNo := c.Query("devNo")
	fmt.Println(devNo)
	log, err := doorModel.FindDoorLogByTcmId(devNo)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doorLogs", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"log": log})
	}
}

//查询门禁记录-根据日期
func (ctrl DoorController) FindDoorLogByDate(c *gin.Context) {
	date := c.Query("date")
	fmt.Println(date)
	log, err := doorModel.FindDoorLogByDate(date)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doorLogs", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"log": log})
	}
}

//根据门牌号查询门禁日志
func (ctrl DoorController) FindDoorLogByDoorId(c *gin.Context) {
	doorId := c.Query("doorid")
	fmt.Println(doorId)

	data, err := doorModel.FindDoorLogByDoorId(doorId)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doorLogs by doorId", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//根据用户id查询门禁日志
func (ctrl DoorController) FindDoorLogByAdminId(c *gin.Context) {
	adminId := c.Query("adminid")
	fmt.Println(adminId)

	data, err := doorModel.FindDoorLogByAdminId(adminId)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doorLogs by adminId", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//根据传参动态查询门禁日志
func (ctrl DoorController) FindDoorLogByCondition(c *gin.Context) {
	var findConditionForm forms.FindDoorConditionForm

	if c.BindJSON(&findConditionForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": findConditionForm})
		c.Abort()
		return
	}

	data,paginatorMap, err := doorModel.FindDoorLogByCondition(findConditionForm)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doorLogs by dynamic condition", "error": err.Error()})
		c.Abort()
		return
	} else {
		//paginatorMap
		c.JSON(200, gin.H{"data": data,"paginatormap":paginatorMap})
	}
}

//分页查询所有门禁记录
func (ctrl DoorController) FindDoorLogByPage(c *gin.Context) {
	var doorPageForm forms.DoorPageForm

	if c.BindJSON(&doorPageForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": doorPageForm})
		c.Abort()
		return
	}
	//res := models.Paginator(pa, pre_page, totals)
	//this.Data["paginator"] = res

	//求日志总数
	sum,err:=doorModel.LogCount()
	if err != nil {
		c.JSON(406, gin.H{"message":"Invalid count", "count":sum})
		c.Abort()
		return
	}
	sumstr:=strconv.Itoa(sum)
	sum64, err := strconv.ParseInt(sumstr, 10, 64)
	doorPageForm.Nums=sum64

	paginatorMap := utils.Paginator(doorPageForm.Page,doorPageForm.Pagesize,doorPageForm.Nums)

	data, err := doorModel.FindDoorLogByPage(paginatorMap,doorPageForm.Pagesize)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the doorLogs by adminId", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data,"paginatormap":paginatorMap})
	}
}


//删除一条门禁日志
func (ctrl DoorController) DeleteOneDoorLog(c *gin.Context) {
	doorId := c.Query("id")
	err := doorModel.DeleteOneDoorLog(doorId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one doorLog", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one doorLog", "doorId": doorId})
	}
}
