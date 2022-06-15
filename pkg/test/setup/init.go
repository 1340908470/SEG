// 用于配制和初始化 auth 模块
package setup

import (
	"github.com/juju/errors"
	"gorm.io/gorm"
	"seg/pkg/base/rpc"
	"seg/pkg/test/models"
	"seg/pkg/test/services"
)

const MOD_NAME = "test"

// 模块初始化
func Setup(reg *rpc.ServiceRegistry, db *gorm.DB) (err error) {
	err = errors.Trace(db.AutoMigrate(&models.A{}))
	if err != nil {
		return
	}

	err = reg.RegisterService(MOD_NAME, services.NewTestService(db))
	if err != nil {
		return
	}

	return
}
