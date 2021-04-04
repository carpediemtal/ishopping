package main

import (
	"encoding/json"
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

	JsonOK(c, gin.H{"commodity_list": ans})
}

func getCommodities() (commodities []Commodity, err error) {
	err = db.Select(&commodities, `select * from commodity`)
	return
}

func commodityAddHandler(c *gin.Context) {
	type params struct {
		ID           int    `json:"user_id"`
		Name         string `json:"commodity_name"`
		Inventory    int    `json:"inventory"`
		Introduction string `json:"introduction"`
		Price        int    `json:"price"`
		//CaID         int    `json:"caid"`
	}
	var p params
	err := c.BindJSON(&p)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}

	err = addCommodity(p.ID, p.Price, p.Inventory, p.Name, p.Introduction)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}

	JsonOK(c, gin.H{})
}

func addCommodity(uid, price, inventory int, name, introduction string) error {
	sid, err := getSidByUid(uid)
	if err != nil {
		return err
	}

	// TODO: sales 和 caid 先插0，以后的版本再改
	_, err = db.Exec(`insert into commodity (sid, name, price, sales, inventory, caid) VALUES (?, ?, ?, ?, ?, ?)`, sid, name, price, 0, inventory, 0)
	if err != nil {
		return err
	}

	cid, err := getCidBySid(sid)
	if err != nil {
		return err
	}
	_, err = db.Exec(`insert into commodity_meta (cid, meta_key, meta_val) values (?, ?, ?)`, cid, "introduction", introduction)
	return err
}

func getSidByUid(uid int) (sid int, err error) {
	row := db.QueryRow(`select sid from shop where uid = ?`, uid)
	err = row.Scan(&sid)
	return
}

func getCidBySid(sid int) (cid int, err error) {
	row := db.QueryRow(`select cid from commodity where sid = ?`, sid)
	err = row.Scan(&cid)
	return
}
