package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//闭包：就是个函数

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//不能如此调用
	// fmt.Println(nextInt)

	newInts := intSeq()
	fmt.Println(newInts())
}
