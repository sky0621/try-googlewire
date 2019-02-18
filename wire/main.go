package main

import (
	"try-googlewire/wire/entity"

	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	entity.ProvideUser(),
	entity.ProvideBook(),
	entity.ProvideBookshelf())

func main() {

}

//type App interface {
//}

type App struct {
}
