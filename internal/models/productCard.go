package models

type ProductCard struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageURL    string  `json:"img_url"`
	CategoryID  int     `json:"category_id"`
	MinPrice    float64 `json:"min_price"`
	TotalCount  int     `json:"total_count"`
}
