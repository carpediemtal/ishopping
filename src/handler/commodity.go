package handler

import (
	"encoding/json"
	"errors"
	"ishopping/src/service"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	CommodityUpdate = 0
	CommodityAdd    = 1
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
	cid, err := strconv.Atoi(c.Query("commodity_id"))
	if err != nil {
		JsonErr(c, "Commodity not found")
		return
	}

	detail, err := service.GetCommodityDetailByCid(cid)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}

	b, err := json.Marshal(detail)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}

	var data gin.H
	if err = json.Unmarshal(b, &data); err != nil {
		JsonErr(c, err.Error())
		return
	}
	JsonOK(c, data)
}

func CategoryListHandler(c *gin.Context) {
	categories, err := service.GetAllCategories()
	if err != nil {
		JsonErr(c, err.Error())
	}

	JsonOK(c, gin.H{"category_list": categories})
}

func CommodityEditHandler(c *gin.Context) {
	var p service.CommodityEdit
	if err := c.ShouldBindJSON(&p); err != nil {
		JsonErr(c, err.Error())
		return
	}

	uid := c.GetInt("UserID")
	switch p.EditType {
	case CommodityUpdate:
		if err := service.UpdateCommodityInfo(p); err != nil {
			JsonErr(c, err.Error())
			return
		}
	case CommodityAdd:
		if err := service.AddCommodity(p, uid); err != nil {
			JsonErr(c, err.Error())
			return
		}
	default:
		JsonErr(c, errors.New("invalid edit_type").Error())
		return
	}

	JsonOK(c, gin.H{})
}

func BuyCommodityToOrderHandler(c *gin.Context) {
	type params struct {
		Cid int `json:"commodity_id"`
	}
	var p params
	err := c.BindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError: "+err.Error())
		return
	}
	uid := c.GetInt("UserID")
	pur_order, err := service.CreatPurchaseOrderByCid(uid, p.Cid)
	if err != nil {
		JsonErr(c, "purchase error: "+err.Error())
		return
	}

	b, err := json.Marshal(pur_order)
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
