package handler

import (
	"encoding/json"
	"ishopping/src/service"

	"github.com/gin-gonic/gin"
)

func BuyerDetailHandler(c *gin.Context) {
	uid := c.GetInt("UserID")
	buyer, err := service.GetBuyerProfileById(uid)
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

func UpdateBuyerInfoHandler(c *gin.Context) {
	type params struct {
		//Uid      int    `json:"uid"`
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

	uid := c.GetInt("UserID")
	err = service.UpdateBuyerInfo(uid, p.Name, p.Address, p.PhoneNum)
	if err != nil {
		JsonErr(c, "update buyer info error: "+err.Error())
		return
	}

	JsonOK(c, gin.H{})
}

func BuyerEvaluateHandler(c *gin.Context) {
	type params struct {
		Order_id int    `json:"order_id"`
		Rate     int    `json:"rate"`
		Content  string `json:"content"`
	}
	var p params
	err := c.BindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError: "+err.Error())
		return
	}
	if p.Rate > 5 || p.Rate < 1 {
		JsonErr(c, "The rate is nonstandard!")
		return
	}
	uid := c.GetInt("UserID")
	err = service.BuyerEvaluateCommodity(uid, p.Order_id, p.Rate, p.Content)
	if err != nil {
		JsonErr(c, "Buyer evaluate error: "+err.Error())
		return
	}

	JsonOK(c, gin.H{})
}
