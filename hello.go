package main

import "fmt"
import "time"

//"sync"专门用于处理并发和同步操作,sync.WaitGroup：用于等待多个 goroutines 执行完毕。

// 标准函数库fmt

// 并发执行 ：并发指的是多个任务同时进行，goroutine就是Go语言中的并发单元。可以通过go关键字来启动一个新的goroutine，它会和主程序并行运行。
// 你可以启动多个 goroutines，让它们同时工作，而不用一个接一个地执行。

// 如果没有写func的话直接go print那么会出现顺序不对的问题因为他们是同时进行的
func hello(i int) {
	fmt.Printf("Hello from goroutine %d\n", i)
}

func main() {
	for i := 0; i < 5; i++ {
		go hello(i)

	}
	time.Sleep(1 * time.Second)
}
