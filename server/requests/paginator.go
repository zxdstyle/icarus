package requests

import (
	"github.com/zxdstyle/icarus/server/responses"
	"gorm.io/gorm"
)

type (
	Paginator struct {
		Total    int64 `json:"total"`
		PageSize int   `json:"pageSize" query:"pageSize"`
		Page     int   `json:"page" query:"page"`
	}

	SimplePaginator struct {
		PageSize int `json:"pageSize" query:"pageSize"`
		Page     int `json:"page" query:"page"`
	}

	DefaultSorter struct {
		ID string `query:"id"`
	}

	DefaultFilter struct {
		ID string `query:"id"`
	}
)

func (p *Paginator) Paginate(tx *gorm.DB) (*gorm.DB, error) {
	if err := tx.Count(&p.Total).Error; err != nil {
		return nil, err
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 20
	}

	offset := (p.Page - 1) * p.PageSize
	return tx.Limit(p.PageSize).Offset(offset), nil
}

func (p *SimplePaginator) Paginate(tx *gorm.DB) (*gorm.DB, error) {
	offset := (p.Page - 1) * p.PageSize
	return tx.Limit(p.PageSize).Offset(offset), nil
}

func (p *Paginator) ToPaginationMeta() *responses.Meta {
	page := p.Page
	if page == 0 {
		page = 1
	}

	return &responses.Meta{Pagination: responses.PaginationMeta{
		CurrentPage: p.Page,
		PerPage:     p.PageSize,
		Total:       p.Total,
	}}
}
