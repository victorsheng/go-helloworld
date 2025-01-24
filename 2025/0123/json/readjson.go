package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

func main() {
	// JSON数据
	jsonData := `{
        "name": "John Doe",
        "age": 30,
        "email": "john.doe@example.com",
        "address": {
            "street": "123 Main St",
            "city": "Anytown",
            "state": "CA",
            "zip": "12345"
        }
    }`

	// 创建一个Person类型的变量
	var person Person

	// 解析JSON数据到person变量中
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 打印解析后的数据
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Email: %s\n", person.Email)
	fmt.Printf("Address: %s, %s, %s %s\n", person.Address.Street, person.Address.City, person.Address.State, person.Address.Zip)
}
