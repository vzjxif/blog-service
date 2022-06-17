package global

import (
	"github.com/vzjxif/blog-service/pkg/logger"
	"github.com/vzjxif/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
