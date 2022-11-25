package preloader

import (
	"github.com/zxdstyle/icarus/helpers"
	"gorm.io/gorm"
)

type preloadWithClause struct {
	clause       string
	resourceName string
	values       []any
}

func WithClause(resource string, args ...any) Preloader {
	return &preloadWithClause{
		resourceName: helpers.ToCamelWithoutDot(resource),
		values:       args,
	}
}

func (p preloadWithClause) Preload(tx *gorm.DB, fields ...string) *gorm.DB {
	return tx.Preload(p.resourceName, p.values...)
}

func (p preloadWithClause) Resource() string {
	return p.resourceName
}
