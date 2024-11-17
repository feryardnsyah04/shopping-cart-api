package services

import (
	"errors"
	"shopping-cart-api/database"
	"shopping-cart-api/models"
)

func GetCartItems() []models.CartItem {
	return database.CartItems
}

func AddToCart(item models.CartItem) {
	// Set ID item berdasarkan urutan
	item.ID = uint(len(database.CartItems) + 1)
	database.CartItems = append(database.CartItems, item)
}

func UpdateCartItem(id uint, quantity int) error {
	for i, item := range database.CartItems {
		if item.ID == id {
			if quantity <= 0 {
				return errors.New("kuantitas tidak valid")
			}
			database.CartItems[i].Quantity = quantity
			return nil
		}
	}
	return errors.New("produk tidak ditemukan")
}

func RemoveFromCart(id uint) error {
	for i, item := range database.CartItems {
		if item.ID == id {
			database.CartItems = append(database.CartItems[:i], database.CartItems[i+1:]...)
			return nil
		}
	}
	return errors.New("produk tidak ditemukan")
}
