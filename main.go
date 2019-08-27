package main

import (
	"fmt"
	"log"
)

func main() {
	var err error

	config := NewConfig()

	err = config.LoadFromEnv()

	if err != nil {
		panic(err)
	}

	log.Printf("Config: %v\n", config)

	loader := NewDBLoader("db.json")

	models, err := loader.Load()

	if err != nil {
		panic(err)
	}

	service := NewLogService(config.PageSize, models)

	controller := NewLogController(config.Port, service)

	fmt.Println("Starting controller at port", config.Port)

	controller.Start()
}
