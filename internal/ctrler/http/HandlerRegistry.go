package http

import (
	"demo/internal/ctrler/http/ping"
	v1 "demo/internal/ctrler/http/v1"

	"github.com/why2go/gostarter/ginstarter"
)

func RegisterHandlers() {
	api := ginstarter.DefaultRouter.Group("/api/")
	{
		api.GET("/ping", ping.PingCtrler.Ping)
	}

	apiV1 := api.Group("v1")
	{
		apiV1.POST("/echo", v1.UserCtrler.GetDemoUser)
	}
}
