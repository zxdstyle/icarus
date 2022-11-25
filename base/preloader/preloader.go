package preloader

import (
	"gorm.io/gorm"
)

type (
	Preloader interface {
		Resource() string
		Preload(tx *gorm.DB, fields ...string) *gorm.DB
	}
)
