package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	GTPL_PATH = "./tpl"
)

type Myform struct {
	username string
	password string
}

//var Vnum int = 1

//func sayHello(w http.ResponseWriter, r *http.Request) {
//	err := r.ParseForm()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("Form data:", r.Form)
//	fmt.Println("PostForm data:", r.PostForm)
//	fmt.Println("Url.Path :", r.URL.Path)
//	fmt.Println("scheme:", r.URL.Scheme)
//	fmt.Println("URL", r.URL)
//	fmt.Println("r :", r)
//	for k, v := range r.Form {
//		fmt.Println("k :", k, " v :", v)
//	}
//	fmt.Fprintf(w, "Hello astx\r\n")
//	fmt.Fprint(w, Vnum)
//	Vnum += 1
//}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		path := GTPL_PATH + "/login.gtpl"
		if exits := isExists(path); !exits {
			fmt.Fprintf(w, "login.gtpl not isExists \r\n")
			fmt.Fprintf(w, path)
			return
		}
		t, _ := template.ParseFiles(path)
		//log.Println(t.Execute(w, nil))
		//token
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t.Execute(w, token)
	} else {
		//方法一
		//这句很重要，要不然数据出不来
		//r.ParseForm()  //解析url传递的参数，对于POST则解析响应包的主体(request body)
		////fmt.Fprint(w, r.ParseForm())
		//fmt.Println("username:", r.Form["username"][0])
		//fmt.Println("password:", r.Form["password"])
		//方法二
		data := Myform{
			username: r.FormValue("username"),
			password: r.FormValue("password")}
		fmt.Println("username:", data.username)
		fmt.Println("password:", data.password)
	}
}

func main() {
	//http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
		return
	}
}
