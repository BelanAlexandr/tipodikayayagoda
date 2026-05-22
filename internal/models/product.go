package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Count       int     `json:"count"`
	SellerID    int     `json:"seller_id"`
	ImageURL    string  `json:"img_url"`
}
