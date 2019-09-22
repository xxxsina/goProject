package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	v := os.Getenv("OS")
	if v == "" {
		//fmt.Println("v is null")
		//os.Setenv("FOO", "LEEXGO")
		//v = os.Getenv("FOO")
	}
	fmt.Println(v)
	//列出所有的环境变量
	for _, e := range os.Environ() {
		par := strings.Split(e, "=")
		fmt.Println("Name :", par[0], ";\tPath :", par[1])
	}
}
