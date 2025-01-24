package main

import (
	"fmt"
	"sync"
)

// 定义一个函数，用于打印问候信息
func greet(name string, wg *sync.WaitGroup) {
	// 当函数执行完毕时，通知WaitGroup该Goroutine已完成
	defer wg.Done()
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// 创建一个WaitGroup实例，用于等待所有Goroutine完成
	var wg sync.WaitGroup

	// 定义一个包含名字的切片
	names := []string{"Alice", "Bob", "Charlie", "David"}

	// 为每个名字启动一个Goroutine
	for _, name := range names {
		// 增加WaitGroup的计数，表示有一个新的Goroutine要执行
		wg.Add(1)
		// 启动一个新的Goroutine来执行greet函数
		go greet(name, &wg)
	}

	// 等待所有Goroutine完成
	wg.Wait()

	fmt.Println("All greetings have been sent!")
}
