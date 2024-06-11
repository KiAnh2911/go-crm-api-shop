package main

import (
	"github.com/KiAnh2911/go-crm-api-shop/internal/routers"
)

func main() {
  r := routers.NewRouter()
  r.Run(":8000")  // DEFAULT : localhost = 8080
}

