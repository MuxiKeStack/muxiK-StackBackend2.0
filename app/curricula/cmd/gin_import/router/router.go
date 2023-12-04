package router

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/gin_import/handler"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	e := gin.Default()
	e.POST("/api/v1/curricula/import", handler.ImportCurricula)
	return e
}
