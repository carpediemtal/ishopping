package service

import (
	"errors"
	"ishopping/src/db"
)

type User struct {
	Uid      int
	Username string
	Password string
	usertype int8
}

func Register(username, password string, userType int) error {
	if len(username) == 0 || len(password) == 0 {
		return errors.New("invalid username or password")
	}
	regSql := `insert into user (username ,password, type) values (?, ?, ?)`
	_, err := db.DB.Exec(regSql, username, password, userType)
	return err
}

func GetUserByUsername(username string) (User, error) {
	row := db.DB.QueryRow("select * from ub_user where username = ?", username)
	var user User
	err := row.Scan(&user.Uid, &user.Username, &user.Password, &user.usertype)
	return user, err
}

func CreateShop(uid int, shopName string) error {
	_, err := db.DB.Exec(`insert into shop (uid, shop_name) VALUES (?, ?)`, uid, shopName)
	return err
}

func GetUserType(uid int) (userType int8, err error) {
	row := db.DB.QueryRow("select type from user where uid = ?", uid)
	err = row.Scan(&userType)
	return
}
