package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

// https://echo.labstack.com/middleware
func setupBasic(e *echo.Echo, apiKey string) {

	// HTTPリクエストのログ出し
	// https://echo.labstack.com/middleware/logger
	e.Use(middleware.Logger())

	// パニックからの復帰
	// https://echo.labstack.com/middleware/recover
	e.Use(middleware.Recover())

	// クロスオリジン対応（ローカルではフロントとサーバとでポート分けて動作確認したりするので）
	// https://echo.labstack.com/middleware/cors
	e.Use(middleware.CORS())

	// リクエスト毎にユニークなIDをHTTPヘッダ「X-Request-ID」に積む
	// https://echo.labstack.com/middleware/request-id
	// TODO: どうも値が入らないのでいったん自前で積んでおくことにする
	//e.Use(middleware.RequestID())

	// 以下から保護
	// ・クロスサイトスクリプティング（XSS）攻撃
	// ・コンテンツタイプスニッフィング
	// ・クリックジャック
	// ・安全でない接続およびその他のコードインジェクション攻撃
	// https://echo.labstack.com/middleware/secure
	e.Use(middleware.Secure())

	// TODO: セキュリティ要件としてボディサイズ絞る必要あるかな？
	// https://echo.labstack.com/middleware/body-limit

	// TODO: クロスサイトリクエストフォージェリ対策は必要？
	// https://echo.labstack.com/middleware/csrf

	// https://echo.labstack.com/middleware/key-auth
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		fmt.Printf("get key: %s, apikey: %s\n", key, apiKey)
		return key == apiKey, nil
	}))

	// ロールによるアクセス制御は↓で多少なりとも実装が楽にならないか？
	// https://echo.labstack.com/middleware/casbin-auth

	// セッション機能有効化
	// https://echo.labstack.com/middleware/session
	// TODO: 使うかな？
	//e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// TODO: 静的ファイル使うなら。
	// https://echo.labstack.com/middleware/static

}

// SetupCustom ...
func SetupCustom(e *echo.Echo, appLgr logger.AppLogger, db *gorm.DB, firebaseApp *firebase.App) {
	// コンテキストにDB接続情報等を積んで引き回していくためのカスタマイズ
	e.Use(customContextMiddleware())

	// 順番がとても大事！
	e.Use(requestIDMiddleware())
	e.Use(customLoggerMiddleware(appLgr))
	e.Use(gormDBMiddleware(db))
}

// CustomContext ... Cloud SQLアクセッサ等をcontrollerで受け取れるよう、Echoコンテキストを拡張
type CustomContext interface {
	echo.Context
	GetLog() logger.AppLogger
	GetDB() *gorm.DB
	GetFirebaseApp() *firebase.App
}

type customContext struct {
	echo.Context
	log         logger.AppLogger
	db          *gorm.DB
	firebaseApp *firebase.App
	requestID   string
}

// GetDB ...
func (c *customContext) GetDB() *gorm.DB {
	return c.db
}

// GetCustomContext ...
func GetCustomContext(c echo.Context) CustomContext {
	return c.(*customContext)
}

func customContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &customContext{
				Context: c,
			}
			return next(cc)
		}
	}
}

func requestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cctx, ok := c.(*customContext)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			requestID, err := uuid.NewUUID()
			if err == nil {
				cctx.requestID = requestID.String()
			} else {
				cctx.log.Errorw(err.Error())
			}
			return next(cctx)
		}
	}
}

func gormDBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cctx, ok := c.(*customContext)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			cctx.db = db
			cctx.db.SetLogger(cctx.log)
			return next(cctx)
		}
	}
}
