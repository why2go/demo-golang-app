package util

import (
	"context"

	"github.com/gin-gonic/gin"
)

func CtxWithRequestId(ctx *gin.Context) context.Context {
	requestId := ctx.GetString("request-id")
	return context.WithValue(ctx, "requestId", requestId)
}
