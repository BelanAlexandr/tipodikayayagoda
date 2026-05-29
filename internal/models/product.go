package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImgURL      string  `json:"img_url"`
	Category_id int     `json:"category_id"`
	Price       float64 `json:"price"`
	Count       int     `json:"count"`
	Offer       bool    `json:"offer"`
}
