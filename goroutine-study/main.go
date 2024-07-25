package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	age  int
}

func main() {
	//当没有给chan设置缓冲时，有数据进chan时，必须有其他任务读取chan中的数据。否则会deadline
	ch1 := make(chan string)
	//从goroutine中获取值的两种方式，1chan，2 *struct{}
	go test1(ch1)
	//select 语句可以接收多个ch，谁的ch有值输出执行那个的case，当两个ch的数据都存在时，会运行一个随机算法，随机选择h一个case执行
	go test2(ch1)

	human := person{}
	ch2 := make(chan string)
	go test3(&human, ch2)

	for {
		k, ok := <-ch2
		if ok {
			fmt.Println("main ch read done", k)
			fmt.Println(human)
			break
		}

	}

	ch3 := make(chan string)
	//在ch没有写入时，关闭ch之后，读取ch的值为空，不关闭ch，读取ch会死锁
	close(ch3)
	ch1Value := <-ch3
	fmt.Println("ch1Value:", ch1Value)
	time.Sleep(time.Second * 3)
	//有缓冲chan会写入缓冲的数量的值之后在退出，无缓冲chan没有人读的情况下无法写入数据
	GOTIME := "2006-01-02 15:04:05"
	ch4 := make(chan string, 3)
	go test4(ch4, GOTIME)
	time.Sleep(time.Second * 10)
	fmt.Println(time.Now().Format(GOTIME), "over")

}

func test1(ch chan string) {

	time.Sleep(time.Second * 1)
	fmt.Println("test1 run")
	ch <- "chan"

}

func test2(ch chan string) {
	var testBool bool
	testBool = false
loop:
	for {
		select {
		case v, ok := <-ch:
			fmt.Println("select is run")
			if ok {
				println("test 3 print ", v)
				testBool = true
				break loop
			}
		case <-time.After(time.Second * 2):
			fmt.Println("time out")
			break loop
		}
		fmt.Println("test3 的select 执行退出")
	}

	fmt.Println(testBool)
}

func test3(human *person, ch chan string) {
	defer func() {
		close(ch)
	}()
	human.age = 18
	human.name = "wty"
	ch <- "test3"
}
func test4(ch chan string, gotime string) {
	defer func() {
		close(ch)
	}()

	for i := 1; i <= 5; i++ {
		v := time.Now().Format(gotime)
		ch <- v
		fmt.Println("i is ", i, "value is ", v)
		time.Sleep(time.Second)
	}
}

func test5(ch chan string) {

}
