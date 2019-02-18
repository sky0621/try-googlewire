//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

func setUp(ctx context.Context, name string) (*App, error) {
	// ProviderをBuildする
	wire.Build(SuperSet)
	// 特に意味はない
	return &App{}, nil
}
