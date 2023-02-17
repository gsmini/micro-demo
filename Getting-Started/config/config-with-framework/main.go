package main

import (
	"fmt"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/config"
)

func main() {
	// setup the service
	srv := service.New(service.Name("confgi-with-framework"))
	srv.Init()

	// read config value
	val, _ := config.Get("mysql")
	fmt.Printf("Value of mysql: %s", val)
}
