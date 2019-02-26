package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewManager(dataSourceStr string) (Manager, error) {
	// MySQLへの接続インスタンスを取得
	db, err := gorm.Open("mysql", dataSourceStr)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &manager{db: db}, nil
}

type Manager interface {
	Ping() error
}

type manager struct {
	db *gorm.DB
}

func (m *manager) Ping() error {
	if err := m.db.DB().Ping(); err != nil {
		return err
	}
	return nil
}
