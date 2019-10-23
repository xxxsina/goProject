package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	tmpPath = "./src/lee/tpl/"
)

func isExits(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func web(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method :", r.Method)
	if r.Method == "GET" {
		//如果请求方法为get显示login.html,并相应给前端
		path := tmpPath + "websocket.gtpl"
		if exits := isExits(path); !exits {
			fmt.Fprintf(w, "websocket.gtpl not isExits \r\n")
			fmt.Fprintf(w, path)
			return
		}
		t, _ := template.ParseFiles(path)
		t.Execute(w, nil)
	} else {
		fmt.Println(r.PostFormValue("message"))
	}
}

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		//websocket接受信息
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("receive failed:", err)
			break
		}
		fmt.Println("reveived from client :" + reply)
		msg := "received :" + reply
		fmt.Println("send to client :" + msg)
		//这里是发送消息
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("send failed :", err)
			break
		}
	}
}

func main() {
	//接受websocket的路由地址
	http.Handle("/websocket", websocket.Handler(Echo))
	//html页面
	http.HandleFunc("/web", web)

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
