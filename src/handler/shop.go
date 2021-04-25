package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"ishopping/src/service"
)

func ShopDetailHandler(c *gin.Context) {
	uid := c.GetInt("UserID")
	shop, err := service.GetShopProfileById(uid)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}

	b, err := json.Marshal(shop)
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

func UpdateShopInfoHandler(c *gin.Context) {
	type params struct {
		//Uid      int    `json:"uid"`
		ShopName string `json:"shopname"`
		Address  string `json:"address"`
		PhoneNum string `json:"seller_phonenumber"`
	}
	var p params
	err := c.BindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError: "+err.Error())
		return
	}

	uid := c.GetInt("UserID")
	err = service.UpdateShopInfo(uid, p.ShopName, p.Address, p.PhoneNum)
	if err != nil {
		JsonErr(c, "update shop info error: "+err.Error())
		return
	}

	JsonOK(c, gin.H{})
}
