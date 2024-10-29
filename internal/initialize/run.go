package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/nhh57/go-ecommerce-backend-api/global"
)

func Run() *gin.Engine {
	InitLoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Port)
	InitLogger()
	// global.Logger.Info("Config Log oke:: ", zap.String("ok", "success"))
	InitMysql()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()
	InitKafka()

	r := InitRouter()
	return r
	//r.Run(":8002")
}
