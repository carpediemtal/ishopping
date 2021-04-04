package main

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Commodity struct {
	CID       int    `json:"cid,map:"cid"`
	SID       int    `json:"sid,map:"sid"`
	NAME      string `json:"name,map:"name"`
	PRICE     int    `json:"price,map:"price"`
	SALES     int    `json:"sales,map:"sales"`
	INVENTORY int    `json:"inventory,map:"inventory"`
	CAID      int    `json:"caid,map:"caid"`
	//INTRODUCTION string `json:"introduction,map:"introduction"`
}

func getCommodityProfileByCid(cid int) (commodity Commodity, err error) {
	row1 := db.QueryRow("select * from commodity where cid = ?", cid)
	//row2 := db.QueryRow("select content from comment where cid = ?", cid)
	err = row1.Scan(&commodity.CID, &commodity.SID, &commodity.NAME, &commodity.PRICE,
		&commodity.SALES, &commodity.INVENTORY, &commodity.CAID) //, &commodity.INTRODUCTION)
	//row2.Scan(&commodity.INTRODUCTION)

	return
}

func commodityDetailHandler(c *gin.Context) {
	cid, err := strconv.Atoi(c.Query("cid"))
	if err != nil {
		JsonErr(c, "Commodity not found")
		return
	}
	commodity, err := getCommodityProfileByCid(cid)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}
	b, err := json.Marshal(commodity)
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

func commoditySearchHandler(c *gin.Context) {
	// TODO: test
}
