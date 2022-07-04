package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vzjxif/blog-service/global"
	"github.com/vzjxif/blog-service/internal/service"
	"github.com/vzjxif/blog-service/pkg/app"
	"github.com/vzjxif/blog-service/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	global.Logger.Infof("===== auth param = %.v", param)
	if !valid {
		global.Logger.Errorf("app.bindAndValid errsL:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Error()))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.checkAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.generateToken err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})

}
