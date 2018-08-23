package db

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbUser := beego.AppConfig.String("db.user")
	dbPass := beego.AppConfig.String("db.pass")
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbName := beego.AppConfig.String("db.database")
	dbDriver := beego.AppConfig.String("db.driver")

	driverMap := make(map[string]orm.DriverType)
	driverMap["mysql"] = orm.DRMySQL
	driverMap["sqlite"] = orm.DRSqlite
	driverMap["oracle"] = orm.DROracle
	driverMap["postgres"] = orm.DRPostgres
	driverMap["tidb"] = orm.DRTiDB
	orm.Debug = true

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", dbUser, dbPass, dbHost, dbPort, dbName)

	beego.Debug(dbSource)

	orm.RegisterDriver(dbDriver, driverMap[dbDriver])
	err := orm.RegisterDataBase("default", dbDriver, dbSource)
	if err != nil {
		fmt.Print(err)
	}

	orm.RegisterModel(new(Message))
}
