package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Users []map[string]interface{} `yaml:"users"` // 数组中的元素是 map
}

func main() {
	// 读取 YAML 文件
	data, err := os.ReadFile("/Users/victor/code/go-helloworld/2025/0228/config.yaml")
	if err != nil {
		panic(fmt.Errorf("读取文件失败: %v", err))
	}

	// 解析 YAML
	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		panic(fmt.Errorf("解析 YAML 失败: %v", err))
	}

	// 遍历数组中的 map
	for _, user := range config.Users {
		fmt.Println("\n用户信息:")
		for key, value := range user {
			// 处理嵌套数据（例如 permissions 数组或 settings map）
			switch key {
			case "name", "role":
				fmt.Printf("- %s: %v\n", key, value)
			case "permissions":
				if permissions, ok := value.([]interface{}); ok {
					fmt.Printf("- 权限: %v\n", permissions)
				}
			case "settings":
				if settings, ok := value.(map[string]interface{}); ok {
					fmt.Println("- 设置:")
					for k, v := range settings {
						fmt.Printf("  - %s: %v\n", k, v)
					}
				}
			}
		}
	}
}
