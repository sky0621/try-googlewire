package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/google/wire"
)

var SuperSet = wire.NewSet(ProviderFoo)

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
