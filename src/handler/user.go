package handler

import (
	"github.com/gin-gonic/gin"
	"ishopping/src/jwt"
	"ishopping/src/service"
	"log"
)

func LoginHandler(c *gin.Context) {
	type params struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var p params
	if err := c.ShouldBindJSON(&p); err != nil {
		JsonErr(c, err.Error())
	}

	u, err := service.GetUserByUsername(p.Username)
	if err != nil || u.Password != p.Password {
		JsonErr(c, "incorrect username or password")
		return
	}

	token, err := jwt.NewJWT(&jwt.CustomClaims{
		UserID:   u.Uid,
		Username: u.Username,
	})
	if err != nil {
		JsonErr(c, "failed to generate token")
		return
	}

	JsonOK(c, gin.H{
		"token": token,
	})
}

func RegisterHandler(c *gin.Context) {
	type params struct {
		Username string `json:"username"`
		Password string `json:"password"`
		UserType int    `json:"usertype"`
	}
	var p params
	if err := c.ShouldBindJSON(&p); err != nil {
		JsonErr(c, err.Error())
		return
	}

	if err := service.Register(p.Username, p.Password, p.UserType); err != nil {
		JsonErr(c, "register failed: username exists")
		return
	}

	// 给 seller 自动创建一个 shop
	if p.UserType == 1 {
		log.Println("create a shop for seller")
		user, err := service.GetUserByUsername(p.Username)
		if err != nil {
			JsonErr(c, err.Error())
			panic(err)
			return
		}

		if err = service.CreateShop(user.Uid, "default shop name"); err != nil {
			JsonErr(c, "create shop failed:"+err.Error())
			return
		}
	}

	JsonOK(c, gin.H{})
}
