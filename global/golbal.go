package global

import (
	"github.com/nhh57/go-ecommerce-backend-api/pkg/logger"
	"github.com/nhh57/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
