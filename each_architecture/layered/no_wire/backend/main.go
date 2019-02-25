package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// MEMO: ここにあるコードはプロダクションレベルのコードではありません。

// MEMO: このシステムでは永続化方法はMySQLへの接続で行う。
// ※接続先やある程度の接続パラメータは環境により動的に変更できるようにするが、
// 他のRDBならびにNoSQLへの変更におけるソース修正コストを最小限にすることまでは求めない。

func main() {
	// 環境変数よりMySQL接続情報を取得
	set := NewMySQLSetting()
	fmt.Println(set)

	// MySQLへの接続インスタンスを取得
	db, err := gorm.Open("mysql", set.DataSourceStr)
	if err != nil {
		panic(err)
	}
}

type MySQLSetting struct {
	DataSourceStr string
	MaxIdleConns  int
	MaxOpenConns  int
}

func NewMySQLSetting() MySQLSetting {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	instance := os.Getenv("DB_INSTANCE")
	dbName := os.Getenv("DB_NAME")
	set := MySQLSetting{
		DataSourceStr: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, instance, dbName),
	}
	return set
}

// 本来なら環境変数は各環境にて事前（もしくはCloudであればインスタンスセットアップ時？）にセットアップするけど、
// 今回は基本的にローカル環境での試行用なので、明示的にここでセットしてしまう。
func init() {
	os.Setenv("DB_USER", "localuser")
	os.Setenv("DB_PASS", "localpass")
	os.Setenv("DB_INSTANCE", "127.0.0.1:3306")
	os.Setenv("DB_NAME", "localdb")
}
