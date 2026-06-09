package models

const (
	EventRegister = "register"
	EventLogin    = "login"
	EventLogout   = "logout"
	EventView     = "view"
	EventPurchase = "purchase"
)

type AnalyticsEvent struct {
	UserID    *int   `json:"user_id"`
	ProductID *int   `json:"product_id"`
	SellerID  *int   `json:"seller_id"`
	EventType string `json:"event_type"`
	Quantity  int    `json:"quantity"`
}

type SellerStats struct {
	ProductID      *int    `json:"product_id"`
	ProductName    string  `json:"product_name"`
	TotalViews     int     `json:"total_views"`
	TotalPurchases int     `json:"total_purchases"`
	TotalUnitsSold int     `json:"total_units_sold"`
	AvgUnitsOrder  float64 `json:"avg_units_per_order"`
}
type AdminStats struct {
	ActiveToday         int     `json:"active_today"`
	ActiveThisMonth     int     `json:"active_this_month"`
	TotalRegistered     int     `json:"total_registered"`
	RegisteredAndBought int     `json:"registered_and_bought"`
	ConversionRate      float64 `json:"conversion_rate"`
}
