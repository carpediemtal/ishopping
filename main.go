package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	r.POST("/api/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		u, err := gerUserByUserName(username)
		if err != nil || u.password != password {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "incorrect username or password", "data": ""})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": ""})
	})

	r.POST("/api/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		userType, err := strconv.Atoi(c.PostForm("usertype"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "invalid user type", "data": ""})
			return
		}
		err = register(username, password, userType)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -2, "msg": "register failed:" + err.Error(), "data": ""})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": ""})
	})

	err := r.Run()
	if err != nil {
		log.Println(err)
	}
}

func gerUserByUserName(username string) (User, error) {
	db := sqlx.MustConnect("mysql", "root:asdf;lkj@(localhost:3306)/ishopping")
	row := db.QueryRow("select * from user where username = ?", username)
	var user User
	err := row.Scan(&user.uid, &user.username, &user.password, &user.usertype)
	return user, err
}

func register(username, password string, userType int) error {
	if len(username) == 0 || len(password) == 0 {
		return errors.New("invalid username or password")
	}
	db := sqlx.MustConnect("mysql", "root:asdf;lkj@(localhost:3306)/ishopping")
	regSql := `insert into user (username ,password, type) values (?, ?, ?)`
	_, err := db.Exec(regSql, username, password, userType)
	return err
}

type User struct {
	uid      int
	username string
	password string
	usertype int
}
