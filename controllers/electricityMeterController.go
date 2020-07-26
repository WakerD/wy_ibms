package controllers

import (
	//"context"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"wy_ibms_demo/forms"
	"wy_ibms_demo/models"
	"wy_ibms_demo/utils"
	//"github.com/bitly/go-simplejson"
	"github.com/garyburd/redigo/redis"
	"math"
	"strconv"
	"time"
)

type ElectricityMeterController struct{}

var electricityMeterModel = new(models.ElectricityMeterModel)

func (ctrl ElectricityMeterController) AllElectricityMeter(c *gin.Context) {
	data, err := electricityMeterModel.AllElectricity()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the electricityMeters", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//getAdminID ...
func getElectricityMeterID(c *gin.Context) int64 {
	session := sessions.Default(c)
	electricityMeterId := session.Get("electricityMeter_id")
	if electricityMeterId != nil {
		return utils.ConvertToInt64(electricityMeterId)
	}
	return 0
}

//getSessionAdminInfo ...
func getSessionElectricityMeterInfo(c *gin.Context) (electricityMeterSessionInfo utils.ElectricityMeterSessionInfo) {
	session := sessions.Default(c)
	electricityMeterId := session.Get("electricityMeter_id")
	if electricityMeterId != nil {
		electricityMeterSessionInfo.Id = session.Get("electricityMeter_id").(string)
		electricityMeterSessionInfo.Name = session.Get("electricityMeter_name").(string)
		electricityMeterSessionInfo.Type = session.Get("electricityMeter_email").(string)
	}
	return electricityMeterSessionInfo
}

//insert one ElectricityMeter
func (ctrl ElectricityMeterController) InsertOneElectricityMeter(c *gin.Context) {
	var insertElectricMeterForm forms.InsertElectricMeterForm

	if c.BindJSON(&insertElectricMeterForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertElectricMeterForm})
		c.Abort()
		return
	}

	electricityMeter, err := electricityMeterModel.InsertOneMeter(insertElectricMeterForm)
	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if electricityMeter.Id != "" {
		session := sessions.Default(c)
		session.Set("electricityMeter_id", electricityMeter.Id)
		session.Set("electricityMeter_name", electricityMeter.Name)
		session.Set("electricityMeter_type", electricityMeter.Type)
		session.Save()
		c.JSON(200, gin.H{"message": "Success insert", "electricityMeter": electricityMeter})
	} else {
		c.JSON(406, gin.H{"message": "Could not insert this electricityMeter", "error": err})
	}
}

func (ctrl ElectricityMeterController) AllElectricityMeterLog(c *gin.Context) {
	data, err := electricityMeterModel.AllMeterLog()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the electricityMetersLog", "error": err.Error()})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

//insert one electricityMeterLog erery minute，不用给前台发message
func (ctrl ElectricityMeterController) InsertOneElectricityMeterLogMin() {
	var insertElectricMeterLogForm forms.InsertElectricMeterLogForm
	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接redis数据库有误", err)
		return
	}
	timer2 := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-timer2.C:
			fmt.Println("-----electricLog定时任务开始-每分钟-----")
			for meterId := 1; meterId <= 4; meterId++ {
				fmt.Println("电表编号：%#V", meterId)

				alist := "CurrentList" + strconv.Itoa(meterId)
				vlist := "VoltageList" + strconv.Itoa(meterId)
				elist := "TotalActiveEList" + strconv.Itoa(meterId)

				a_values, _ := redis.Values(r.Do("lrange", alist, "0", "100"))
				v_values, _ := redis.Values(r.Do("lrange", vlist, "0", "100"))
				e_values, _ := redis.Values(r.Do("lrange", elist, "0", "100"))

				a_avg, err := averageValue(a_values)
				v_avg, err := averageValue(v_values)
				e_avg, err := averageValue(e_values)

				//fmt.Println("测试sum是否为NAN",a_avg,v_avg,e_avg)
				if math.IsNaN(a_avg) || math.IsNaN(v_avg) || math.IsNaN(e_avg) {
					fmt.Println("electricMeter插入的数值有NAN值,不存入mongo", a_avg, v_avg, e_avg)
				} else {
					insertElectricMeterLogForm.SubordinateId = strconv.Itoa(meterId)
					insertElectricMeterLogForm.Current = a_avg
					insertElectricMeterLogForm.Voltage = v_avg
					insertElectricMeterLogForm.TolalActiveEnergy = e_avg

					electricityMeterLog, err := electricityMeterModel.InsertOneMeterLogMin(insertElectricMeterLogForm)
					if err != nil {
						fmt.Println("electricityMeterLog插入数据库有误", electricityMeterLog)
						return
					}
				}
				//redis里面相关list清空
				v1, err := r.Do("ltrim", alist, 1, 0)
				v2, err := r.Do("ltrim", vlist, 1, 0)
				v3, err := r.Do("ltrim", elist, 1, 0)
				if err != nil {
					fmt.Println("清空redis中的list失败,电流,电压,总电量", v1, v2, v3)
					return
				} else {
					fmt.Println("清空redis中的list成功,电流,电压,总电量", v1, v2, v3)
				}
				fmt.Println("-----electricLog定时任务结束-每分钟-----")
			}
		}
	}
}

