package main

import (
	"fmt"
	"os"
	"try-googlewire/each_architecture/layered/no_wire/backend/infrastructure/persistence"

	"github.com/labstack/echo"
)

// Note: ここにあるコードはプロダクションレベルのコードではありません。
// Note: ログ出力やエラーハンドリングなどプロダクションレベルで必要なコードは省略しています。

func main() {
	// 環境変数より各種セットアップ情報を取得
	env := ReadEnv()

	// 永続化層アクセス用のマネージャを生成
	m, err := persistence.NewManager(env.m.dataSourceStr())
	if err != nil {
		panic(err)
	}
	fmt.Println(m.Ping())

	// https://echo.labstack.com/guide
	e := echo.New()

}

// 本来なら環境変数は各環境にて事前（もしくはCloudであればインスタンスセットアップ時？）にセットアップするけど、
// 今回は基本的にローカル環境での試行用なので、明示的にここでセットしてしまう。
func init() {
	os.Setenv("DB_USER", "localuser")
	os.Setenv("DB_PASS", "localpass")
	os.Setenv("DB_INSTANCE", "127.0.0.1:3306")
	os.Setenv("DB_NAME", "localdb")
}
