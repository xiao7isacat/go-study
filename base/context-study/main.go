package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	test()
}

func test() {
	flag := make(chan int)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		wg.Add(1)
		t := time.Tick(time.Second)
		for range t {
			select {
			case v, ok := <-flag:
				if ok {
					fmt.Println("接收请求", v)
				}
			case <-ctx.Done():
				defer wg.Done()
				fmt.Println("goroutine 退出")
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		flag <- i
	}
	close(flag)
	cancel()
	wg.Wait()
	fmt.Println("主程序执行完成")
}
