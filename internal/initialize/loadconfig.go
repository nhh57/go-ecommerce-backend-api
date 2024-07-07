package initialize

import (
	"fmt"

	"github.com/nhh57/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func InitLoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt(("server.port")))
	fmt.Println("Server Port::", viper.GetString(("security.jwt.key")))

	//configure structure

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
