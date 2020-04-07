package test

import (
	"bytes"
	"com.mark/ginFirst/initRouter"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

//测试登陆接口
func TestUserRegister(t *testing.T) {
	value := url.Values{}
	value.Add("email", "huangtao@oraro.net")
	value.Add("password", "2e31e13ej")
	value.Add("password-again", "2e31e13")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
