package main

import "github.com/KiAnh2911/go-crm-api-shop/internal/initialize"

func main() {
	// r := routers.NewRouter()

	// InitMysql()
	// InitRedis()
	// InitKafka()

	// r.Run(":8000")  // DEFAULT : localhost = 8080
	initialize.Run()
}
