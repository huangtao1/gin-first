package handler

import (
	"com.mark/ginFirst/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//登陆处理
func UserLogin(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	log.Println(email, password)
}

//注册处理
func UserRegister(ctx *gin.Context) {
	var user model.UserModel
	if err := ctx.ShouldBind(&user); err != nil {
		log.Println("err:", err)
		ctx.String(http.StatusBadRequest, "输入数据不合法")
		return
	}
	user.Save()
	log.Printf("email is %s,password is %s,password-again is %s", user.Email, user.Password, user.PasswordAgain)
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "注册成功"})
	return
}
