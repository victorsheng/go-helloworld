package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "text/plain")
	// 写入响应内容
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// 注册处理函数
	http.HandleFunc("/", helloHandler)

	// 启动HTTP服务器，监听8080端口
	fmt.Println("Starting server at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
