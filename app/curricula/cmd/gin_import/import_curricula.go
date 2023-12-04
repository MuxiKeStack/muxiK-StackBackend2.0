package main

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/gin_import/dao/mysql"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/gin_import/router"
)

func main() {
	e := router.RouterInit()
	mysql.InitMysql()
	err := e.Run(":8088")
	if err != nil {
		panic(err)
		return
	}
}
