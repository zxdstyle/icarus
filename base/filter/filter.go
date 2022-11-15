package filter

import (
	"fmt"
	"gorm.io/gorm"
)

type Filter struct {
	Eq      string `query:"eq"`
	Gt      string `query:"gt"`
	Gte     string `query:"gte"`
	Lt      string `query:"lt"`
	Lte     string `query:"lte"`
	Like    string `query:"like"`
	Contain string `query:"contain"`
}

func (f Filter) Filter(tx *gorm.DB, filed string) (res *gorm.DB) {
	res = tx
	if len(f.Eq) > 0 {
		res = tx.Where(fmt.Sprintf("`%s` = ?", filed), f.Eq)
	}

	if len(f.Gt) > 0 {
		res = tx.Where(fmt.Sprintf("`%s` > ?", filed), f.Gt)
	}

	if len(f.Gte) > 0 {
		res = tx.Where(fmt.Sprintf("`%s` >= ?", filed), f.Gte)
	}

	if len(f.Lt) > 0 {
		res = tx.Where(fmt.Sprintf("`%s` < ?", filed), f.Lt)
	}

	if len(f.Lte) > 0 {
		res = tx.Where(fmt.Sprintf("`%s` <= ?", filed), f.Lte)
	}

	if len(f.Like) > 0 {
		res = tx.Where(fmt.Sprintf("`%s` LIKE ?", filed), fmt.Sprintf("%s%%", f.Like))
	}

	if len(f.Contain) > 0 {
		res = tx.Where(fmt.Sprintf("`%s` LIKE ?", filed), fmt.Sprintf("%%%s%%", f.Contain))
	}

	return
}
