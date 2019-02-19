//+build wireinject

// 「wireinject」というビルドタグを付けると、「wire」コマンドによってDIソースが自動生成される

package main

import "github.com/google/wire"

// 依存する構造体群を生成する関数をすべて渡して、最終的にEventを構築
func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
