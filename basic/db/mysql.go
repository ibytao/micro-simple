package db

import (
	"database/sql"
	"micro-simple/basic/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/micro/go-micro/v2/logger"
)

func initMysql() {
	var err error

	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	mysqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.GetMysqlConfig().GetConnMaxLifetime()))

	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
