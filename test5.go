package main

import (
	"fmt"
	"time"
)

//Go 方法
//一般的函数定义叫做函数，定义在结构体上面的函数叫做该结构体的方法。
type rect struct {
	width, height int
}

//对结构体实现一个可以改变参数值的area的方法
func (r *rect) area() int {
	return r.width * r.height
}
//对结构体实现一个不可以改变参数值的perim的方法
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
//方法的调用
type Person struct {
	ID int
	Name string
}
func (this Person) test(x int) {
	fmt.Println("ID :", this.ID, "\tName", this.Name)
	fmt.Println("args :", x)
}
//匿名函数 结构体的方法重写
type body struct {
	id int
	name string
	weight int
	height int
}
type Student struct {
	body	//身份属性
	score int	//成绩属性
}
func (this body) test1(x int) {
	fmt.Println("this body struct")
	fmt.Println("My name is", this.name)
}
func (this Student) test1(x int)  {
	fmt.Println("this Student struct")
	fmt.Println("My name is", this.name, "\tMy score is", this.score)
}
func main() {
	r := rect{width:10, height:30}
	//调用方法
	fmt.Println("area :", r.area())
	fmt.Println("perim :", r.perim())
	//一个拷贝
	rp := &r
	fmt.Println("area :", rp.area())
	fmt.Println("perim :", rp.perim())
	//另外一个种调用形式
	p := Person{ID: 1, Name: "LEE"}
	p.test(100)
	Person.test(p, 200)
	//还原成函数形式 (type).func(instance,args)
	var f1 func(int2 int) = p.test
	f1(120)
	var f2 func(Person, int) = Person.test
	f2(p, 122)
	/****嵌套与重写****/
	s := Student{body{id:1, name:"LEE", weight:120, height: 170}, 90}
	s.test1(99)
	//时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano() / 1000000)
}