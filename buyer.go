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

func queryBuyerByIdHandler(c *gin.Context) {
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

func queryBuyerByUsernameHandler(c *gin.Context) {
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

func updateBuyerInfoHandler(c *gin.Context) {
	type params struct {
		Uid      int    `json:"uid"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		PhoneNum string `json:"phone_num"`
	}
	var p params
	err := c.BindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError: "+err.Error())
		return
	}

	err = updateBuyerInfo(p.Uid, p.Name, p.Address, p.PhoneNum)
	if err != nil {
		JsonErr(c, "update buyer info error: "+err.Error())
		return
	}

	JsonOK(c, gin.H{})
}

func updateBuyerInfo(uid int, name, address, phoneNum string) (err error) {
	_, err = getBuyerProfileById(uid)
	// 如果买家的信息已经存在
	if err == nil {
		_, err = db.Exec(`update buyer set name = ?, address = ?, phone_num = ? where uid = ?`, name, address, phoneNum, uid)
	} else {
		_, err = db.Exec(`insert into buyer (uid, name, address, phone_num) values (?, ?, ?, ?)`, uid, name, address, phoneNum)
	}
	return
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
