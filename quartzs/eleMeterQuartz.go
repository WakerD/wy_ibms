package quartzs

import (
	"time"
	"fmt"
	//"wy_ibms_demo/controllers"
	"log"
	"os"
	"github.com/goburrow/modbus"
	//"math"
	//"wy_ibms_demo/db"
	//"wy_ibms_demo/utils"
	"github.com/garyburd/redigo/redis"
	//"reflect"
	//"bytes"
	"strconv"
	"wy_ibms_demo/utils"
	"encoding/binary"
)

type ElectricityMeterQuartz struct {}

//电表modbus参数查询并存入redis定时任务，暂未用
func (qz ElectricityMeterQuartz) Timer1()  {
	//设置定时任务的时间
	timer1:=time.NewTicker(1 * time.Second)
	meterId := 0
	for{
		select {
			case <-timer1.C:
			fmt.Println("-----modbus插入redis-定时开始-----")
			meterId=meterId%4+1
			fmt.Println("电表编号：%#V",meterId)
			// Modbus TCP CONNECT
			handler := modbus.NewTCPClientHandler("172.16.1.190:26")
			handler.Timeout = 10 * time.Second
			handler.SlaveId = byte(meterId)
			handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
			// Connect manually so that multiple requests are handled in one connection session
			err := handler.Connect()
			if err != nil{
				fmt.Println("CONNECTION ERROR")
			}
			client := modbus.NewClient(handler)
			//从modbus里面取值	电流、电压、总电量
			current, err := client.ReadHoldingRegisters(3009, 2)
			voltage, err := client.ReadHoldingRegisters(3035, 2)
			energy, err := client.ReadHoldingRegisters(3203, 4)
			if err == nil {
				Ea := utils.BigByteToFloat32(current)
				Ev := utils.BigByteToFloat32(voltage)
				Ewh := binary.BigEndian.Uint64(energy)
				//fmt.Printf("当前电流%#V\n", Ea)
				//fmt.Printf("当前电压%#V\n", Ev)
				//fmt.Printf("总电量%#V\n", Ewh)
				r,err:=redis.Dial("tcp","127.0.0.1:6379")
				defer r.Close()
				alist:="CurrentList"+strconv.Itoa(meterId)
				vlist:="VoltageList"+strconv.Itoa(meterId)
				elist:="TotalActiveEList"+strconv.Itoa(meterId)
				//给redis里面存	电流、电压、总电量
				a,err:=r.Do("lpush",alist,Ea)
				v,err:=r.Do("lpush",vlist,Ev)
				e,err:=r.Do("lpush",elist,Ewh)
				if err!=nil {
					fmt.Println("redis insert false:",err)
					return
				}
				fmt.Println("电流",a,"电压",v,"总电量",e)
			} else {
				fmt.Println("ERROR")
			}
			defer handler.Close()
			fmt.Println("-----modbus插入redis-定时结束-----")
		}
	}
}

func (qz ElectricityMeterQuartz) Timer2()  {
	//设置定时任务的时间
	timer1:=time.NewTicker(1 * time.Minute)
	//floor := new(controllers.FloorController)
	for{
		select {
		case <-timer1.C:
			for {
				fmt.Println("-----定时任务开始-----")
				//testTimer2()
				fmt.Println("-----定时任务结束-----")
			}
		}
	}
}