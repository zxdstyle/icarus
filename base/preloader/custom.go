package preloader

import (
	"github.com/zxdstyle/icarus/helpers"
	"gorm.io/gorm"
	"strings"
)

type customPreload struct {
	customClause func(tx *gorm.DB) *gorm.DB
	resourceName string
}

func NewCustom(resource string, customClause func(tx *gorm.DB) *gorm.DB) Preloader {
	return &customPreload{
		customClause: customClause,
		resourceName: helpers.ToCamelWithoutDot(resource),
	}
}

func (p customPreload) Preload(tx *gorm.DB, fields ...string) *gorm.DB {
	return tx.Preload(p.resourceName, func(query *gorm.DB) *gorm.DB {
		query = p.customClause(query)
		if len(fields) == 0 {
			return query
		}
		return query.Select(strings.Join(fields, ","))
	})
}

func (p customPreload) Resource() string {
	return p.resourceName
}
