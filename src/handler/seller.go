package handler

import (
	"errors"
	"ishopping/src/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrderListHandler(c *gin.Context) {
	status, err := strconv.Atoi(c.Query("order_status"))
	if err != nil {
		JsonErr(c, errors.New("order_status is required").Error())
		return
	}

	userID := c.GetInt("UserID")
	orderList, err := service.GetOrderListByStatus(status, userID)
	if err != nil {
		JsonErr(c, errors.New("searching err").Error())
		return
	}

	JsonOK(c, gin.H{"orderList": orderList})
}

// 将订单由未发货状态（1）改变为发货状态（2）
func OrderDeliveryHandler(c *gin.Context) {
	type params struct {
		Oid int `json:"order_id"`
	}
	var p params
	if err := c.ShouldBindJSON(&p); err != nil {
		JsonErr(c, err.Error())
		return
	}

	uid := c.GetInt("UserID")
	if err := service.OrderDelivery(p.Oid, uid); err != nil {
		JsonErr(c, err.Error())
		return
	}

	JsonOK(c, gin.H{})
}
