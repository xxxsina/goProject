package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	getRows(1)
}

type MysqlInfo struct {
	username string
	password string
	dbase string
	charset string
}

var xgoto bool = true

func PanicErr(err error, xgoto *bool) {
	if err != nil && *xgoto {
		panic(err)
	}
}

//var dbInfo MysqlInfo= MysqlInfo{"root", "root", "ai_marketing", "utf8"}
var dbInfo MysqlInfo= MysqlInfo{"ai_marketing", "fdsafhsai2349fjhndsafhui", "ai_marketing", "utf8"}

func dbConnect() *sql.DB {
	var db, err = sql.Open("mysql", dbInfo.username+":"+dbInfo.password+"@tcp(192.168.3.237:3306)/"+dbInfo.dbase+"?charset="+dbInfo.charset)
	PanicErr(err, &xgoto)
	return db
}

func getRows(limit int) {
	stmt, err := dbConnect().Prepare("select * from am_user limit 10")
	fmt.Println(stmt)
	//PanicErr(err)
	//rows, err := stmt.Exec(limit)
	rows, err := dbConnect().Query("select id from am_demand limit ?", limit)
	PanicErr(err, &xgoto)
	xgoto = false
	for rows.Next() {
		var id int
		//var roleId int
		//var alias string
		//var username string
		err = rows.Scan(&id/*, &roleId, &username, &alias*/)
		PanicErr(err, &xgoto)
		fmt.Println("id =", id)
		//fmt.Println("roleId =", roleId)
		//fmt.Println("username =", username)
		//fmt.Println("alias =", alias)
	}
}