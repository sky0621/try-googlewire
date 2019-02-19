package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"wire_sample2/entity"

	"github.com/google/wire"
)

var SuperSet = wire.NewSet(entity.ProviderFoo)

func main() {
	n := flag.String("n", "foo", "foo name")
	flag.Parse()

	foo, err := setUp(context.Background(), *n)
	if err != nil {
		log.Fatalln(err)
	}

	// fooの名前を出力
	fmt.Println(foo.Name)
}
