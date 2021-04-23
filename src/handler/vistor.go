package handler

import (
	"ishopping/src/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

func VistorViewHandle(c *gin.Context) {
	Caid, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		JsonErr(c, "no commodity found")
		return
	}
	Page, err := strconv.Atoi(c.Query("page_num"))
	if err != nil {
		JsonErr(c, "no commodity found")
		return
	}
	ans, err := service.GetCommodityListByCategoryAndPage(Caid, Page)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}
	JsonOK(c, gin.H{"commodity_list": ans})
}