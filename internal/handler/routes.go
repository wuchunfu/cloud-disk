// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	file "cloud-disk/internal/handler/file"
	user "cloud-disk/internal/handler/user"
	"cloud-disk/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/registerCode",
				Handler: RegisterCodeHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/detail",
					Handler: user.UserDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/repository/link",
					Handler: user.UserRepositoryLinkHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/file/list",
					Handler: user.UserFileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/fileName/edit",
					Handler: user.UserFileNameEditHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/dir/create",
					Handler: user.UserDirCreateHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: file.UploadFileHandler(serverCtx),
				},
			}...,
		),
	)
}
