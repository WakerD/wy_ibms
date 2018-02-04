package quartzs

import (
	"fmt"
	"time"
	//"wy_ibms_demo/controllers"
	"github.com/goburrow/modbus"
	"log"
	"os"
	//"math"
	"wy_ibms_demo/db"
	//"wy_ibms_demo/utils"
	// "github.com/garyburd/redigo/redis"
	//"reflect"
	//"bytes"
	"strconv"
	// "wy_ibms_demo/utils"
	"encoding/binary"
	// "encoding/json"
	// "github.com/gorilla/websocket"
	"wy_ibms_demo/utils"
)

type ElectricityMeterModbus struct{}

type Message struct {
	Value   float32
	Type    string
	SlaveId int
}

type MessageEnergy struct {
	Value   uint64
	Type    string
	SlaveId int
}

//电表modbus参数查询并存入redis定时任务
func (em ElectricityMeterModbus) Timer1() {
	fmt.Println("-----modbus参数插入redis-开始-----")
	for meterId := 1; meterId <= 4; meterId++ {
		// Modbus TCP CONNECT
		handler := modbus.NewTCPClientHandler("172.16.1.190:26")
		handler.Timeout = 10 * time.Second
		handler.SlaveId = byte(meterId)
		handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
		// Connect manually so that multiple requests are handled in one connection session
		err := handler.Connect()
		if err != nil {
			fmt.Println("CONNECTION ERROR",err)
		}
		client := modbus.NewClient(handler)
		//从modbus里面取值	电流、电压、总电量
		current, err := client.ReadHoldingRegisters(3009, 2)
		voltage, err := client.ReadHoldingRegisters(3025, 2)
		energy, err := client.ReadHoldingRegisters(3203, 4)
		defer handler.Close()
		//fmt.Println("电流byte：", current)
		if err == nil {
			//redis connect
			r := db.GetRedisPool().Get()
			alist := "CurrentList" + strconv.Itoa(meterId)
			vlist := "VoltageList" + strconv.Itoa(meterId)
			elist := "TotalActiveEList" + strconv.Itoa(meterId)

			fmt.Println("电表编号：%#v", meterId,)
			if current != nil {
				Ea := utils.BigByteToFloat32(current)
				fmt.Println("电流%#V", Ea)
				a, err := r.Do("lpush", alist, Ea)
				//给redis里面存	电流、电压、总电量
				if err != nil {
					fmt.Println("redis insert false:",a,err)
					return
				}
				//电流信息用Message接收后，写入websocket通道
				m := Message{Ea, "current", meterId}
				if utils.GetWs() != nil {
					utils.GetWs().WriteJSON(m)
				}
			}

			if voltage != nil {
				Ev := utils.BigByteToFloat32(voltage)
				fmt.Printf("电压%#V\n", Ev)
				v, err := r.Do("lpush", vlist, Ev)
				//给redis里面存	电流、电压、总电量
				if err != nil {
					fmt.Println("redis insert false:",v,err)
					return
				}
				//电压信息用Message接收后，写入websocket通道
				m := Message{Ev, "voltage", meterId}
				if utils.GetWs() != nil {
					utils.GetWs().WriteJSON(m)
				}
			}

			if energy != nil {
				Ewh := binary.BigEndian.Uint64(energy)
				fmt.Printf("总电量%#V\n", Ewh)
				e, err := r.Do("lpush", elist, Ewh)
				//给redis里面存	电流、电压、总电量
				if err != nil {
					fmt.Println("redis insert false:",e, err)
					return
				}
				//总电量信息用Message接收后，写入websocket通道
				m := MessageEnergy{Ewh, "energy", meterId}
				if utils.GetWs() != nil {
					utils.GetWs().WriteJSON(m)
				}
			}
			r.Close()
		} else {
			fmt.Println("ERROR",err)
		}
	}
	fmt.Println("-----modbus参数插入redis-结束-----")
	//延时，为了保证数据稳定
	time.Sleep(time.Millisecond * 8000)
	//递归调用自身函数进行循环
	em.Timer1()
}
