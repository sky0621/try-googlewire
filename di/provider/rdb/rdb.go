package rdb

import (
	"errors"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var Set = wire.NewSet(
	NewRDBConfig,
	NewRDB)

type RDBConfig struct {
	DriverName string
	DataSource string
}

func NewRDBConfig(driverName, dataSource string) RDBConfig {
	return RDBConfig{
		DriverName: driverName,
		DataSource: dataSource,
	}
}

type RDB struct {
	db *gorm.DB
}

func (r *RDB) Close() error {
	if r.db != nil {
		if err := r.db.Close(); err != nil {
			return err
		}
	}
	return nil
}

func NewRDB(config RDBConfig) (*RDB, error) {
	db, err := gorm.Open(config.DriverName, config.DataSource)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("can not connect to Cloud SQL")
	}
	return &RDB{db: db}, nil
}
