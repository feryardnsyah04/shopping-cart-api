package controllers

import (
	"net/http"
	"shopping-cart-api/models"
	"shopping-cart-api/services"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetCartItems(c *gin.Context) {
	items := services.GetCartItems()
	c.JSON(http.StatusOK, items)
}

func AddToCart(c *gin.Context) {
	var newItem models.CartItem
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.AddToCart(newItem)
	c.JSON(http.StatusCreated, newItem)
}

func UpdateCartItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var requestBody struct {
		Quantity int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateCartItem(uint(id), requestBody.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item berhasil diperbarui"})
}

func RemoveFromCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.RemoveFromCart(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item berhasil dihapus"})
}
