package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vzjxif/blog-service/pkg/app"
	"github.com/vzjxif/blog-service/pkg/errcode"
	"github.com/vzjxif/blog-service/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := l.Key(ctx)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(ctx)
				response.ToErrorResponse(errcode.TooManyRequests)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
