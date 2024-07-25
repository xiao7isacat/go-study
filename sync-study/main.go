package main

import (
	"fmt"
	"sync"
	"time"
)

/*
sync包，包含有
Mutex			互斥锁
RWMutex			读写锁
WaitGroup		并发等待组
Map				并发安全字典
Once			单例模式

Cond			同步等待条件
Pool			临时对象池
*/

func main() {
	GOTIME := "2006-01-02 15:04:05"
	//互斥锁
	MeutxStudy(GOTIME)
	//读写锁
	RWMeutxStudy(GOTIME)
	//等待组
	WaitGroupStudy(GOTIME)
	//并发安全字典
	MapStudy()
	fmt.Println(time.Now().Format(GOTIME))
}

func MeutxStudy(GOTIME string) {
	//保证同一时间内只有一个goroutine持有锁
	var lock sync.Mutex

	go func() {
		fmt.Println("start func1", time.Now().Format(GOTIME))
		defer lock.Unlock()
		lock.Lock()
		fmt.Println("func 1 get lock at", time.Now().Format(GOTIME))
		time.Sleep(time.Second)
		fmt.Println("func 1 release lock at", time.Now().Format(GOTIME))
	}()
	time.Sleep(time.Second / 100)
	go func() {
		fmt.Println("start func2", time.Now().Format(GOTIME))
		defer lock.Unlock()
		lock.Lock()
		fmt.Println("func 2 get lock at", time.Now().Format(GOTIME))
		time.Sleep(time.Second)
		fmt.Println("func 2 release lock at", time.Now().Format(GOTIME))
	}()

	time.Sleep(time.Second * 2)
}

func RWMeutxStudy(GOTIEM string) {
	/*
		1.同一个时间段只有一个goroutine可以获取写锁
		2.同一个时间段可以有多个goroutine获取读锁,读锁之间相互独立
		3.同一个时间只能存在读锁或者写锁
	*/
	var lock sync.RWMutex
	//read lock
	fmt.Println("start read")
	for i := 1; i <= 10; i++ {
		go func(i int) {
			defer func() {
				lock.RUnlock()
				fmt.Println("    read func", i, "release read lock", time.Now().Format(GOTIEM))

			}()
			fmt.Println("read func start", i, time.Now().Format(GOTIEM))
			lock.RLock()
			fmt.Println("  read func", i, "get read lock", time.Now().Format(GOTIEM))
			//time.Sleep(time.Second / 1000)
		}(i)
	}

	time.Sleep(time.Second / 10)

	//write lock
	fmt.Println("start write")
	for i := 1; i <= 5; i++ {
		go func(i int) {
			defer func() {
				lock.Unlock()
			}()
			fmt.Println("write func start", i, time.Now().Format(GOTIEM))
			lock.Lock()
			fmt.Println("  write func", i, "get write lock", time.Now().Format(GOTIEM))
			time.Sleep(time.Second)
			fmt.Println("    write func", i, "release write lock", time.Now().Format(GOTIEM))
		}(i)
	}
	time.Sleep(time.Second * 8)
}

func WaitGroupStudy(GOTIME string) {
	/*
		使用waitgroup的goroutine会等待设置好数量的goroutine都执行结束后，才会继续往下执行
		在goroutine调用waitgroup之前要确保waitgroup中等待的数量大于1
		保证waitgroup.done() 的执行次数和waitGroup.add()相同，过少会死锁，过多会panic
	*/

	var wg sync.WaitGroup
	//wg.Add(5)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("func start", i, time.Now().Format(GOTIME))
			fmt.Println("func", i, "start add wg", time.Now().Format(GOTIME))
			time.Sleep(time.Second)
			wg.Done()
			fmt.Println("func", i, "add wg done", time.Now().Format(GOTIME))
		}(i)
	}
	wg.Wait()
	fmt.Println("main func", time.Now().Format(GOTIME))

}

type MapTest struct {
	name string
	age  int
}

func (this *MapTest) Test() {
	this.name = "wty"
	this.age = 24
}

func MapStudy() {
	/*
		go中原生的map并不是并发安全的，多个goroutine同时操作时，原生的mao会报错
		常用方法：
			1. Load 根据k获取v
			2. Store 存储k，v
			3. LoadOrStore 存在则返回，不存在则保存
			4. Delete 删除
			5. Range 遍历
	*/
	var (
		dict sync.Map
		wg   sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			dict.LoadOrStore(i, i)
		}(i)
	}
	wg.Wait()

	dict.Store(1, "aaa")
	v, ok := dict.Load(1)
	fmt.Println("v is ", v, "ok is ", ok)
	dict.Delete(10)
	v, ok = dict.Load(10)
	fmt.Println("v is ", v, "ok is ", ok)
	v, ok = dict.LoadAndDelete(9)
	fmt.Println("v is ", v, "ok is ", ok)
	v, ok = dict.Load(9)
	fmt.Println("v is ", v, "ok is ", ok)
	dict.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	var mapTest MapTest
	dict.Store(11, mapTest)
	v, ok = dict.Load(11)
	switch value := v.(type) {
	case MapTest:
		value.Test()
		fmt.Println(value)

	}
}
