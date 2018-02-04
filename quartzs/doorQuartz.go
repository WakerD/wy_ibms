package quartzs

import (
	"fmt"
	"time"
	//"math"
	"wy_ibms_demo/db"
	// "wy_ibms_demo/utils"
	"github.com/snluu/uuid"

	"database/sql"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"strconv"
	"wy_ibms_demo/models"
)

type DoorQuartz struct{}

func (dq DoorQuartz) DqTimer() {
	//设置定时任务的时间
	timer1 := time.NewTicker(60 * time.Second)
	//floor := new(controllers.FloorController)
	for {
		select {
		case <-timer1.C:
			// for {
			fmt.Println("-----提取门禁数据-定时任务开始-----")
			//testTimer2()

			row, err := db.GetMssqlDB().Query("SELECT COUNT(*) as count FROM dbo.DmsOpenRecord")
			// fmt.Println("Total count:",checkCount(row))
			if err != nil {
				// return err
			}
			count := checkCount(row)
			c := db.GetDB().C("doorLog")
			mCount, err := c.Find(nil).Count()
			if err != nil {
				// return err
			}
			fmt.Println("Total count:", count, mCount)

			if mCount < count {
				// 	rows, err := db.GetMssqlDB().Query("SELECT * From dbo.DmsOpenRecord")

				// defer rows.Close()
				// cols, err := rows.Columns()

				// vals := make([]interface{}, len(cols))
				// for i := 0; i < len(cols); i++ {
				// 	vals[i] = new(interface{})
				// 	// if i != 0 {
				// 	// 	fmt.Print("\t")
				// 	// }
				// 	// fmt.Print(cols[i])
				// }
				// // fmt.Println()
				// for rows.Next() {
				// 	err = rows.Scan(vals...)
				// 	if err != nil {
				// 		fmt.Println(err)
				// 		continue
				// 	}
				// 	for i := 0; i < len(vals); i++ {
				// 		// if i != 0 {
				// 		// 	fmt.Print("\t")
				// 		// }
				// 		// utils.PrintValue(vals[i].(*interface{}))

				// 	}
				// 		matchColumn(vals, cols)

				// 	// fmt.Println()

				// }

				// returns rows *sql.Rows
				// rows, err := db.GetMssqlDB().Query("SELECT * From dbo.DmsOpenRecord")

				cmd := "SELECT TOP " + strconv.Itoa(count-mCount) + " *, PicAddr From dbo.DmsOpenRecord o CROSS APPLY (SELECT p.PicAddr + ',' FROM dbo.DmsOpenPic p Where p.OpenPicNo = o.OpenPicNo For XML PATH('') ) D (PicAddr) order by OpenDate DESC"

				// j, err := getJSON("SELECT TOP "+ strconv.Itoa(count - mCount) +" * From dbo.DmsOpenRecord order by OpenDate DESC")
				// if err != nil{

				// }

				rows, err := db.GetMssqlDB().Query(cmd)
				if err != nil {
					fmt.Println("mssql query fault",rows,err)
				}
				defer rows.Close()
				//查询表中列名
				cols, err := rows.Columns()
				if err != nil {
					fmt.Println("column query fault",cols,err)
				}

				//fmt.Println("rows中的列名",cols)

				count := len(cols)
				//fmt.Println("rows中的列名总数",count)
				values := make([]interface{}, count)
				valuePtrs := make([]interface{}, count)

				n:=0

				for rows.Next() {

					n++

					for i := 0; i < count; i++ {
						valuePtrs[i] = &values[i]
					}
					rows.Scan(valuePtrs...)

					var door_id interface{}
					var door_name interface{}
					var admin_id interface{}
					var admin_name interface{}
					var organization_name interface{}
					var tcm_name interface{}
					var open_date interface{}

					entry := make(map[string]interface{})
					for i := 0; i < count; i++ {
						// fmt.Print(cols[i])
						// fmt.Println(values[i])

						if cols[i] == "DevNo" {
							door_id = values[i]
						} else if cols[i] == "DevName" {
							door_name = values[i]
						} else if cols[i] == "StaffNo" {
							admin_id = values[i]
						} else if cols[i] == "StaffName" {
							admin_name = values[i]
						} else if cols[i] == "OrganizationName" {
							organization_name = values[i]
						} else if cols[i] == "TcmName" {
							tcm_name = values[i]
						} else if cols[i] == "OpenDate" {
							open_date = values[i]
						}

						var v interface{}
						val := values[i]
						b, ok := val.([]byte)
						if ok {
							v = string(b)
						} else {
							v = val
						}
						entry[cols[i]] = v
					}
					jsonData, err := json.Marshal(entry)

					var door_t = time.Now().Unix()
					var doorid = uuid.Rand()
					c := db.GetDB().C("doorLog")

					err = c.Insert(
						&models.DoorLog{
							Id:         doorid.Hex(),
							Door_id:    strconv.Itoa(int(door_id.(int64))),
							Door_name:    door_name.(string),
							Admin_id:   admin_id.(string),
							Admin_name:   admin_name.(string),
							Organization_name:       organization_name.(string),
							Tcm_name:       tcm_name.(string),
							//时间差8小时，需要转化
							Date:       strconv.Itoa(int(open_date.(time.Time).Unix()-28800)),
							Raw:        jsonData,
							Updated_at: door_t,
							Created_at: door_t})
					var doorLog models.DoorLog
					if err == nil {
						err = c.Find(bson.M{"id": doorid.Hex()}).One(&doorLog)
						if err != nil {
							 fmt.Println("can't find doorLog",err)
						}
					}else {
						fmt.Println("doorLog insert fault",err)
					}
					fmt.Println("循环次数：",n)
				}
				//    fmt.Print(j)
				// return err
			}
			fmt.Println("-----提取门禁数据-定时任务结束-----")
			// }
		}
	}
}

// func getJSON(sqlString string) (utils.JSONRaw, error) {
//   rows, err := db.GetMssqlDB().Query(sqlString)
//   if err != nil {
//       return nil, err
//   }
//   defer rows.Close()
//   columns, err := rows.Columns()
//   if err != nil {
//       return nil, err
//   }
//   count := len(columns)
//   tableData := make([]map[string]interface{}, 0)
//   values := make([]interface{}, count)
//   valuePtrs := make([]interface{}, count)
//   for rows.Next() {
//       for i := 0; i < count; i++ {
//           valuePtrs[i] = &values[i]
//       }
//       rows.Scan(valuePtrs...)
//       entry := make(map[string]interface{})
//       for i, col := range columns {
//           var v interface{}
//           val := values[i]
//           b, ok := val.([]byte)
//           if ok {
//               v = string(b)
//           } else {
//               v = val
//           }
//           entry[col] = v
//       }
//       tableData = append(tableData, entry)
//   }
//   jsonData, err := json.Marshal(tableData)
//   if err != nil {
//       return nil, err
//   }
//   fmt.Println(string(jsonData))
//   return jsonData, nil
// }

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		// checkErr(err)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return count
}
