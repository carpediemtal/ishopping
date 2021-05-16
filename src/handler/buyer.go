package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"ishopping/src/service"
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
