package main

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func main () {
	println("开始创建文件")
	f, err := os.Create(".\\file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("开始写文件")
	fmt.Fprintln(f, "hello world")
	//sha1散列值
	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
	println("关闭文件")
	f.Close()
}
