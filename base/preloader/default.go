package preloader

import (
	"github.com/zxdstyle/icarus/helpers"
	"gorm.io/gorm"
	"strings"
)

type defaultPreloader struct {
	resourceName string
}

func NewDefault(resourceName string) Preloader {
	return &defaultPreloader{resourceName: helpers.ToCamelWithoutDot(resourceName)}
}

func (p defaultPreloader) Preload(tx *gorm.DB, fields ...string) *gorm.DB {
	if len(fields) == 0 {
		return tx.Preload(p.resourceName)
	}

	return tx.Preload(p.resourceName, func(db *gorm.DB) *gorm.DB {
		return db.Select(strings.Join(fields, ","))
	})
}

func (p defaultPreloader) Resource() string {
	return p.resourceName
}
