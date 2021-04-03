package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	uid      int
	username string
	password string
	usertype int
}

func loginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	u, err := gerUserByUserName(username)
	if err != nil || u.password != password {
		JsonErr(c, "incorrect username or password")
		return
	}
	JsonOK(c, gin.H{})
}

func registerHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	userType, err := strconv.Atoi(c.PostForm("usertype"))
	if err != nil {
		JsonErr(c, "invalid user type")
		return
	}
	err = register(username, password, userType)
	if err != nil {
		JsonErr(c, "register failed:"+err.Error())
		return
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

func gerUserByUserName(username string) (User, error) {
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
