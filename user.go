package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	uid      int
	username string
	password string
	usertype int
}

func loginHandler(c *gin.Context) {
	type params struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var p params
	if err := c.BindJSON(&p); err != nil {
		JsonErr(c, err.Error())
	}

	u, err := getUserByUsername(p.Username)
	if err != nil || u.password != p.Password {
		JsonErr(c, "incorrect username or password")
		return
	}

	token, err := NewJWT(&CustomClaims{
		UserID:   u.uid,
		Username: u.username,
	})
	if err != nil {
		JsonErr(c, "failed to generate token")
		return
	}

	JsonOK(c, gin.H{
		"token": token,
	})
}

func registerHandler(c *gin.Context) {
	type params struct {
		Username string `json:"username"`
		Password string `json:"password"`
		UserType int    `json:"usertype"`
	}
	var p params
	if err := c.BindJSON(&p); err != nil {
		JsonErr(c, err.Error())
		return
	}

	if err := register(p.Username, p.Password, p.UserType); err != nil {
		JsonErr(c, "register failed: username exists")
		return
	}

	// 给 seller 自动创建一个 shop
	if p.UserType == 1 {
		log.Println("create a shop for seller")
		user, err := getUserByUsername(p.Username)
		if err != nil {
			JsonErr(c, err.Error())
			panic(err)
			return
		}

		if err = createShop(user.uid, "default shop name"); err != nil {
			JsonErr(c, "create shop failed:"+err.Error())
			return
		}
	}

	JsonOK(c, gin.H{})
}

func register(username, password string, userType int) error {
	if len(username) == 0 || len(password) == 0 {
		return errors.New("invalid username or password")
	}
	regSql := `insert into user (username ,password, type) values (?, ?, ?)`
	_, err := db.Exec(regSql, username, password, userType)
	return err
}

func getUserByUsername(username string) (User, error) {
	row := db.QueryRow("select * from user where username = ?", username)
	var user User
	err := row.Scan(&user.uid, &user.username, &user.password, &user.usertype)
	return user, err
}

func getIdByUsername(username string) (uid int, err error) {
	row := db.QueryRow("select uid from user where username = ?", username)
	err = row.Scan(&uid)
	return
}

func createShop(uid int, shopName string) error {
	_, err := db.Exec(`insert into shop (uid, shop_name) VALUES (?, ?)`, uid, shopName)
	return err
}
