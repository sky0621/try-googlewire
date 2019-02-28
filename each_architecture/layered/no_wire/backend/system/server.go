package system

import (
	"fmt"
	"try-googlewire/each_architecture/layered/no_wire/backend/infrastructure/persistence"

	"github.com/labstack/echo"
)

type Server interface {
	Serve()
}

type server struct {
	env
}

func NewServer(e env) Server {
	// 永続化層アクセス用のマネージャを生成
	m, err := persistence.NewManager(e.m.dataSourceStr())
	if err != nil {
		panic(err)
	}
	fmt.Println(m.Ping())

	return &server{env: e}
}

func (s *server) Serve() {
	// https://echo.labstack.com/guide
	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