//insert one electricityMeterLog erery hour不用给前台发message
func (ctrl ElectricityMeterController) InsertOneElectricityMeterLogHour() {
	var insertElectricMeterLogForm forms.InsertElectricMeterLogForm
	//定时任务开始
	timer2 := time.NewTicker(180 * time.Second)
	for {
		select {
		case <-timer2.C:
			fmt.Println("-----electricLog按天插入-定时开始-每小时-----")
			for meterId := 1; meterId <= 4; meterId++ {
				fmt.Println("电表编号：%#V", meterId)
				insertElectricMeterLogForm.Type = "every_minute"
				//按时间段查询
				//startstamp := utils.TimeToUnixInt64(insertElectricMeterLogForm.Created_at)
				// startstamp := time.Now().Unix()
				// endstamp := startstamp + 86400
				endstamp := time.Now().Unix()
				//startstamp := endstamp - 3600
				startstamp := endstamp - 180

				electricityLogs, err := electricityMeterModel.FindMeterLogByTypeAndDate(strconv.Itoa(meterId), insertElectricMeterLogForm.Type, startstamp, endstamp)
				//fmt.Printf("electricityLogs len %d \n", len(electricityLogs))
				if err != nil {
					fmt.Println(err)
				}
				//求一个电表每天的电表参数均值
				a_sum, v_sum, e_sum, n := 0.00, 0.00, 0.00, 0.00
				for _, v := range electricityLogs {
					//insertElectricMeterLogForm.Created_at=utils.UnixInt64ToTime(v.Created_at)
					a_sum = a_sum + v.Current
					v_sum = v_sum + v.Voltage
					e_sum = e_sum + v.TolalActiveEnergy
					n++
				}

				if math.IsNaN(a_sum) || math.IsNaN(v_sum) || math.IsNaN(e_sum) {
					fmt.Println("electricityMeterLog插入的数值有误，存在NaN值,不存入mongo", a_sum, v_sum, e_sum)
				} else {
					insertElectricMeterLogForm.SubordinateId = strconv.Itoa(meterId)
					insertElectricMeterLogForm.Current = a_sum / n
					insertElectricMeterLogForm.Voltage = v_sum / n
					insertElectricMeterLogForm.TolalActiveEnergy = e_sum / (n*1000)
					//fmt.Println("测试sum是否为NAN",a_sum,v_sum,e_sum)
					fmt.Printf("每小时平均值Current: %#v\n", insertElectricMeterLogForm.Current)
					fmt.Printf("每小时平均值Voltage: %#v\n", insertElectricMeterLogForm.Voltage)
					fmt.Printf("每小时平均值TolalActiveEnergy: %#v\n", insertElectricMeterLogForm.TolalActiveEnergy)

					electricityLog, err := electricityMeterModel.InsertOneMeterLogHour(insertElectricMeterLogForm)
					if err != nil {
						fmt.Println("InsertOneElectricityMeterLogHour插入数据库有误", electricityLog)
						return
					}
				}
			}
			fmt.Println("-----electricLog按天插入-定时结束-每小时-----")
		}
	}
}

