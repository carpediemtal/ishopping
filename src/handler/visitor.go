package handler

import (
	"github.com/gin-gonic/gin"
	"ishopping/src/service"
	"strconv"
)

<<<<<<< HEAD:src/handler/visitor.go
func VisitorViewHandle(c *gin.Context) {
=======
func VisitorViewHandler(c *gin.Context) {
>>>>>>> b4516d78f659b8ac2274475b821f49e02cfa2338:src/handler/vistor.go
	Caid, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		JsonErr(c, "invalid caid")
		return
	}

	Page, err := strconv.Atoi(c.Query("page_num"))
	if err != nil {
		JsonErr(c, "invalid page_num")
		return
	}

	ans, err := service.GetCommodityListByCategoryAndPage(Caid, Page)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}
	JsonOK(c, gin.H{"commodity_list": ans})
}
