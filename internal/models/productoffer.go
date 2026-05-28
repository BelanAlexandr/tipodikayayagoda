package models

type ProductOffer struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	SellerID  int     `json:"seller_id"`
	Price     float64 `json:"price"`
	Count     int     `json:"count"`
}
