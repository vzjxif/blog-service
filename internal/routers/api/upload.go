package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vzjxif/blog-service/global"
	"github.com/vzjxif/blog-service/internal/service"
	"github.com/vzjxif/blog-service/pkg/app"
	"github.com/vzjxif/blog-service/pkg/convert"
	"github.com/vzjxif/blog-service/pkg/errcode"
	"github.com/vzjxif/blog-service/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.uploadFile err:%v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccseeUrl,
	})
}
