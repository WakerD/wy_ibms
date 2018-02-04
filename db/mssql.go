package db

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "123456", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "172.16.1.250", "the database server")
	user          = flag.String("user", "ibms", "the database user")
)

var Mssqldb *sql.DB

func ConnectMssqlDB() *sql.DB{
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;encrypt=disable", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)

	// fmt.Printf("%v", conn.(T))
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	// defer conn.Close()

	return conn
}

//GetDB ...
// func GetMssqlDB() *MssqlConn {
// 	return mssqldb
// }

func InitMssqlDB(){
	Mssqldb = ConnectMssqlDB()
}

func GetMssqlDB() *sql.DB{
	return Mssqldb
}

// func Test() {
// 	// flag.Parse()

// 	// if *debug {
// 	// 	fmt.Printf(" password:%s\n", *password)
// 	// 	fmt.Printf(" port:%d\n", *port)
// 	// 	fmt.Printf(" server:%s\n", *server)
// 	// 	fmt.Printf(" user:%s\n", *user)
// 	// }

// 	// connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;encrypt=disable", *server, *user, *password, *port)
// 	// if *debug {
// 	// 	fmt.Printf(" connString:%s\n", connString)
// 	// }
// 	// conn, err := sql.Open("mssql", connString)
// 	// if err != nil {
// 	// 	log.Fatal("Open connection failed:", err.Error())
// 	// }
// 	// defer conn.Close()


// 	// conn.Query("SELECT * From dbo.DmsOpenRecord")

// 	// stmt, err := conn.Prepare("SELECT * From dbo.DmsOpenRecord")
// 	// if err != nil {
// 	// 	log.Fatal("Prepare failed:", err.Error())
// 	// }
// 	// defer stmt.Close()

// 	// row := stmt.QueryRow()
// 	// // var retval interface{}
// 	// cols, err := row.Columns() 
// 	// vals := make([]interface{}, len(cols))
// 	// err = row.Scan(vals...)
// 	// // var somenumber string
// 	// // var somechars string
// 	// // err = row.Scan(&somenumber, &somechars)
// 	// if err != nil {
// 	// 	log.Fatal("Scan failed:", err.Error())
// 	// }
// 	// // fmt.Printf("somenumber:%s\n", somenumber)
// 	// // fmt.Printf("somechars:%s\n", somechars)
// 	// fmt.Printf("somechars:%v\n", retval)

// 	// fmt.Printf("bye\n")

	

// 	rows, err := GetMssqlDB().Query("SELECT * From dbo.DmsOpenRecord")
	
// 	defer rows.Close()
// 	cols, err := rows.Columns()
	
// 	vals := make([]interface{}, len(cols))
// 	for i := 0; i < len(cols); i++ {
// 		vals[i] = new(interface{})
// 		// if i != 0 {
// 		// 	fmt.Print("\t")
// 		// }
// 		// fmt.Print(cols[i])
// 	}
// 	// fmt.Println()
// 	for rows.Next() {
// 		err = rows.Scan(vals...)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		for i := 0; i < len(vals); i++ {
// 			// if i != 0 {
// 			// 	fmt.Print("\t")
// 			// }
// 			utils.PrintValue(vals[i].(*interface{}))
// 		}
// 		// fmt.Println()

// 	}
// }