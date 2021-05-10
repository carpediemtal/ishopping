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
	r.GET("/api/category_list", handler.CategoryListHandler)

	// release2 index_category_channel
	r.GET("/api/index_category_channel", handler.VisitorViewHandler)

	auth := r.Group("/api")
	auth.Use(handler.AuthorizationHandler)
	auth.GET("/seller/order_list", handler.OrderListHandler)
	auth.POST("/seller/order_delivery", handler.OrderDeliveryHandler)
	auth.POST("/seller/commodity_edit", handler.CommodityEditHandler)
	auth.GET("/seller/shop_information", handler.ShopDetailHandler)
	auth.POST("/seller/shop_information_modify", handler.UpdateShopInfoHandler)

	auth.GET("/buyer/information", handler.BuyerDetailHandler)
	auth.POST("/buyer/information_modify", handler.UpdateBuyerInfoHandler)
	auth.POST("/buyer/commodity_buy", handler.BuyCommodityToOrderHandler)
	auth.POST("/userType", handler.UserTypeHandler)

	if err := r.Run(":7001"); err != nil {
		log.Println(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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
