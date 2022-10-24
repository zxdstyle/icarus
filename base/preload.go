package base

import (
	"github.com/zxdstyle/liey/pkg/helpers"
	"gorm.io/gorm"
	"strings"
)

type (
	PreloadAble interface {
		Resource() string
		Preload(tx *gorm.DB, fields ...string) *gorm.DB
	}

	Preload struct {
		resourceName string
	}

	CustomPreload struct {
		customClause func(tx *gorm.DB) *gorm.DB
		resourceName string
	}
)

func NewPreload(resource string) PreloadAble {
	return &Preload{resourceName: helpers.ToCamelWithoutDot(resource)}
}

func NewCustomPreload(resource string, customClause func(tx *gorm.DB) *gorm.DB) PreloadAble {
	return &CustomPreload{
		customClause: customClause,
		resourceName: helpers.ToCamelWithoutDot(resource),
	}
}

func (p Preload) Preload(tx *gorm.DB, fields ...string) *gorm.DB {
	if len(fields) == 0 {
		return tx.Preload(p.resourceName)
	}

	return tx.Preload(p.resourceName, func(db *gorm.DB) *gorm.DB {
		return db.Select(strings.Join(fields, ","))
	})
}

func (p Preload) Resource() string {
	return p.resourceName
}

func (p CustomPreload) Preload(tx *gorm.DB, fields ...string) *gorm.DB {
	return tx.Preload(p.resourceName, func(query *gorm.DB) *gorm.DB {
		query = p.customClause(query)
		if len(fields) == 0 {
			return query
		}
		return query.Select(strings.Join(fields, ","))
	})
}

func (p CustomPreload) Resource() string {
	return p.resourceName
}
