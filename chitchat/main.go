package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ServerMux = HTTPのマルチプレクサ
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	// URLパターンと対応するhandlerをServerMuxの中に登録 → 内部的にhttp.Handler()と同じ処理をしている
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// HandleFuncは第一引数にURLをとり、第二引数にハンドラ関数名をとる
	mux.HandleFunc("/", index)

	mux.HandleFunc("/test", testHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello world")
}

func testHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("test test test")
}
