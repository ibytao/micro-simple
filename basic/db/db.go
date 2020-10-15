package db

import (
	"database/sql"
	"fmt"
	"micro-simple/basic/config"
	"sync"

	log "github.com/micro/go-micro/v2/logger"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化")
		log.Error(err)
	}

	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}

	inited = true
}

func GetDB() *sql.DB {
	return mysqlDB
}
