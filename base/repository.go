package base

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/strutil"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cast"
	"github.com/zxdstyle/icarus/base/filter"
	"github.com/zxdstyle/icarus/base/preloader"
	"github.com/zxdstyle/icarus/server/requests"
	"github.com/zxdstyle/icarus/server/responses"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"reflect"
	"strings"
)

type Repository[M any] interface {
	List(ctx context.Context, req requests.Request, response *responses.ApiResponse) error
	Show(ctx context.Context, primaryKey uint, req requests.Request, mo *M) error
	Create(ctx context.Context, mo *M) error
	BatchCreate(ctx context.Context, mos *[]M) error
	Update(ctx context.Context, primaryKey uint, mo *M) error
	Destroy(ctx context.Context, primaryKey uint) error

	First(ctx context.Context, primaryKey uint, mo *M) error
	Exists(ctx context.Context, primaryKey uint) (bool, error)
}

var _ Repository[RepoModel] = GormRepository[RepoModel, any, any]{}

type (
	GormRepository[M any, F any, O any] struct {
		Orm      *gorm.DB
		preloads map[string]preloader.Preloader
	}

	RepoModel interface {
		Filterable() any
	}
)

func NewGormRepository[M any, F any, O any](orm *gorm.DB) *GormRepository[M, F, O] {
	var mo M
	return &GormRepository[M, F, O]{
		Orm:      orm.Model(mo),
		preloads: make(map[string]preloader.Preloader),
	}
}

func NewGormRepositoryWithPreload[M any, F any, O any](orm *gorm.DB, preloads ...preloader.Preloader) *GormRepository[M, F, O] {
	var mo M
	ps := make(map[string]preloader.Preloader)
	for _, preload := range preloads {
		ps[preload.Resource()] = preload
	}

	return &GormRepository[M, F, O]{
		Orm:      orm.Model(mo),
		preloads: ps,
	}
}

func (g GormRepository[M, F, O]) List(ctx context.Context, req requests.Request, response *responses.ApiResponse) error {
	var query struct {
		requests.Paginator
		Where F
		Order O
	}
	if err := req.ScanQueries(&query); err != nil {
		return err
	}

	tx := g.Orm.WithContext(ctx)

	tx = g.doFilter(tx, query.Where)

	tx, err := query.Paginate(tx)
	if err != nil {
		return err
	}
	tx = g.doSorter(tx, query.Order)

	tx, err = g.doPreload(req, tx)
	if err != nil {
		return err
	}

	response.Meta = query.ToPaginationMeta()
	return tx.Find(response.Data).Error
}

func (g GormRepository[M, F, O]) Show(ctx context.Context, primaryKey uint, req requests.Request, mo *M) error {
	tx, err := g.doPreload(req, g.Orm.WithContext(ctx))
	if err != nil {
		return err
	}
	return tx.Where("`id` = ?", primaryKey).First(mo).Error
}

func (g GormRepository[M, F, O]) Create(ctx context.Context, mo *M) error {
	return g.Orm.WithContext(ctx).Create(mo).Error
}

func (g GormRepository[M, F, O]) BatchCreate(ctx context.Context, mos *[]M) error {
	return g.Orm.WithContext(ctx).CreateInBatches(mos, 100).Error
}

func (g GormRepository[M, F, O]) Update(ctx context.Context, primaryKey uint, mo *M) error {
	if err := g.Orm.WithContext(ctx).Where("`id` = ?", primaryKey).Updates(mo).Error; err != nil {
		return err
	}
	// 修改完立马查询强制使用主节点，以防数据库主从延迟导致查询到修改之前的数据
	return g.Orm.WithContext(ctx).Clauses(dbresolver.Write).First(mo, primaryKey).Error
}

func (g GormRepository[M, F, O]) Destroy(ctx context.Context, primaryKey uint) error {
	var mo M
	return g.Orm.WithContext(ctx).Where("`id` = ?", primaryKey).Delete(&mo).Error
}

func (g GormRepository[M, F, O]) First(ctx context.Context, primaryKey uint, mo *M) error {
	return g.Orm.WithContext(ctx).Where("`id` = ?", primaryKey).First(mo).Error
}

func (g GormRepository[M, F, O]) Exists(ctx context.Context, primaryKey uint) (exist bool, err error) {
	var count int64
	err = g.Orm.WithContext(ctx).Where("`id` = ?", primaryKey).Count(&count).Error
	return count > 0, err
}

func (g GormRepository[M, F, O]) doSorter(tx *gorm.DB, sorter O) *gorm.DB {
	t := reflect.TypeOf(sorter)
	v := reflect.ValueOf(sorter)
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Tag.Get("query")
		if len(name) == 0 {
			name = t.Field(i).Tag.Get("json")
		}
		if len(name) == 0 {
			name = strutil.SnakeCase(t.Field(i).Name)
		}
		value := v.Field(i).Interface()
		val := strings.ToUpper(cast.ToString(value))
		if len(val) == 0 || len(name) == 0 {
			continue
		}
		tx = tx.Order(fmt.Sprintf("`%s` %s", name, val))
	}
	return tx
}

func (g GormRepository[M, F, O]) doFilter(tx *gorm.DB, where F) *gorm.DB {
	t := reflect.TypeOf(where)
	v := reflect.ValueOf(where)
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("query")
		if len(fieldName) == 0 {
			fieldName = t.Field(i).Tag.Get("json")
		}
		if len(fieldName) == 0 {
			fieldName = strutil.SnakeCase(t.Field(i).Name)
		}
		fl, ok := v.Field(i).Interface().(filter.Filter)
		if !ok {
			continue
		}
		tx = fl.Filter(tx, fieldName)
	}
	return tx
}

func (g GormRepository[M, F, O]) doPreload(req requests.Request, tx *gorm.DB) (*gorm.DB, error) {
	for name, able := range g.preloads {
		resource := "with." + name
		value := req.GetQuery(resource)
		if len(value) == 0 {
			resource = "with." + strcase.ToSnake(name)
			value = req.GetQuery(resource)
		}
		if len(value) == 0 {
			continue
		}
		tx = able.Preload(tx, value)
	}
	return tx, nil
}
