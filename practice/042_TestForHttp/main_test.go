package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert" // 這是一個斷言庫，你也可以直接比較
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	// 獲取一個請求實例
	w := httptest.NewRecorder()
	// 構造請求
	// 參數依次是 請求方法、路由、參數
	req, _ := http.NewRequest("GET", "/ping", nil)
	// 執行
	router.ServeHTTP(w, req)
	// 斷言
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
	// 完了
}

// > go test .\yoytang\015_Test
