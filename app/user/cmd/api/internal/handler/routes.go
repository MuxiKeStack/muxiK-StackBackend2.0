// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthorizeHandler},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: updateInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/share",
					Handler: sharingPlanHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/info",
					Handler: getMyInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/info/:sid",
					Handler: getOtherInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/user"),
	)
}