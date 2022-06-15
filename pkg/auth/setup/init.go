// 用于配制和初始化 auth 模块
package setup

import (
	"github.com/juju/errors"
	"gorm.io/gorm"
	"seg/pkg/auth/models"
	"seg/pkg/auth/services"
	//"seg/pkg/base"
	"seg/pkg/base/rpc"
)

const MOD_NAME = "auth"

// 模块初始化
func Setup(reg *rpc.ServiceRegistry, db *gorm.DB) (err error) {
	err = errors.Trace(db.AutoMigrate(&models.User{}, &models.Student{}))
	if err != nil {
		return
	}

	err = reg.RegisterService(MOD_NAME, services.NewUserService(db))
	if err != nil {
		return
	}

	return
}
