package http

import (
	"demo/internal/ctrler/http/ping"
	v1 "demo/internal/ctrler/http/v1"

	"github.com/why2go/gostarter/ginstarter"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func RegisterHandlers() {

	ginstarter.DefaultRouter.Use(otelgin.Middleware("demo-service"))

	api := ginstarter.DefaultRouter.Group("/api/")
	{
		api.GET("/ping", ping.PingCtrler.Ping)
	}

	apiV1 := api.Group("v1")
	{
		apiV1.POST("/demo_user", v1.UserCtrler.GetDemoUser)
	}
}
