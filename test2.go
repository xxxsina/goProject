package main

import (
	"fmt"
	"log"
	"syscall"
)

func ifelseFunc() {
	//n只在if块中作用，下面的n是没的
	if n, err := syscall.ComputerName(); err != nil {
		log.Printf("%s", err)
	} else {
		fmt.Printf(n)
	}
	//这里获取不到n
	//fmt.Printf(n)
}
func gotoFunc() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	if i > 10000 {
		return
	}
	goto Here //跳转到Here去，然后就是死循环了，这个时候需要if跳出
}
func forFunc() {
	sum := 0
	//一
	for index := 0; index < 10; index++ {
		sum += index
	}
	println("first =>", sum)

	//二
	sum2 := 1
	for ;sum2 < 10; {
		sum2 += sum2
	}
	println(sum2)
	//三
	sum3 := 1
	for sum3 < 10  {
		sum3 += sum3
	}
	println(sum3)
	//rang map
	for k,v:=range map[string]int{"a":1,"b":2,"c":3} {
		println(k, "=>", v)
	}
	for _, v:=range map[int]string{1:"a",2:"sb",3:"dfdf"}{
		println("key", "=>", v)
	}
}
func switchFunc() {
	i := 2
	switch i {
	case 1:
		println("i1 =", i)
	case 2:
		println("i2 =", i)
		fallthrough	//默认自带break，fallthrough为强制执行后面的代码
	default:
		println("default =", i)
	}
}
func main() {
	switchFunc()
	//forFunc()
	//gotoFunc()
	//ifelseFunc()
}
