package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vzjxif/blog-service/pkg/app"
	"github.com/vzjxif/blog-service/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := ctx.GetQuery("token"); exist {
			token = s
		} else {
			token = ctx.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}

		}

		if ecode != errcode.Success {
			response := app.NewResponse(ctx)
			response.ToErrorResponse(ecode)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
