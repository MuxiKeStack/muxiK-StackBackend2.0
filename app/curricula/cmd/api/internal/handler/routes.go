// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: addCurriculaHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/curricula"),
	)
}
