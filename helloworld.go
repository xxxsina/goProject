package main

import (
	"errors"
	"fmt"
)

//这种方式一般声明全局变量
var vname1 = 10

//这里声明一个字符串类型的变量要使用双引号
var vnameString = "abc"

/**
定义常量
Go 常量和一般程序语言不同的是，
可以指定相当多的小数位数(例如200位)，
若指定給float32自动缩短为32bit，
指定给float64自动缩短为64bit
*/
const PI float64 = 3.1415926
const Prefix string = "LEE learning" //字符串常量
//布尔值的类型
var isAction bool = true
var enabled, disabled = true, false // 忽略类型的声明
//分组声明变量
const (
	xi = 100
	pii = 3.14
	pre string = "GO_"
)
var (
	ii int
	ss string
	piii float32 = 3.12
)
/**
Go里面有一个关键字iota，
这个关键字用来声明enum的时候采用，
它默认开始值是0，const中每增加一行加1
 */
const (
	x = iota //x == 0
	y = iota //y == 1
	z = iota //z == 2
	w // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	o = "b"
	g = iota
	e,f = iota,iota
	q = iota
)
const v = iota //每遇到一个const关键字，iota就会重置，此时v == 0
const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)
/**
大写字母开头:公有;小写字母开头:私有
函数也一样
 */
var Public = "public var"
var private = "private var"
func Pub(){}
func pri(){}
//数组 array 的定义方式 : var arr [n]type
var arr [10]int


func test() {
	//:= 这种方式只能用于函数内部
	vname2, vname3 := 11, 12
	//_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃
	_, vname4 := 43, 34
	//这里声明一个变量如果不使用就要报错
	var i int
	//boolean赋值操作
	isAction = false
	/**
	Go里面也有直接定义好位数的类型：
	rune, int8, int16, int32, int64和
	byte, uint8, uint16, uint32, uint64。
	其中rune是int32的别称，byte是uint8的别称
	注意：这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。
	如下
	*/
	//var a int8 = 1
	//var b int16 = 2
	//c := a + b	//报错
	//fmt.Println(c)
	//浮点数
	var c complex64 = 5 + 5i
	fmt.Printf("Value is: %v", c)
	fmt.Println(vname2)
	fmt.Println(vname3)
	fmt.Println(vname4)
	//fmt.Print(vnameString)
	fmt.Println(i)
	fmt.Println(PI)
	fmt.Println(isAction)
}
func testSting() {
	var s string = "hello"	//声明字符串要用双引号
	c := []byte(s) //如果要替换第一个字母h为c，应该先将字符串转化成[]byte类型
	c[0] = 'c'	//这里要用单引号，要不然报错
	s2 := string(c)
	fmt.Println(s2)
	//使用+来连接字符串
	m := " world"
	a := s + m
	fmt.Println(a)
	//修改字符串也可写为
	s3 := "A" + s[1:]	//字符串虽不能更改，但可进行切片操作
	fmt.Println(s3)
	//如果要声明一个多行的字符串怎么办？可以通过`来声明
	mm := `hello 
	world`
	fmt.Println(mm)
	//Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}
func testConst() {
	fmt.Println(i)
	fmt.Println(z)
	fmt.Println(v)
	fmt.Println(j)
	fmt.Println(w,o)
	fmt.Println(g)
	fmt.Println(e,f)
	fmt.Println(q)
}
func testArr() {
	arr[0] = 10
	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组
	b := []int{2,3,4,5,6,7}	//slice
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	//多维数组
	arrArr := [2][3]int{[3]int{1,2}, [3]int{3,4,5}}
	fmt.Println(arr)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(arrArr)
	fmt.Printf(" ------------------ \n")
	/**
	这里数组与切片
	可以这么说指定长度就是数组，未指定为切片
	 */
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}//声明一个数组
	//声明2个切片slice
	slice1 := []byte{}
	slice2 := []byte{}
	//slice3 := []byte{}
	slice1 = array[1:2]
	slice2 = array[3:7]
	fmt.Printf("%c\t%d\n", array[0], array[0])
	fmt.Println(slice1[:])
	fmt.Println(slice2[:])
	fmt.Println(cap(slice1))
	fmt.Println(cap(slice2))
	fmt.Println(append(slice2, 'a','c','d','d','c','d','c'))
	fmt.Println(cap(slice2))
	slice1 = array[3:3:10]
	fmt.Println(cap(slice1))
}
func testMap() {
	//字典
	//声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int
	// 另一种map的声明方式
	numbers = make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	numbers["four"] = 4
	delete(numbers, "four")
	fmt.Println(numbers)
	fmt.Println(numbers["two"])
	fmt.Println(len(numbers))
	// 初始化一个字典
	rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	m := rating
	m["Go"] = 1.18
	fmt.Println(m)
	fmt.Println(rating)
}
func main() {
	//fmt.Println("hello world！")
	testMap()
	//testArr()
	//testConst()
	//testSting()
	//test()
	//http.HandleFunc('/', test)
	//http.ListenAndServe(':9090', nil)
}
