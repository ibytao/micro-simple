package db

import (
	"database/sql"
	"fmt"
	"micro-simple/basic"
	"sync"

	log "github.com/micro/go-micro/v2/logger"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func init() {
	basic.Register(initDB)
}

func initDB() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[initDB] db 已经初始化过")
		log.Info(err)
		return
	}

	initMysql()

	inited = true
}

func GetDB() *sql.DB {
	return mysqlDB
}
