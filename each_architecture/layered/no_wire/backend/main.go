package main

import (
	"os"
)

// Note: ここにあるコードはプロダクションレベルのコードではありません。
// Note: ログ出力やエラーハンドリングなどプロダクションレベルで必要なコードは省略しています。

func main() {
	// 環境変数より各種セットアップ情報を取得
	env := ReadEnv()

	// MEMO: toml等に書いた設定の読み込みも必要になるかも

	// MEMO: 各種メッセージの読み込みも必要になるかも

	// APIサーバのセットアップ
	NewServer(env).Serve()
}

// 本来なら環境変数は各環境にて事前（もしくはCloudであればインスタンスセットアップ時？）にセットアップするけど、
// 今回は基本的にローカル環境での試行用なので、明示的にここでセットしてしまう。
func init() {
	os.Setenv("API_KEY", "dummy")

	os.Setenv("DB_USER", "localuser")
	os.Setenv("DB_PASS", "localpass")
	os.Setenv("DB_INSTANCE", "127.0.0.1:3306")
	os.Setenv("DB_NAME", "localdb")
}
