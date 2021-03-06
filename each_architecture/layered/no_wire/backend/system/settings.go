package system

// MEMO: システム全体の設定情報を担う。

func newSettings() Settings {
	return &settings{}
}

type Settings interface {
	GetSecuritySettings() SecuritySettings
	GetPersistenceSettings()
}

type settings struct {
	s SecuritySettings
	m MySQL
}

type SecuritySettings interface {
	GetAPIKey() string
}

type securitySettings struct {
	ApiKey string
}

type MySQL struct {
	User     string
	Password string
	Instance string
	DBName   string
}

// システムのセットアップを行う。
// 主に以下のようなロジックを想定。
// ・環境変数の読み込み
// ・設定ファイルの読み込み
// ・
func SetUp() Settings {
	s := &settings{}

	return s
}