//insert one electricityMeterLog erery day不用给前台发message
func (ctrl ElectricityMeterController) InsertOneElectricityMeterLogDay() {
	var insertElectricMeterLogForm forms.InsertElectricMeterLogForm
	//定时任务开始
	timer2 := time.NewTicker(540 * time.Second)
	for {
		select {
		case <-timer2.C:
			fmt.Println("-----electricLog按天插入-定时开始-每天-----")
			for meterId := 1; meterId <= 4; meterId++ {
				fmt.Printf("电表编号：%#V\n", meterId)
				insertElectricMeterLogForm.Type = "every_minute"
				//按时间段查询
				//startstamp := utils.TimeToUnixInt64(insertElectricMeterLogForm.Created_at)
				// startstamp := time.Now().Unix()
				// endstamp := startstamp + 86400
				endstamp := time.Now().Unix()
				//正式环境
				//startstamp := endstamp - 86400
				//测试环境
				startstamp := endstamp - 540
				electricityLogs, err := electricityMeterModel.FindMeterLogByTypeAndDate(strconv.Itoa(meterId), insertElectricMeterLogForm.Type, startstamp, endstamp)
				//fmt.Printf("electricityLogs len %d \n", len(electricityLogs))
				if err != nil {
					fmt.Println(err)
				}
				//求一个电表每天的电表参数均值
				a_sum, v_sum, e_sum, n := 0.00, 0.00, 0.00, 0.00
				for _, v := range electricityLogs {
					a_sum = a_sum + v.Current
					v_sum = v_sum + v.Voltage
					e_sum = e_sum + v.TolalActiveEnergy
					n++
				}

				if math.IsNaN(a_sum) || math.IsNaN(v_sum) || math.IsNaN(e_sum) {
					fmt.Println("electricityMeterLog插入的数值有NAN值,不存入mongo", a_sum, v_sum, e_sum)
				} else {
					insertElectricMeterLogForm.SubordinateId = strconv.Itoa(meterId)
					insertElectricMeterLogForm.Current = a_sum / n
					insertElectricMeterLogForm.Voltage = v_sum / n
					insertElectricMeterLogForm.TolalActiveEnergy = e_sum / (n*1000)

					fmt.Printf("每天平均值Current: %#v\n", insertElectricMeterLogForm.Current)
					fmt.Printf("每天平均值Voltage: %#v\n", insertElectricMeterLogForm.Voltage)
					fmt.Printf("每天平均值TolalActiveEnergy: %#v\n", insertElectricMeterLogForm.TolalActiveEnergy)

					electricityLog, err := electricityMeterModel.InsertOneMeterLogDay(insertElectricMeterLogForm)
					if err != nil {
						fmt.Println("InsertOneElectricityMeterLogDay插入数据库有误", electricityLog)
						return
					}
				}
			}
			fmt.Println("-----electricLog按天插入-定时任务结束-每天-----")
		}
	}
}

//insert one electricityMeterLog erery month不用给前台发message
func (ctrl ElectricityMeterController) InsertOneElectricityMeterLogMonth() {
	var insertElectricMeterLogForm forms.InsertElectricMeterLogForm
	//定时任务开始
	timer2 := time.NewTicker(1620 * time.Second)
	for {
		select {
		case <-timer2.C:
			fmt.Println("-----electricLog按月插入-定时任务开始-每月-----")
			for meterId := 1; meterId <= 4; meterId++ {
				fmt.Println("电表编号：%#V", meterId)
				insertElectricMeterLogForm.Type = "every_minute"
				//按时间段查询
				//startstamp := utils.TimeToUnixInt64(insertElectricMeterLogForm.Created_at)
				endstamp := time.Now().Unix()
				//正式环境
				//endstamp := startstamp + 86400
				//测试环境
				startstamp := endstamp - 1620
				electricityLogs, err := electricityMeterModel.FindMeterLogByTypeAndDate(strconv.Itoa(meterId), insertElectricMeterLogForm.Type, startstamp, endstamp)
				if err != nil {
					fmt.Println(err)
				}
				//求一个电表每天的电表参数均值
				a_sum, v_sum, e_sum, n := 0.00, 0.00, 0.00, 0.00
				for _, v := range electricityLogs {
					//insertElectricMeterLogForm.Created_at=utils.UnixInt64ToTime(v.Created_at)
					a_sum = a_sum + v.Current
					v_sum = v_sum + v.Voltage
					e_sum = e_sum + v.TolalActiveEnergy
					n++
				}

				if math.IsNaN(a_sum) || math.IsNaN(v_sum) || math.IsNaN(e_sum) {
					fmt.Println("electricityMeterLog插入的数值有NAN值,不存入mongo", a_sum, v_sum, e_sum)
				} else {
					insertElectricMeterLogForm.SubordinateId = strconv.Itoa(meterId)
					insertElectricMeterLogForm.Current = a_sum / n
					insertElectricMeterLogForm.Voltage = v_sum / n
					insertElectricMeterLogForm.TolalActiveEnergy = e_sum / (n*1000)

					fmt.Printf("每月平均值Current: %#v\n", insertElectricMeterLogForm.Current)
					fmt.Printf("每月平均值Voltage: %#v\n", insertElectricMeterLogForm.Voltage)
					fmt.Printf("每月平均值TolalActiveEnergy: %#v\n", insertElectricMeterLogForm.TolalActiveEnergy)

					electricityLog, err := electricityMeterModel.InsertOneMeterLogMonth(insertElectricMeterLogForm)
					if err != nil {
						fmt.Println("InsertOneElectricityMeterLogMonth插入数据库有误", electricityLog)
						return
					}
				}
			}
			fmt.Println("-----electricLog按月插入-定时结束-每月-----")
		}
	}
}

