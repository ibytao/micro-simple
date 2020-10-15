package basic

import (
	"micro-simple/basic/config"
	"micro-simple/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
