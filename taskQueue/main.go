package main

import (
	"fmt"
	"sync"
	"time"
)

var GOTIME = "2006-01-02 15:04:05"

func main() {

	var taskInstall taskQueue
	taskInstall.New(3)
	for i := 0; i < 10; i++ {
		taskInstall.AddTask(func() {
			time.Sleep(5 * time.Second)
			fmt.Println(time.Now().Format(GOTIME))
		})
	}

	taskInstall.Run()

}

func (this *taskQueue) New(i int) *taskQueue {
	this.asyncNum = i
	this.TaskNum = 0
	this.TaskNumNow = 0
	this.waitList = make(map[int]func())
	this.runList = make(map[int]func())
	return this
}

type taskQueue struct {
	mu         sync.Mutex
	waitList   map[int]func()
	runList    map[int]func()
	asyncNum   int
	TaskNum    int
	TaskNumNow int
}

// 当函数退出时，将函数从安装队列删除
func (this *taskQueue) RunTask(taksNum int) {
	taskFunc := this.runList[taksNum]
	taskFunc()
	defer this.mu.Unlock()
	this.mu.Lock()
	delete(this.runList, taksNum)
}

func (this *taskQueue) AddTask(f func()) {
	//将任务添加进来

	this.waitList[this.TaskNum] = f
	this.TaskNum += 1

}

func (this *taskQueue) Run() {
	for {
		//退出条件
		if len(this.waitList) == 0 && len(this.runList) == 0 {
			return
		}

		//并发控制
		if len(this.runList) == this.asyncNum || len(this.waitList) == 0 {
			continue
		}

		this.mu.Lock()

		//从等待队列拿出并且删除一个task
		//将task加入安装队列并且启动协程执行该函数

		taskFunc := this.waitList[this.TaskNumNow]
		if taskFunc == nil {
			fmt.Println(this.TaskNumNow)
		}
		this.runList[this.TaskNumNow] = taskFunc
		delete(this.waitList, this.TaskNumNow)
		go this.RunTask(this.TaskNumNow)
		this.TaskNumNow += 1

		this.mu.Unlock()

	}

}
