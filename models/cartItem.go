package models

type CartItem struct {
	ID       uint   `json:"id"`
	Product  string `json:"product"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
