package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//queue1 := make(chan string, 2)
	//queue1 <- "a"
	//queue1 <- "b"
	//for elem := range queue1 {
	//	fmt.Println(elem)
	//}

	//进程触发
	// 我们从一个简单的命令开始，这个命令不需要任何参数
	// 或者输入，仅仅向stdout输出一些信息。`exec.Command`
	// 函数创建了一个代表外部进程的对象
	dateCmd := exec.Command("date")
	// `Output`是另一个运行命令时用来处理信息的函数，这个
	// 函数等待命令结束，然后收集命令输出。如果没有错误发
	// 生的话，`dateOut`将保存date的信息
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Print(dateOut)
}
