package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	dataStore map[string]string
	mutex     sync.RWMutex
)

func main() {
	dataStore = make(map[string]string)

	// ルーティングの設定
	http.HandleFunc("/", handler)

	// サーバーの起動
	fmt.Println("サーバーを起動します... ポート 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("サーバーの起動に失敗しました:", err)
	}
}

// リクエストを処理するハンドラー関数
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		handlePut(w, r)
	case http.MethodGet:
		handleGet(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	// リクエストボディからデータを読み取る
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// データを保存
	key := "/"
	mutex.Lock()
	defer mutex.Unlock()
	dataStore[key] = string(body)

	// 成功を返す
	w.WriteHeader(http.StatusOK)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	key := "/"
	mutex.RLock()
	defer mutex.RUnlock()

	// データを取得
	data, ok := dataStore[key]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// レスポンスにデータを書き込む
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
