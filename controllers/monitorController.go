package controllers

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"wy_ibms_demo/forms"
	"wy_ibms_demo/models"
	"wy_ibms_demo/utils"
	//"github.com/bitly/go-simplejson"
	//"strconv"
)

type MonitorController struct{}

var monitorModel = new(models.MonitorModel)

func (ctrl MonitorController) AllMonitor(c *gin.Context) {
	data, err := monitorModel.AllMonitor()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the monitors", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//getAdminID ...
func getMonitorID(c *gin.Context) int64 {
	session := sessions.Default(c)
	monitorId := session.Get("monitor_id")
	if monitorId != nil {
		return utils.ConvertToInt64(monitorId)
	}
	return 0
}

//getSessionAdminInfo ...
func getSessionMonitorInfo(c *gin.Context) (monitorSessionInfo utils.MonitorSessionInfo) {
	session := sessions.Default(c)
	monitorId := session.Get("monitor_id")
	if monitorId != nil {
		monitorSessionInfo.Id = session.Get("monitor_id").(string)
		monitorSessionInfo.Name = session.Get("monitor_name").(string)
		monitorSessionInfo.Type = session.Get("monitor_email").(string)
	}
	return monitorSessionInfo
}

//insert one monitor
func (ctrl MonitorController) InsertOneMonitor(c *gin.Context) {
	var insertMonitorForm forms.InsertMonitorForm

	if c.BindJSON(&insertMonitorForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertMonitorForm})
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
	//var Monitor_position forms.Monitor_position
	//
	//if c.BindJSON(&Monitor_position) != nil {
	//	c.JSON(406, gin.H{"message": "Invalid form", "form": Monitor_position})
	//	c.Abort()
	//	return
	//}

	monitor, err := monitorModel.InsertOneMonitor(insertMonitorForm)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if monitor.Id != "" {
		session := sessions.Default(c)
		session.Set("monitor_id", monitor.Id)
		session.Set("monitor_name", monitor.Name)
		session.Set("monitor_type", monitor.Type)
		session.Save()
		c.JSON(200, gin.H{"message": "Success insert", "monitor": monitor})
	} else {
		c.JSON(406, gin.H{"message": "Could not insert this monitor", "error": err})
		c.Abort()
		return
	}
}

//查询一个用户
func (ctrl MonitorController) FindOneMonitor(c *gin.Context) {
	monitorId := c.Query("id")
	fmt.Println(monitorId)
	monitor, err := monitorModel.FindOneMonitor(monitorId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one monitor", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one monitor", "monitor": monitor})
	}
}

//根据设备id更新
func (ctrl MonitorController) UpsertMonitorById(c *gin.Context) {
	var insertMonitorForm forms.InsertMonitorForm

	if c.BindJSON(&insertMonitorForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertMonitorForm})
		c.Abort()
		return
	}

	err := monitorModel.UpsertMonitorById(insertMonitorForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one monitor", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one monitor"})
	}
}

//删除一个用户
func (ctrl MonitorController) DeleteOneMonitor(c *gin.Context) {
	monitorId := c.Query("id")
	err := monitorModel.DeleteOneMonitor(monitorId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one monitor", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one monitor", "monitorId": monitorId})
	}
}

//根据传参动态查询门禁日志
func (ctrl MonitorController) FindMonitorByCondition(c *gin.Context) {
	var findConditionForm forms.FindMonitorConditionForm

	if c.BindJSON(&findConditionForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": findConditionForm})
		c.Abort()
		return
	}

	////求日志总数
	//sum,err:=monitorModel.MonitorCount()
	//if err != nil {
	//	c.JSON(406, gin.H{"message":"Invalid count", "count":sum})
	//	c.Abort()
	//	return
	//}
	//sumstr:=strconv.Itoa(sum)
	//sum64, err := strconv.ParseInt(sumstr, 10, 64)
	//findConditionForm.Nums=sum64
	//
	//paginatorMap := utils.Paginator(findConditionForm.Page,findConditionForm.Pagesize,findConditionForm.Nums)

	data,paginatorMap, err := monitorModel.FindMonitorByCondition(findConditionForm)
	//for k, v := range data {
	//	fmt.Println("key值",k)
	//	fmt.Println("value值",v)
	//}
	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the Monitors by dynamic condition", "error": err.Error()})
		c.Abort()
		return
	} else {
		//paginatorMap
		c.JSON(200, gin.H{"data": data,"paginatormap":paginatorMap})
	}
}
