package handler

import (
	"ishopping/src/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CartAddHandler(c *gin.Context) {
	type params struct {
		Commodity_id int `json:"commodity_id"`
		Count        int `json:"count"`
	}
	var p params
	err := c.BindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError: "+err.Error())
		return
	}
	uid := c.GetInt("UserID")
	err = service.CartAddCommodity(uid, p.Commodity_id, p.Count)
	if err != nil {
		JsonErr(c, "Cart add error: "+err.Error())
		return
	}

	JsonOK(c, gin.H{})
}

func CartDeleteHandler(c *gin.Context) {
	type params struct {
		Commodity_id int `json:"commodity_id"`
	}
	var p params
	err := c.BindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError: "+err.Error())
		return
	}
	uid := c.GetInt("UserID")
	err = service.CartDeleteCommodity(uid, p.Commodity_id)
	if err != nil {
		JsonErr(c, "Cart delete error: "+err.Error())
		return
	}

	JsonOK(c, gin.H{})
}

func CartGetHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		JsonErr(c, "Page not found")
		return
	}
	size, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		JsonErr(c, "Page_size not found")
		return
	}
	uid := c.GetInt("UserID")
	ret, err := service.CartGetCommodityByUid(uid)
	if err != nil {
		JsonErr(c, "Cart search error: "+err.Error())
		return
	}
	if (page-1)*size >= len(ret) {
		JsonErr(c, "page out of range")
		return
	}

	var ans []service.CartCommodity
	for i, cnt := (page-1)*size, 0; i < len(ret) && cnt < size; i, cnt = i+1, cnt+1 {
		ans = append(ans, ret[i])
		//cnt++
	}

	JsonOK(c, gin.H{"list": ans})

}

func OrderHistoryHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		JsonErr(c, "Page not found")
		return
	}
	size, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		JsonErr(c, "Page_size not found")
		return
	}
	uid := c.GetInt("UserID")
	ret, err := service.GetOrderHistoryByUid(uid)
	if err != nil {
		JsonErr(c, "Cart search error: "+err.Error())
		return
	}
	if (page-1)*size >= len(ret) {
		JsonErr(c, "page out of range")
		return
	}

	var ans []service.OrderHistory
	for i, cnt := (page-1)*size, 0; i < len(ret) && cnt < size; i, cnt = i+1, cnt+1 {
		ans = append(ans, ret[i])
	}

	JsonOK(c, gin.H{"list": ans})
}
