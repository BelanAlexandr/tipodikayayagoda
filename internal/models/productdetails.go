package models

type OfferDetail struct {
	SellerID   int     `json:"seller_id"`
	SellerName string  `json:"seller_name"`
	Price      float64 `json:"price"`
	Count      int     `json:"count"`
}
type ProductDetails struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	ImageURL    string        `json:"img_url"`
	CategoryID  int           `json:"category_id"`
	Offers      []OfferDetail `json:"offers"`
}
