package system

import (
	"github.com/labstack/echo"
)

type Server interface {
	Serve()
}

type server struct {
	set Settings
}

func NewServer(set Settings) Server {
	//// 永続化層アクセス用のマネージャを生成
	//m, err := persistence.NewManager(e.m.dataSourceStr())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(m.Ping())

	return &server{set: set}
}

func (s *server) Serve() {
	// https://echo.labstack.com/guide
	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
