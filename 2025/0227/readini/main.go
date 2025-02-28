package main

import (
	"fmt"
	"log"

	"github.com/vick/demo/2025/0227/readini/cfg"
	"gopkg.in/ini.v1"
)

func main() {
	cfg := new(cfg.ConfigSt)
	cfgIni, err := ini.Load("/Users/victor/code/go-helloworld/2025/0227/readini/config.ini")
	if err != nil {
		log.Println("Failed to load config file:", err)
		return
	}

	// 打印所有节和键值对
	for _, section := range cfgIni.Sections() {
		fmt.Printf("Section: %s\n", section.Name())
		for _, key := range section.Keys() {
			fmt.Printf("  %s = %s\n", key.Name(), key.Value())
		}
	}

	err = cfgIni.StrictMapTo(cfg)
	if err != nil {
		log.Println("Failed to map config:", err)
		return
	}
	log.Printf("Loaded Config: %+v\n", cfg)
	log.Printf("Loaded Config: %+v\n", cfg.StreamEngine.Enable)

}
