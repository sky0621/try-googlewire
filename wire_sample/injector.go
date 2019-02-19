// buildタグを付与
//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

// 実装
func setUp(ctx context.Context, name string) (Foo, error) {
	// ProviderをBuildする
	wire.Build(SuperSet)
	// 特に意味はない
	return Foo{}, nil
}
