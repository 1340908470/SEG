package models

import (
	"seg/pkg/base"
)

type A struct {
	AID base.ID `gorm:"type:bigint;primaryKey;not null"`

	Name string
}
