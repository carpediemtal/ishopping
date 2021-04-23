package main

import (
	"ishopping/src/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("/api/login", handler.LoginHandler)
	r.POST("/api/register", handler.RegisterHandler)
	r.GET("/api/commodity_search", handler.CommoditySearchHandler)
	r.GET("/api/commodity_detail", handler.CommodityDetailHandler)
	r.GET("/api/category_list",handler.CategoryListHandler)

	// release2 index_category_channel
	r.GET("/api/index_category_channel", handler.VistorViewHandle)

	r.GET("/api/buyer/information", handler.BuyerDetailHandler) //以uid查buyer信息
	r.POST("/api/buyer/information_modify", handler.UpdateBuyerInfoHandler)
	r.GET("/api/seller/shop_information", handler.ShopDetailHandler)
	r.POST("/api/seller/shop_information_modify", handler.UpdateShopInfoHandler)

	auth := r.Group("/api")
	auth.Use(handler.AuthorizationHandler)
	auth.GET("/seller/order_list", handler.OrderListHandler)
	auth.POST("/commodity_add", handler.CommodityAddHandler)
	auth.POST("/seller/order_delivery",handler.OrderDeliveryHandler)

	if err := r.Run(":7001"); err != nil {
		log.Println(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:63342")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// auth.Use(authMiddleware.MiddlewareFunc())
// auth.GET("/refresh_token", authMiddleware.RefreshHandler)
// auth.GET("/query_buyer_by_id", queryBuyerByIdHandler)
// auth.GET("/query_buyer_by_username", queryBuyerByUsernameHandler)
// auth.POST("/update_buyer_info", updateBuyerInfoHandler)
