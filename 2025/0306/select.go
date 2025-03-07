package main

import (
	"fmt"
	"time"
)

// https://www.runoob.com/go/go-select-statement.html
// select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 100; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", i, msg1)
		case msg2 := <-c2:
			fmt.Println("received", i, msg2)
		default:
			fmt.Println("no message received", i)
			time.Sleep(1 * time.Second)
		}
	}
}
