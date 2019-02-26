package main

import (
	"fmt"
	"os"
)

type env struct {
	m mysql
}

type mysql struct {
	user     string
	password string
	instance string
	dbName   string
}

func (m mysql) dataSourceStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		m.user, m.password, m.instance, m.dbName)
}

func ReadEnv() env {
	e := env{}
	e.m.user = os.Getenv("DB_USER")
	e.m.password = os.Getenv("DB_PASS")
	e.m.instance = os.Getenv("DB_INSTANCE")
	e.m.dbName = os.Getenv("DB_NAME")
	return e
}
