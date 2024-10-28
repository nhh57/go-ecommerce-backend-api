package initialize

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"time"

	"github.com/nhh57/go-ecommerce-backend-api/global"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MYSQL Successfully")
	global.Mdbc = db
	// set Pool
	SetPool()
	//migrateTables()
}

func SetPoolC() {

	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Println("mysql error:: %s::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxIdleTime(time.Duration(m.ConnMaxLifetime))
}

func migrateTablesC() {
	err := global.Mdb.AutoMigrate(
		//&po.User{},
		//&po.Role{},

	)
	if err != nil {
		fmt.Println("Migrating tables error")
	}
}
