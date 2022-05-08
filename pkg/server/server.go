package server

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"42tokyo-road-to-dojo-go/pkg/http/middleware"
	"42tokyo-road-to-dojo-go/pkg/infra/sql"
	"42tokyo-road-to-dojo-go/pkg/presen/handler"
	"42tokyo-road-to-dojo-go/pkg/wire"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {

	/* ===== URLマッピングを行う ===== */
	http.HandleFunc("/setting/get", get(handler.HandleSettingGet()))
	db := sql.NewDriver()
	userHandler := wire.InitUserHandler(db)
	userCreate := http.HandlerFunc(userHandler.Create)
	http.HandleFunc("/user/create", post(middleware.Logger(userCreate)))
	userGet := http.HandlerFunc(userHandler.Get)
	http.HandleFunc("/user/get", get(middleware.Logger(userGet)))
	// TODO: 認証を行うmiddlewareを実装する
	// middlewareは pkg/http/middleware パッケージを利用する
	// http.HandleFunc("/user/get",
	//   get(middleware.Authenticate(handler.HandleUserGet())))

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
