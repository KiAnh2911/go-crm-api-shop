package initialize

import (
	"fmt"

	"github.com/KiAnh2911/go-crm-api-shop/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") //path to config
	viper.SetConfigName("local")     //tên file
	viper.SetConfigType("yaml")      // loại file

	// read configuration
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Failed to read config file: %w \n", err))
	}

	fmt.Println("Server Port:: ", viper.GetInt("server.port"))
	fmt.Println("Security :: ", viper.GetString("security.jwt.key"))

	// configure structure

	if err := viper.Unmarshal((&global.Config)); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

}
