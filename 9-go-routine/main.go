// https://learn.microsoft.com/zh-cn/training/modules/go-concurrency/3-buffered-channels
// 不是通过共享内存通信；而是通过通信共享内存。

package main

import (
	"fmt"
	"net/http"
	"time"
)

// chan 默认是传指针
// chan<- type // it's a channel to only send data
// <-chan type // it's a channel to only receive data
func checkApi(api string, ch chan<- string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!", api)
	} else {
		ch <- fmt.Sprintf("SUCCESS: %s is up and running!", api)
	}
}

func process(ch chan<- string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}

	// 有缓冲 channel 在不阻止程序的情况下发送和接收数据，因为有缓冲 channel 的行为类似于队列。
	// 创建 channel 时，可以限制此队列的大小。
	// 无缓冲 channel 同步通信。它们保证每次发送数据时，程序都会被阻止，直到有人从 channel 中读取数据。
	// 相反，有缓冲 channel 将发送和接收操作解耦。它们不会阻止程序，但你必须小心使用，因为可能最终会导致死锁。
	var ch chan string = make(chan string, 5)

	for _, api := range apis {
		// channel 与 goroutine 有着紧密的联系。
		// 如果没有另一个 goroutine 从 channel 接收数据，则整个程序可能会永久处于被阻止状态。
		// 我们建议在使用 channel 时始终使用 goroutine
		go checkApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

	// ------------------------------------------
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		process(ch1)
	}()

	go func() {
		replicate(ch2)
	}()

	for i := 0; i < 2; i++ {
		// select 语句的工作方式类似于 switch 语句，但它适用于 channel。
		// 它会阻止程序的执行，直到它收到要处理的事件。
		// 如果它收到多个事件，则会随机选择一个。
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}

	// ------------------------------------------
	printFib()
}