//查询一个用户
func (ctrl ElectricityMeterController) FindOneElectricityMeter(c *gin.Context) {
	electricityMeterId := c.Query("id")
	fmt.Println(electricityMeterId)
	electricityMeter, err := electricityMeterModel.FindOneMeter(electricityMeterId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find one electricityMeter", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find one electricityMeter", "electricityMeter": electricityMeter})
	}
}

//根据设备id更新
func (ctrl ElectricityMeterController) UpsertElectricityMeterById(c *gin.Context) {
	var insertElectricityMeterForm forms.InsertElectricMeterForm

	if c.BindJSON(&insertElectricityMeterForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": insertElectricityMeterForm})
		c.Abort()
		return
	}

	err := electricityMeterModel.UpsertMeterById(insertElectricityMeterForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not upsert one electricityMeter", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success upsert one electricityMeter"})
	}
}

//删除一个用户
func (ctrl ElectricityMeterController) DeleteOneElectricityMeter(c *gin.Context) {
	electricityMeterId := c.Query("id")
	err := electricityMeterModel.DeleteOneMeter(electricityMeterId)

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not delete one electricityMeter", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success delete one electricityMeter", "electricityMeterId": electricityMeterId})
	}
}

//根据设备id，查询相应时间段电表日志
func (ctrl ElectricityMeterController) FindElectricityMeterLogByType (c *gin.Context)  {
	//meterId string, logtype string, starttime int64, endtime int64
	meterId := c.Query("subordinateid")
	logtype := c.Query("type")
	starttime, err := strconv.Atoi(c.Query("starttime"))
	endtime, err := strconv.Atoi(c.Query("endtime"))

	//strconv.ParseInt(c.Query("starttime"), 10, 64)
	//endtime, err := strconv.ParseInt(c.Query("endtime"), 10, 64)

	fmt.Println("meterId",meterId)
	fmt.Println("starttime",starttime)

	if err != nil {
		fmt.Println("timestamp error")
		c.JSON(406, gin.H{"message": "timestamp error", "error": err})
		c.Abort()
		return
	}
	electricityMeter, err := electricityMeterModel.FindMeterLogByTypeAndDate(meterId,logtype,int64(starttime),int64(endtime))

	if err != nil {
		c.JSON(406, gin.H{"message": "Could not find electricityMeter", "error": err})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{"message": "Success find electricityMeter", "electricityMeter": electricityMeter})
	}
}

//从redis里面取值，然后求平均值
func averageValue(values []interface{}) (avg float64, err error) {
	var sum, n, val float64
	sum, n = 0, 0
	for _, v := range values {
		val, err = redis.Float64(v, nil)
		if err != nil {
			fmt.Println("redis获取值失败")
			return
		}
		sum = sum + val
		n++
	}
	avg1 := sum / n
	if avg1 > 0 && avg1 < 1000 {
		fmt.Println("redis取值求得平均值：%#v\n", avg1)
		return avg1, err
	} else {
		fmt.Println("redis总电量求得平均值：%#v\n", avg1/1000)
		return avg1, err
	}
}
