package handler

import (
	"encoding/json"
	"ishopping/src/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShopDetailHandler(c *gin.Context) {
	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		JsonErr(c, "user not found")
		return
	}

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
		Uid      int    `json:"uid"`
		ShopName string `json:"shop_name"`
		Address  string `json:"address"`
		PhoneNum string `json:"phone_num"`
	}
	var p params
	err := c.BindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError: "+err.Error())
		return
	}

	err = service.UpdateShopInfo(p.Uid, p.ShopName, p.Address, p.PhoneNum)
	if err != nil {
		JsonErr(c, "update shop info error: "+err.Error())
		return
	}

	JsonOK(c, gin.H{})
}
