package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Buyer struct {
	Uid      int    `json:"uid" map:"uid"`
	Name     string `json:"name" map:"uid"`
	Address  string `json:"address" map:"address"`
	PhoneNum string `json:"phone_num" map:"phone_num"`
}

func getBuyerProfileById(uid int) (buyer Buyer, err error) {
	row := db.QueryRow("select * from buyer where uid = ?", uid)
	err = row.Scan(&buyer.Uid, &buyer.Name, &buyer.Address, &buyer.PhoneNum)
	return
}

func getBuyerProfileByUsername(username string) (buyer Buyer, err error) {
	uid, err := getIdByUsername(username)
	if err != nil {
		return
	}
	return getBuyerProfileById(uid)
}

func queryBuyerById(c *gin.Context) {
	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		JsonErr(c, "user not found")
		return
	}
	buyer, err := getBuyerProfileById(uid)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}
	b, err := json.Marshal(buyer)
	if err != nil {
		panic(err)
	}
	var data gin.H
	err = json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}
	JsonOK(c, data)
}

func queryBuyerByUsername(c *gin.Context) {
	username := c.Query("username")
	buyer, err := getBuyerProfileByUsername(username)
	if err != nil {
		JsonErr(c, "user not found")
		return
	}
	b, err := json.Marshal(buyer)
	if err != nil {
		panic(err)
	}
	var data gin.H
	err = json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}
	JsonOK(c, data)
}
