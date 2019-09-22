package main

import "fmt"
//闭包，这发返回值是函数，这个函数又返回int型数据
func inSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
//闭包，这个构造了一个x+y的函数
func closure(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}
//闭包，一个切片的闭包组
func sliceFunc() {
	//声明一个切片的闭包
	var fs []func() int
	//向切片入值
	for i := 0; i < 3; i++ {
		j := i
		fs = append(fs, func() int {
			return j
		})
	}
	//遍历切片
	for _, fx := range fs {
		fmt.Printf("%p = %v\n", fx, fx())
	}
}
//闭包，默认
func adder () func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	//调用闭包函数返回一个函数
	val := inSeq()
	//这里调用方式是val()
	fmt.Println("i =", val())	//i = 1
	fmt.Println("i =", val())	//i = 2
	fmt.Println("i =", val())	//i = 3
	fmt.Println("i =", val())	//i = 4
	//闭包测试二
	add10 := closure(10)
	fmt.Println("value =", add10(10))	//value = 20
	fmt.Println("value =", add10(20))	//value = 30
	add20 := closure(20)
	fmt.Println("value =", add20(5))	//value = 25
	//闭包测试三
	sliceFunc()
	//闭包测试四
	rsult := adder()
	for i := 0; i < 20; i++ {
		fmt.Println(rsult(i))
	}
}