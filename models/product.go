package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ShopName        string  `json:"shop_name"`
	ProductName     string  `json:"product_name"`
	ProductPrice    float32 `json:"product_price"`
	ProductCategory string  `json:"product_category"`
	ProductImage    string  `json:"product_image"`
	ProductDetails  string  `json:"product_details"`
	Rating          int     `json:"rating"`
	Quantity        int     `json:"quantity"`
}