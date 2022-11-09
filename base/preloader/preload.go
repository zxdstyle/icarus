package preloader

import (
	"github.com/zxdstyle/icarus/helpers"
	"gorm.io/gorm"
)

type (
	Preloader interface {
		Preload(tx *gorm.DB) *gorm.DB
	}

	defaultPreloader struct {
		resourceName string
	}
)

func DefaultPreloader(name string) Preloader {
	return &defaultPreloader{resourceName: helpers.ToCamelWithoutDot(name)}
}

func (p defaultPreloader) Preload(tx *gorm.DB) *gorm.DB {
	return tx.Preload(p.resourceName)
}
