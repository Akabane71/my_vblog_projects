package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	// 自己创建一个路由，防止循环引用
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	body := `{
		"community": 1,
		"title": "test",
		"content": "just a test"
	}`
	// 字节 ----->
	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// 判断响应内容是否

	// 方法1: 判断响应内容中是不是包含指定的字符串
	// assert.Contains(t, "pong", w.Body.String(), "需要登录")

	// 方法2: 将响应的内容反序列化到ResponseDate,然后判断字段与预期是否一致
	res := new(ResponseDate)
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatal("json.Unmarshal err:", err)
	}
	assert.Equal(t, res.Code, CodeNeedKLogin)
}
