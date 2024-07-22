package initialize

import (
	"fmt"

	"github.com/KiAnh2911/go-crm-api-shop/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Load configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
