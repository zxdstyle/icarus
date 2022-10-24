package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func Connect(cfg Config) (*gorm.DB, error) {
	orm, err := gorm.Open(mysql.Open(cfg.DataSource), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	//if !env.IsProduction() {
	// show sql statements
	orm = orm.Debug()
	//}
	//
	//err = orm.Use(dbresolver.Register(dbresolver.Config{
	//	Sources:  []gorm.Dialector{mysql.Open(cfg.DataSource)},
	//	Replicas: []gorm.Dialector{mysql.Open(cfg.SlaveDataSource)},
	//}).SetMaxOpenConns(cfg.MaxOpenConns).SetMaxIdleConns(cfg.MaxIdleConns))
	//if err != nil {
	//	return nil, err
	//}
	if err := orm.Use(dbresolver.Register(dbresolver.Config{}).SetMaxOpenConns(20).SetMaxIdleConns(10)); err != nil {
		return orm, err
	}
	return orm, nil
}
