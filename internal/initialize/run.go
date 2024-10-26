package initialize

import (
	"fmt"

	"github.com/nhh57/go-ecommerce-backend-api/global"
)

func Run() {
	InitLoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Port)
	InitLogger()
	// global.Logger.Info("Config Log oke:: ", zap.String("ok", "success"))
	InitMysql()
	InitMysqlC()
	InitRedis()
	InitKafka()

	r := InitRouter()
	r.Run(":8002")
}
