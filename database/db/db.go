package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/database"
	"github.com/goravel/framework/contracts/database/db"
	contractsdriver "github.com/goravel/framework/contracts/database/driver"
	"github.com/goravel/framework/errors"
)

type DB struct {
	config   database.Config
	instance *sqlx.DB
}

func NewDB(config database.Config, instance *sqlx.DB) db.DB {
	return &DB{config: config, instance: instance}
}

func BuildDB(config config.Config, connection string) (db.DB, error) {
	driverCallback, exist := config.Get(fmt.Sprintf("database.connections.%s.via", connection)).(func() (contractsdriver.Driver, error))
	if !exist {
		return nil, errors.DatabaseConfigNotFound
	}

	driver, err := driverCallback()
	if err != nil {
		return nil, err
	}

	instance, err := driver.DB()
	if err != nil {
		return nil, err
	}

	return NewDB(driver.Config(), sqlx.NewDb(instance, driver.Config().Driver)), nil
}

func (r *DB) Table(name string) db.Query {
	return NewQuery(r.config, r.instance, name)
}
