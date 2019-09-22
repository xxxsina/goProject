package main

import (
	. "fmt" //这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")
	o "os" //别名o
)

//有返回值函数 一个返回值
func funcOne(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//多个返回值函数
func funcTwo(a, b int) (int, int) {
	return a + b, a * b
}
func funcThree(a, b int) (add int, Multi int) {
	add = a + b
	Multi = a * b
	return
}
//变参
func funcFour(argAll ...string) {
	for k, v := range argAll {
		if v == "3" {
			println(k, "=>", v)
		}
	}
}
//指针传递
func funcFive(x *int) int {
	*x++
	return *x
}
//defer 延迟
func funcSix() bool {
	i := 0
	for i < 5 {
		defer println(i)
		i++
	}
	//var _, err := file.Open("./test2.go")
	//defer file.Close()
	//if err != nil {
	//	println(err)
	//}
	return true
}
//Panic
func funcPanic() {
	//这里recover恢复了
	defer func() {
		if x := recover(); x != nil {
			println(x)
			return
		}
	}()
	//这里就panic了
	user := o.Getenv("USER")
	if user == "" {
		panic("no value for $USER")
	}
}
func init()  {
	Println("init和main函数为保留函数，任何一个package里面都会执行这两个函数，init可选，但是main是必须的")
}
func main() {
	//param := funcOne(12,3)
	//param1, param2 := funcTwo(12,3)
	//param1, param2 := funcThree(12,3)
	//println(param1, param2)
	//funcFour("1", "2","3","4")
	//x := 3
	//funcFive(&x)
	//println(x)
	//funcSix()
	funcPanic()
}