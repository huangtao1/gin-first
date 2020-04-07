package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func IndexFuncHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.gohtml", gin.H{"title": "hello gin " + strings.ToLower(ctx.Request.Method) + " method"})
}
func RetHelloGin(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello gin "+strings.ToLower(ctx.Request.Method))
}
