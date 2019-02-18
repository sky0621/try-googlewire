package main

import (
	"try-googlewire/di/provider/rdb"

	"github.com/google/wire"
)

func initializeApp() (*app, func(), error) {
	wire.Build(
		rdb.Set,
	)
	return nil, nil, nil
}
