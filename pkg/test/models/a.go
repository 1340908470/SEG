package models

import (
	"github.com/lantu-dev/puki/pkg/base"
)

type A struct {
	AID base.ID `gorm:"type:bigint;primaryKey;not null"`

	Name string
}
