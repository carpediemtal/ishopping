package main

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Commodity struct {
	Cid       int     `json:"cid" map:"cid"`
	Sid       int     `json:"sid" map:"sid"`
	Name      string  `json:"name" map:"name"`
	Price     float64 `json:"price" map:"price"`
	Sales     int     `json:"sales" map:"sales"`
	Inventory int     `json:"inventory" map:"inventory"`
	Caid      int     `json:"caid" map:"caid"`
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
	if err = json.Unmarshal(b, &data); err != nil {
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
			log.Println("Whats this all about ?",commodity.Name, content)
			ans = append(ans, commodity)
		}
	}

	JsonOK(c, gin.H{"commodity_list": ans})
}

func commodityAddHandler(c *gin.Context) {
	type params struct {
		Name         string  `json:"commodity_name"`
		Inventory    int     `json:"inventory"`
		Introduction string  `json:"introduction"`
		Price        float64 `json:"price"`
		//CaID         int    `json:"caid"`
	}

	var p params
	if err := c.ShouldBindJSON(&p); err != nil {
		JsonErr(c, err.Error())
		return
	}
	userID := c.GetInt("UserID")
	cid, err := addCommodity(userID, p.Price, p.Inventory, p.Name, p.Introduction)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}

	JsonOK(c, gin.H{"commodity_id": cid})
}

func getCommodityProfileByCid(cid int) (commodity Commodity, err error) {
	row1 := db.QueryRow("select * from commodity where cid = ?", cid)
	err = row1.Scan(&commodity.Cid, &commodity.Sid, &commodity.Name, &commodity.Price, &commodity.Sales, &commodity.Inventory, &commodity.Caid)
	return
}

func getCommodities() (commodities []Commodity, err error) {
	err = db.Select(&commodities, `select * from commodity`)
	return
}

// 添加商品成功时返回商品ID
func addCommodity(uid int, price float64, inventory int, name, introduction string) (cid int, err error) {
	sid, err := getSidByUid(uid)
	if err != nil {
		return 0, errors.New("no matched shop id")
	}

	// TODO: sales 和 caid 先插0，以后的版本再改
	_, err = db.Exec(`insert into commodity (sid, name, price, sales, inventory, caid) VALUES (?, ?, ?, ?, ?, ?)`, sid, name, price, 0, inventory, 0)
	if err != nil {
		return 0, err
	}

	cid, err = getCidBySid(sid)
	if err != nil {
		return 0, err
	}
	_, err = db.Exec(`insert into commodity_meta (cid, meta_key, meta_val) values (?, ?, ?)`, cid, "introduction", introduction)
	return cid, err
}

func getSidByUid(uid int) (sid int, err error) {
	log.Println("uid:", uid)
	row := db.QueryRow(`select sid from shop where uid = ?`, uid)
	err = row.Scan(&sid)
	return
}

func getCidBySid(sid int) (cid int, err error) {
	row := db.QueryRow(`select cid from commodity where sid = ?`, sid)
	err = row.Scan(&cid)
	return
}
