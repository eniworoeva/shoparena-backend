package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	SellerId  uint `json:"seller_id"`
	Seller    Seller
	BuyerId   uint `json:"buyer_id"`
	Buyer     Buyer
	ProductId uint `json:"product_id"`
	Product   Product
}

type OrderProducts struct {
	CategoryName string
	Title        string
	Description  string
	Price        uint
	Images       []Image
	Quantity     uint
}
