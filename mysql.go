package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var Id int64 = 1

func main() {
	http.HandleFunc("/", sayHello2)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
		return
	}
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

func getRows(limit int64) int64 {
	stmt, err := dbConnect().Prepare("select * from am_user limit 10")
	fmt.Println(stmt)
	//PanicErr(err)
	//rows, err := stmt.Exec(limit)
	rows, err := dbConnect().Query("select id from am_demand limit ?", limit)
	PanicErr(err, &xgoto)
	xgoto = false
	var Reid int
	for rows.Next() {
		var id int
		//var roleId int
		//var alias string
		//var username string
		err = rows.Scan(&id/*, &roleId, &username, &alias*/)
		PanicErr(err, &xgoto)
		Reid = id
		//fmt.Println("id =", id)
		//fmt.Println("roleId =", roleId)
		//fmt.Println("username =", username)
		//fmt.Println("alias =", alias)
	}
	return int64(Reid)
}

func sayHello2(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w,"id =", getRows(Id))
	Id += 1
}