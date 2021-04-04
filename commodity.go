package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Commodity struct {
	Cid       int    `json:"cid" map:"cid"`
	Sid       int    `json:"sid" map:"sid"`
	Name      string `json:"name" map:"name"`
	Price     int    `json:"price" map:"price"`
	Sales     int    `json:"sales" map:"sales"`
	Inventory int    `json:"inventory" map:"inventory"`
	Caid      int    `json:"caid" map:"caid"`
}

func getCommodityProfileByCid(cid int) (commodity Commodity, err error) {
	row1 := db.QueryRow("select * from commodity where cid = ?", cid)
	err = row1.Scan(&commodity.Cid, &commodity.Sid, &commodity.Name, &commodity.Price, &commodity.Sales, &commodity.Inventory, &commodity.Caid)
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
	content := c.Query("content")
	commodities, err := getCommodities()
	if err != nil {
		JsonErr(c, err.Error())
	}

	var ans []Commodity
	for _, commodity := range commodities {
		if strings.Contains(commodity.Name, content) {
			ans = append(ans, commodity)
		}
	}

	b, err := json.Marshal(ans)
	if err != nil {
		panic(err)
	}
	log.Println(string(b))
	JsonOK(c, gin.H{"result": string(b)})
}

func getCommodities() (commodities []Commodity, err error) {
	err = db.Select(&commodities, `select * from commodity`)
	return
}
