package model

import (
	"com.mark/ginFirst/utils"
	"log"
)

type UserModel struct {
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}

func (user *UserModel) Save() int64 {
	result, e := utils.Db.Exec("insert into users (email,password_hash) values (?,?);", user.Email, user.Password)
	if e != nil {
		log.Panic("insert into db failed:" + e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panic("get insert id failed:" + err.Error())
	}
	return id
}
