package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"web3-go-blog/models"
	"web3-go-blog/router"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = router.SetupRouter()
	models.InitDB()
}

// TestClearDB 清理测试数据库
func TestClearDB(t *testing.T) {
	models.ClearDB()
}

// TestRegister_Init 注册测试用户
func TestRegister_Init(t *testing.T) {
	register("zhangsan")
	register("lisi")
	register("wangwu")
	register("zhaoliu")
}

func register(name string) {
	body := bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s","password":"123456","email":"%s@example.com"}`,
		name, name)))
	req, _ := http.NewRequest("POST", "/api/register", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w)
}

// TestLogin 登录测试
func TestLogin(t *testing.T) {
	fmt.Println(login("zhangsan", "123456"))
}

func login(username, password string) string {
	body := bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, password)))
	req, _ := http.NewRequest("POST", "/api/login", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := make(map[string]interface{})
	json.Unmarshal(w.Body.Bytes(), &res)
	return res["data"].(string)
}

// TestCreatePost 创建文章测试
func TestCreatePost(t *testing.T) {
	token := login("zhangsan", "123456")
	body := bytes.NewBuffer([]byte(`{"title":"这是一篇测试文章","content":"这是一篇测试文章的内容"}`))
	req, _ := http.NewRequest("POST", "/api/post", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
}

// TestListPost 列出文章测试
func TestListPost(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/posts", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
}

// TestCreateComment 创建评论测试
func TestCreateComment(t *testing.T) {
	token := login("lisi", "123456")
	body := bytes.NewBuffer([]byte(`{"content":"这是一段测试评论"}`))
	req, _ := http.NewRequest("POST", "/api/post/1/comment", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
}

// TestCreateCommentMustLogin 创建文章必须登录
func TestCreatePostMustLogin(t *testing.T) {
	body := bytes.NewBuffer([]byte(`{"title":"这是一篇测试文章","content":"这是一篇测试文章的内容"}`))
	req, _ := http.NewRequest("POST", "/api/post", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String()) //{"code":-1,"message":"Unauthorized"}
}

// TestUpdatePostMustLogin 修改文章必须作者本人
func TestUpdatePostMustLogin(t *testing.T) {
	token := login("lisi", "123456")
	body := bytes.NewBuffer([]byte(`{"title":"这是一篇测试文章","content":"这是一篇测试文章的内容"}`))
	req, _ := http.NewRequest("PUT", "/api/post/1", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String()) //{"code":-1,"message":"permission denied"}
}

// TestDeletePostMustLogin 删除文章必须作者本人
func TestDeletePostMustLogin(t *testing.T) {
	token := login("lisi", "123456")
	req, _ := http.NewRequest("DELETE", "/api/post/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String()) //{"code":-1,"message":"permission denied"}
}
