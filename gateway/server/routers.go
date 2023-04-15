package server

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	boxv1 "github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/box/v1"
	dashv1 "github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/dash/v1"
)

func NewDashApiV1RouterGroup(server *server.Hertz) *dashv1.DashApiV1RouterGroup {
	return &dashv1.DashApiV1RouterGroup{RouterGroup: server.Group("/api/dash/v1")}
}

func NewBoxApiV1RouterGroup(server *server.Hertz) *boxv1.BoxApiV1RouterGroup {
	return &boxv1.BoxApiV1RouterGroup{RouterGroup: server.Group("/api/box/v1")}
}
