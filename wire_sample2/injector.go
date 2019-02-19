// buildタグを付与
//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"

	"wire_sample2/entity"
)

// 実装
func setUp(ctx context.Context, name string) (entity.Foo, error) {
	// ProviderをBuildする
	wire.Build(SuperSet)
	// 特に意味はない
	return entity.Foo{}, nil
}
