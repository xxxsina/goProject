package main

import (
	"fmt"
	"time"
)

func worker(done chan int, i int) {
	fmt.Println("working =>")
	time.Sleep(time.Second * 1)
	fmt.Println(i)
	//写入通道
	done <- i
	//close(done)
}

//select关键字可以让你同时等待多个通道操作
func selecter() {
	//声明2个通道
	c1 := make(chan string)
	c2 := make(chan string)
	// 为了模拟并行协程的阻塞操作，我们让每个通道在一段时间后再写入一个值
	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1 :", msg1)
		case msg2 := <-c2:
			fmt.Println("msg2 ", msg2)
		}
	}
	//close(c1)
	//close(c2)
}
func main() {
	//声明一个通道, 缓冲区大小为1, 设置了就为带缓冲区的通道
	//如果设置为2  可以同时发送两个数据
	/**
	设置例子：
		message <- "one chan"
		message <- "two chan"
	获取例子：
		<-message
		<-message
	*/
	//done := make(chan int, 1)
	//使用协程
	//go worker(done, 2)
	//for i := 1; i <= 10; i++ {
	//	go worker(done, i)
	//	// 一直阻塞，直到从worker所在协程获得一个worker执行完成的数据
	//	//这里通过<-阻塞方式达到同步效果
	//	<-done
	//}
	//fmt.Println("End")

	//select方法
	//selecter()

	//遍历通道
	queue := make(chan string, 2)
	queue <- "first"
	queue <- "second"
	//fmt.Println(len(queue))	//通道大小
	for elem := range queue {
		fmt.Println(elem)
	}

	fmt.Println("ok")
}
