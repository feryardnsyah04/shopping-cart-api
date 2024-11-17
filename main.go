package main

import (
	"shopping-cart-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "shopping-cart-api Koneksi Berhasil",
		})
	})

	router.GET("/cart", controllers.GetCartItems)

	router.POST("/cart", controllers.AddToCart)

	router.PUT("/cart/:id", controllers.UpdateCartItem)

	router.DELETE("/cart/:id", controllers.RemoveFromCart)

	router.Run(":8080")
}
