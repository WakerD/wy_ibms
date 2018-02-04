package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math/rand"
	"reflect"
	"time"
	// "log"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
)

//UserSessionInfo ...
type UserSessionInfo struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//UserSessionInfo ...
type AdminSessionInfo struct {
	Id            string `json:"id"`
	Account       string `json:"account"`
	Password      string `json:"password"`
	Phone_number  string `json:"phone_number"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Address       string `json:"address"`
	Gender        bool   `json:"gender"`
	Identity_no   string `json:"identity_no"`
	Identity_type string `json:"identity_type"`
	Department    string `json:"department"`
	Position      string `json:"position"`
	Updated_at    int64  `json:"updated_at"`
	Created_at    int64  `json:"created_at"`
}

//UserSessionInfo ...
type MonitorSessionInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Ip        string `json:"ip"`
	Channel   string `json:"channel"`
	Status    string `json:"status"`
	Image_url string `json:"image_url"`
	//Location   Location `json:"location"`
	//Position    Position  `json:"position"`

	Location_name string `json:"location_name" binding:"required"`
	Floor         string `json:"floor" binding:"required"`

	Monitor_position_x int `json:"monitor_position_x" binding:"required"`
	Monitor_position_y int `json:"monitor_position_y" binding:"required"`

	Updated_at int64 `json:"updated_at"`
	Created_at int64 `json:"created_at"`
}

//type Location struct {
//	Location_name string `json:"location_name"`
//	Floor         int    `json:"floor"`
//}
//
//type Position struct {
//	X int	`json:"x"`
//	Y int	`json:"y"`
//}

type FloorSessionInfo struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Floor     int     `json:"floor"`
	Image_url string  `json:"image_url"`
	Grid      JSONRaw `json:"grid"`

	//Floor_grid_x int	`json:"floor_grid_x"`
	//Floor_grid_y int `json:"floor_grid_y"`
}

type Grid struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type DepartmentSessionInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PositionSessionInfo struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Department_id string `json:"department_id"`
}

type ElectricityMeterSessionInfo struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Ip        string  `json:"ip"`
	Channel   string  `json:"channel"`
	Status    string  `json:"status"`
	Image_url string  `json:"image_url"`
	Location  JSONRaw `json:"location"`
	Position  JSONRaw `json:"position"`
}

//需要测试调整
type Location struct {
	Location_name string `form:"location_name" json:"location_name" binding:"required"`
	Floor         int    `form:"floor" json:"floor" binding:"required"`
}

//需要测试调整
type Position struct {
	X int `form:"x" json:"x" binding:"required"`
	Y int `form:"y" json:"y" binding:"required"`
}

type InsertElectricMeterLogForm struct {
	Id                string  `json:"id"`
	SlaveId           string  `json:"slaveId"`
	Type              string  `json:"type"`
	TolalActiveEnergy float64 `json:"tolalActiveEnergy"`
	Updated_at        int64   `json:"updated_at"`
	Created_at        int64   `json:"created_at"`
}

type DoorSessionInfo struct {
	Id string `json:"id"`
	//门牌号
	Door_id string `json:"door_id"`
	//门名称
	Name    string `json:"name"`
	Type    string `json:"type"`
	Ip      string `json:"ip"`
	Channel string `json:"channel"`
	//门的（开、关）状态
	Status string `json:"status"`
	//门的在线、离线
	Online string `json:"online"`
	//人脸识别图片（证件照）
	Image_url string `json:"image_url"`

	//Location utils.JSONRaw	`form:"location" json:"location" binding:"required"`
	//Door_position utils.JSONRaw	`form:"door_position" json:"door_position" binding:"required"`

	//门详细位置
	Location_name string `form:"location_name" json:"location_name" binding:"required"`
	Floor         string `form:"floor" json:"floor" binding:"required"`

	Door_position_x int `form:"door_position_x" json:"door_position_x" binding:"required"`
	Door_position_y int `form:"door_position_y" json:"door_position_y" binding:"required"`

	Updated_at int64 `json:"updated_at"`
	Created_at int64 `json:"created_at"`
}

type InsertDoorLog struct {
	Id string `json:"id"`
	//门的唯一Id
	Door_id string `json:"door_id"`
	//人员的唯一Id
	Admin_id string `json:"admin_id"`
	//进出门的时间
	Date string `json:"date"`
}

//JSONRaw ...
type JSONRaw json.RawMessage

//Value ...
func (j JSONRaw) Value() (driver.Value, error) {
	byteArr := []byte(j)
	return driver.Value(byteArr), nil
}

//Scan ...
func (j *JSONRaw) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}
	err := json.Unmarshal(asBytes, &j)
	if err != nil {
		return error(errors.New("Scan could not unmarshal to []string"))
	}
	return nil
}

//MarshalJSON ...
func (j *JSONRaw) MarshalJSON() ([]byte, error) {
	return *j, nil
}

//UnmarshalJSON ...
func (j *JSONRaw) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

//ConvertToInt64 ...
func ConvertToInt64(number interface{}) int64 {
	if reflect.TypeOf(number).String() == "int" {
		return int64(number.(int))
	}
	// var key int
	return number.(int64)
}

//生成验证码verification code
func VcodeTool() int {
	rand.Seed(time.Now().Unix())
	Rnd := rand.Intn(999999)
	//生成六位数验证码
	// Vcode := vcode()
	// fmt.Printf("rand is %v\n", rnd)
	return Rnd
}

//小端-byte转换为float32格式
func LittleByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

//小端-float32转换为byte格式
func LittleFloat32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

//小端-byte转换为float64格式
func LittleByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

//小端-float64转换为byte格式
func LittleFloat64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

//大端-byte转换为float32格式
func BigByteToFloat32(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

//大端-float32转换为byte格式
func BigFloat32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint32(bytes, bits)
	return bytes
}

//大端-byte转换为float64格式
func BigByteToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

//大端-float64转换为byte格式
func BigFloat64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)
	return bytes
}

//int64时间戳转化为日期格式
func UnixInt64ToTime(timestamp int64) string {
	timeLayout := "2006-01-02 15:04:05"
	timestr := strconv.FormatInt(timestamp, 10)
	i, err := strconv.ParseInt(timestr, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0).Format(timeLayout)
	return tm
}

//日期格式转化为int64时间戳
func TimeToUnixInt64(timestr string) int64 {
	tm, _ := time.Parse("2006-01-02 15:04:05", timestr)
	return tm.Unix() - 28800
}

func PrintValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
}

//对象转换成map
func Obj2map(obj interface{}) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}