package main

import (
	"fmt"
	// "net/http"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"wy_ibms_demo/controllers"
	"wy_ibms_demo/db"
	//"wy_ibms_demo/quartzs"

	"wy_ibms_demo/quartzs"
	"wy_ibms_demo/utils"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	//服务器gin开始
	r := gin.Default()

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("wy-ibms-demo-session", store))

	r.Use(CORSMiddleware())

	db.Init()

	v1 := r.Group("/v1")
	{
		//定时任务每秒redis插入数据
		//meterQuartzs := new(quartzs)
		//electricityMeter := new(quartzs.ElectricityMeterQuartz)
		//println(electricityMeter)
		//定时任务每分钟从redis取出数据，并存入mongo
		//new()

		/*** START USER ***/
		//user := new(controllers.UserController)
		//v1.GET("/user", user.All)
		//v1.POST("/user/signin", user.Signin)
		//v1.POST("/user/signup", user.Signup)
		//v1.GET("/user/signout", user.Signout)
		//v1.POST("/user/findOne", user.FindOne)
		//v1.POST("/user/deleteOne", user.DeleteOne)
		// v1.POST("/user/insertVcode", user.InsertVcode)

		/*** START Project ***/
		project := new(controllers.ProjectController)
		// v1.POST("/project", project.Create)
		v1.GET("/project", project.All)
		// v1.GET("/project/:id", project.One)
		// v1.PUT("/project/:id", project.Update)
		// v1.DELETE("/project/:id", project.Delete)

		/*** START admin ***/
		admin := new(controllers.AdminController)
		v1.GET("/admin", admin.AllAdmin)
		v1.POST("/admin/signin", admin.AdminSignin)
		v1.POST("/admin/signup", admin.AdminSignup)
		v1.GET("/admin/signout", admin.AdminSignout)
		v1.GET("/admin/findOne", admin.FindOneAdmin)
		v1.POST("/admin/upsertOne", admin.UpsertAdminById)
		v1.GET("/admin/deleteOne", admin.DeleteOneAdmin)

		/*** START monitor ***/
		monitor := new(controllers.MonitorController)
		v1.GET("/monitor", monitor.AllMonitor)
		v1.POST("/monitor/insertone", monitor.InsertOneMonitor)
		v1.GET("/monitor/findone", monitor.FindOneMonitor)
		v1.POST("/monitor/upsertbyid", monitor.UpsertMonitorById)
		v1.GET("/monitor/deleteone", monitor.DeleteOneMonitor)
		v1.POST("/monitor/findbycondition", monitor.FindMonitorByCondition)

		/*** START floor ***/
		floor := new(controllers.FloorController)
		v1.GET("/floor", floor.AllFloor)
		v1.POST("/floor/insertone", floor.InsertOneFloor)
		v1.GET("/floor/findone", floor.FindOneFloor)
		v1.POST("/floor/upsertbyid", floor.UpsertFloorById)
		v1.GET("/floor/deleteone", floor.DeleteOneFloor)

		/*** START department ***/
		department := new(controllers.DepartmentController)
		v1.GET("/department", department.AllDepartment)
		v1.POST("/department/insertone", department.InsertOneDepartment)
		v1.GET("/department/findone", department.FindOneDepartment)
		v1.POST("/department/upsertbyid", department.UpsertDepartmentById)
		v1.GET("/department/deleteone", department.DeleteOneDepartment)

		/*** START position ***/
		position := new(controllers.PositionController)
		v1.GET("/position", position.AllPosition)
		v1.POST("/position/insertone", position.InsertOnePosition)
		v1.GET("/position/findone", position.FindOnePosition)
		v1.POST("/position/upsertbyid", position.UpsertPositionById)
		v1.GET("/position/deleteone", position.DeleteOnePosition)

		/*** START door***/
		door := new(controllers.DoorController)
		v1.GET("/door/alldoor", door.AllDoor)
		v1.POST("/door/insertone", door.InsertOneDoor)
		v1.GET("/door/findone", door.FindOneDoor)
		v1.POST("/door/upsertbyid", door.UpsertDoorById)
		v1.POST("/door/upsertstatusbyid", door.UpsertDoorStatusById)
		v1.GET("/door/deleteone", door.DeleteOneDoor)
		//log
		v1.GET("/doorlog/alldoorlog", door.AllDoorLog)
		//v1.POST("/doorlog/insertone", door.InsertOneDoorLog)
		v1.GET("/doorlog/findlogbyorganizationno", door.FindDoorLogByOrganizationNo)
		v1.GET("/doorlog/findlogbytcmid", door.FindDoorLogByTcmId)
		v1.GET("/doorlog/findlogbydevno", door.FindDoorLogByDevNo)
		v1.GET("/doorlog/findlogbydate", door.FindDoorLogByDate)
		v1.GET("/doorlog/findlogbydoorid", door.FindDoorLogByDoorId)
		v1.GET("/doorlog/findlogbyadminid", door.FindDoorLogByAdminId)
		v1.GET("/doorlog/deleteone", door.DeleteOneDoorLog)

		v1.POST("/doorlog/findlogbycondition", door.FindDoorLogByCondition)
		v1.POST("/doorlog/finddoorlogbypage", door.FindDoorLogByPage)

		/*** START electricityMeter***/
		meter := new(controllers.ElectricityMeterController)
		v1.GET("/meter/allmeter", meter.AllElectricityMeter)
		v1.POST("/meter/insertone", meter.InsertOneElectricityMeter)
		v1.GET("/meter/findone", meter.FindOneElectricityMeter)
		v1.POST("/meter/upsertbyid", meter.UpsertElectricityMeterById)
		v1.GET("/meter/deleteone", meter.DeleteOneElectricityMeter)

		v1.GET("/meterlog/allmeterlog", meter.AllElectricityMeterLog)
		v1.GET("/meterlog/findlogbydate", meter.FindElectricityMeterLogByType)

		/*** START role ***/
		role := new(controllers.RoleController)
		v1.GET("/role", role.AllRole)
		v1.POST("/role/insertone", role.InsertOneRole)
		v1.GET("/role/findone", role.FindOneRole)
		v1.POST("/role/upsertbyid", role.UpsertRoleById)
		v1.GET("/role/deleteone", role.DeleteOneRole)

		/*** START system ***/
		system := new(controllers.SystemController)
		v1.GET("/system", system.AllSystem)
		v1.POST("/system/insertone", system.InsertOneSystem)
		v1.GET("/system/findone", system.FindOneSystem)
		v1.POST("/system/upsertbyid", system.UpsertSystemById)
		v1.GET("/system/deleteone", system.DeleteOneSystem)


		ws := new(utils.Websocket)
		v1.GET("/ws", func(c *gin.Context) {
			ws.Wshandler(c.Writer, c.Request)
		})
	}
	/***初始化连接数据库***/
	db.RedisInit()
	db.InitMssqlDB()

	/***定时任务***/
	doorQuartz := new(quartzs.DoorQuartz)
	//每分钟从提取富士门禁表数据
	go doorQuartz.DqTimer()
	//modbus递归轮询
	electricityMeter := new(quartzs.ElectricityMeterModbus) //new(quartzs.ElectricityMeterQuartz)
	// electricityMeter.Init()
	go electricityMeter.Timer1()
	//电表定时插入log
	electricityMeterController := new(controllers.ElectricityMeterController)
	//每分钟从redis取出数据，并存入mongo
	go electricityMeterController.InsertOneElectricityMeterLogMin()
	//每小时取出数据，并存入mongo
	go electricityMeterController.InsertOneElectricityMeterLogHour()
	//每天电表数据整理后存入mongo
	go electricityMeterController.InsertOneElectricityMeterLogDay()
	//每月电表数据整理后存入mongo
	go electricityMeterController.InsertOneElectricityMeterLogMonth()

	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}
