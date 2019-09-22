package main

import "fmt"

//非阻塞通道
func chan1() {
	c1 := make(chan string)
	c2 := make(chan bool)
	select {
	case msg := <-c1:
		fmt.Print("this is :", msg)
	default:
		fmt.Println("this is default")
	}

	msg := "chan one"
	select {
	case c1 <- msg:
		fmt.Println("Push c1 :", msg)
	default:
		fmt.Println("this is in default")
	}

	select {
	case msg := <-c1:
		fmt.Println("c1 case", msg)
	case msg2 := <-c2:
		fmt.Println("c2 case", msg2)
	default:
		fmt.Println("this is double default")
	}
}

//带缓存的非阻塞通道
func chan2() {
	c3 := make(chan string, 1)
	c4 := make(chan bool)
	select {
	case msg := <-c3:
		fmt.Println("this is case msg :", msg)
	default:
		fmt.Println("this is one default")
	}
	msg1 := "hi"
	select {
	case c3 <- msg1:
		fmt.Println("this is case msg1:", msg1)
	default:
		fmt.Println("this is two default")
	}

	select {
	case msg := <-c3:
		fmt.Println("this is 2 case:", msg)
	case msg2 := <-c4:
		fmt.Println("this is 2 case 2:", msg2)
	default:
		fmt.Println("this is 22 default")
	}
}
func main() {
	chan2()
}
