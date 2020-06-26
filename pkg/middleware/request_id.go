package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

const HeaderRequestId = "X-Request-Id"

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestId string
		if requestId = ctx.GetHeader(HeaderRequestId); requestId == "" {
			requestId = strings.ReplaceAll(uuid.New().String(), "-", "")
		}
		ctx.Request.Header.Set(HeaderRequestId, requestId)
		ctx.Set(HeaderRequestId, requestId)
		ctx.Next()
	}
}
