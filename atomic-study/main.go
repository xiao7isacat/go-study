package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	/*
		简介：
			1. 当想要对某个变量进行并发安全修改的时候，除了官方提供的mutex之外，可以使用sync/atomic 包的原子操作
			   它能保证对变量的读取或者修改期间不被其他的协程所影响
			2. go的原子操作都是非嵌入式的，直接由底层cpu硬件支持，也就是在硬件层次去实现的，性能比较好，并不像mutex那样记录很多状态
			   mutex并不只是对变量进行并发控制，而是可以对代码段进行并发控制。两者侧重点不同

		原理：
			原子操作是不能中断的，在针对某个变量进行原子操作时，cpu不会进行其他针对该值的操作。具体的原子操作在cpu上实现是不同的
			在inter的cpu上，主要使用的是总线锁的方式。当一个cpu需要操作一个内存块时，向总线发送一个LOCK的信号，所有的cpu接收到该信号之后，就不会对该内存进行操作，直到使用该内存的cpu操作完成，发送UNLOCK信号
			AMD的cpu上使用的是MESI一致性协议来实现的

		mutex和atomic的区别：
			1. 加锁代价比较高，耗时多，需要切换上下文
			2. 原子操作只能针对基础数据，不支持结构体，自定义数据类型
			3. 原子操作在用户态即可完成，性能比较高
			4. 针对特定需求的原子操作步骤加单，无需加锁解锁的步骤

		为什么atomic比mutex块
			1. 原子操作依赖于cpu指令，而不是依赖于外部锁，使用互斥锁时，每次获取锁，goroutine都会短暂暂停或中断，这种阻塞占使用锁花费的很大一部分时间
			2. 原子操作时能保证执行期间是连续且不会中断的，临界区（mutex上下文）只能保证访问共享数据是按照顺序访问的，并不能保证访问期间不会被上下文切换

		CSA
			CAS是cpu的硬件同步的原语。 Compare And Swap
			go中CAS操作，是借用cpu通用的原子性指令实现的，CAS操作修改共享变量的时候不需要对共享变量加锁，而是同过类似乐观锁进行检查，本质还是不断占用cpu来换取加锁的开销
			原子操作中的CAS，在atomic包中，这类原子操作以CompareAndSwap为前缀
			优势：
				可以在不形成临界区和创建互斥量的情况下完成并发安全的值替换操作，减少同步对程序性能的损耗
			劣势：
				在操作值被频繁更新的情况下，CAS操作并不是很容易成功的，因为需要对old的值进行匹配，只有匹配成功了才会进行下一步操作
			当前atomic包有以下几个原子操作： Add CompareAndSwap Load Store Swap
	*/

	//增或减
	//AddStudy()

	//比较并且替换
	//CompareAndSwapStudy()

	//读
	//LoadStudy()

	//存储
	//StoreStudy()

	//交换
	SwapStduy()

	time.Sleep(time.Second * 2)
}

func AddStudy() {
	/*
		1. 增或者减
		2. 原子操作的数据类型只能是数值行，int32 int64 uint32 uint64 uintptr
		3. 原子增减哈书的第一个值为原值，第二个值为增加或者减少多少
	*/
	var num int64
	fmt.Println(num)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(atomic.AddInt64(&num, 2))
		}()
	}
}

func CompareAndSwapStudy() {
	var num int64
	num = 2
	fmt.Println(num)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(atomic.CompareAndSwapInt64(&num, 2, 4))
		}()
	}
}

func LoadStudy() {
	var num int64
	var ch = make(chan int64)
	fmt.Println(num)
	for i := 0; i < 5; i++ {
		go func() {
			atomic.AddInt64(&num, 1)
		}()
	}
	for i := 0; i < 5; i++ {
		go func() {
			ch <- atomic.LoadInt64(&num)
		}()
	}
	for i := 0; i < 5; i++ {
		go func() {
			atomic.AddInt64(&num, 1)
		}()
	}
	for i := 0; i < 5; i++ {
		go func() {
			ch <- atomic.LoadInt64(&num)
		}()
	}

	for i := 0; i < 10; i++ {
		test := <-ch
		fmt.Println(test)
		fmt.Println(num)
	}
}

func StoreStudy() {
	var num, i int64
	fmt.Println(num)

	for i = 0; i < 10; i++ {
		go func(i int64) {
			atomic.StoreInt64(&num, i)
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(atomic.LoadInt64(&num))
		}()
	}
}

func SwapStduy() {

	var num, i int64
	fmt.Println(num)

	for i = 0; i < 10; i++ {
		go func(i int64) {
			fmt.Println("第", i, "次写入打印", atomic.SwapInt64(&num, i))
		}(i)
	}
	for i = 0; i < 10; i++ {
		go func(i int64) {
			fmt.Println("第", i, "次读打印", atomic.LoadInt64(&num))
		}(i)
	}
}
