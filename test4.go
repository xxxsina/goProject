package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//mutex 互斥
func main() {
	//声明一个字典
	var state = make(map[int]int)
	//声明一个mutex
	var mutex = &sync.Mutex{}
	//声明一个计算器
	var ops int64 = 0
	//启动100个协程 不停的读取
	for i := 0; i < 100; i++ {
		//并发方式
		go func() {
			//统计总数
			total := 0
			//这里不停的循环操作
			for {
				//0-5之间取随机数
				key := rand.Intn(6)
				//上锁
				mutex.Lock()
				//开始操作
				total += state[key]	//这里这个没得默认值是0
				//解锁
				mutex.Unlock()
				//计数器操作一次增加1
				atomic.AddInt64(&ops, 1)
				//手动地释放资源控制权
				runtime.Gosched()
			}
		}()
	}
	//启动10个协程不停的写入
	for j := 0; j < 10; j++ {
		go func() {
			for {
				key := rand.Intn(6)
				val := rand.Intn(100)	//取100以内的随机数
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	//主协程Sleep，让那10个协程能够运行一段时间，睡个1秒钟
	time.Sleep(time.Second)
	//输出总操作次数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops =>", opsFinal)
	//最后锁定并输出状态
	mutex.Lock()
	fmt.Println("state=>", state)
	mutex.Unlock()
}