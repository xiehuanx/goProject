package main

import (
	"fmt"
	"runtime"
)

type staudent struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Class string `json:"class"`
}

func main() {

	SelectFunc()
}
func SelectFunc() {
	chananel := make(chan int, 5)
	chananel2 := make(chan string, 5)
	for i := 0; i < 5; i++ {
		chananel <- i
		chananel2 <- string(rune(i))
	}
	for {
		select {
		case v := <-chananel:
			fmt.Println("chananel-----", v)
		case v := <-chananel2:
			fmt.Println("chananel2-----", v)
		default:
			fmt.Println("结束了结束了")
			return
		}

	}
}

func WriteAndRead() {
	initChal := make(chan int, 1)
	exitChal := make(chan bool, 1)
	go writeData(initChal)
	go readData(initChal, exitChal)
	for {
		_, ok := <-exitChal
		if !ok {
			break
		}
	}
}

/**
写入数据
*/
func writeData(initChal chan int) {
	for i := 1; i <= 50; i++ {
		initChal <- i
		fmt.Println("writeData 写入当前数据位：", i)
	}
	fmt.Println("写入数据完成，关闭channel!")
	close(initChal)
}

func readData(initChal chan int, exitChal chan bool) {
	for {
		v, ok := <-initChal
		if !ok {
			fmt.Println("读取数据结束！")
			break
		}
		fmt.Println("readData 读取当前数据位：", v)
	}
	exitChal <- true
	close(exitChal)
}

func channel() {
	//获取cpu数量
	fmt.Println("121212")
	cpu := runtime.NumCPU()
	fmt.Println(cpu)

	var chal chan string
	chal = make(chan string, 3)
	fmt.Println(chal, &chal)

	// 向管道中加入数据

	chal <- "谢环"
	chal <- "测试"
	chal <- "侧模"

	fmt.Println(<-chal)

	var objectChal chan interface{}
	objectChal = make(chan interface{}, 4)
	//objectChal <- 10
	//objectChal <- "谢环"
	//objectChal <- 10.02

	xiehuan := staudent{
		12,
		"谢环",
		"13",
		"yiban",
	}
	//关闭chan

	objectChal <- xiehuan

	student := <-objectChal
	a := student.(staudent)
	fmt.Println(a)
	var test chan int64
	test = make(chan int64, 100)
	for i := 0; i < 100; i++ {
		test <- int64(i*2 + 1)
	}

	for i := range test {
		fmt.Println(i)
	}
}
