package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"ishopping/src/service"
	"log"
	"strconv"
	"strings"
)

func CommoditySearchHandler(c *gin.Context) {
	content := c.Query("content")
	commodities, err := service.GetCommodities()
	if err != nil {
		JsonErr(c, err.Error())
	}

	var ans []service.Commodity
	for _, commodity := range commodities {
		if strings.Contains(commodity.Name, content) {
			log.Println("Whats this all about ?", commodity.Name, content)
			ans = append(ans, commodity)
		}
	}

	JsonOK(c, gin.H{"commodity_list": ans})
}

func CommodityDetailHandler(c *gin.Context) {
	cid, err := strconv.Atoi(c.Query("cid"))
	if err != nil {
		JsonErr(c, "Commodity not found")
		return
	}

	commodity, err := service.GetCommodityProfileByCid(cid)
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

func CommodityAddHandler(c *gin.Context) {
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
	cid, err := service.AddCommodity(userID, p.Price, p.Inventory, p.Name, p.Introduction)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}

	JsonOK(c, gin.H{"commodity_id": cid})
}


type order struct {
	Oid         int     `json:"oid" map:"oid"`
	Total_price float64 `json:"total_price" map:"total_price"`
}

func order_list_handler(c *gin.Context) {
	var status int
	err := c.Query("order_status")

	var order_list []order

	order_list, _ = getorderlistBystatus(status)
	JsonOK(c, gin.H{"order_list": order_list})
	if err != nil {
		JsonErr(c, err.Error())
		return
	}
}
func getorderlistBystatus(status int) (order_list []order, err error) {
	err = db.Select(&order_list, `select oid,sum(price) from commodity,purchase_order where commodity.cid=purchase_order.cid&&status=?`, status)
	return
}