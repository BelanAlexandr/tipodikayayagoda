package repository

import (
	"tipodikayayagoda/internal/models"
)

func TrackEvent(event models.AnalyticsEvent) error {
	query := `
		INSERT INTO site_analytics (user_id, product_id, seller_id, event_type, quantity)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := db.Exec(query, event.UserID, event.ProductID, event.SellerID, event.EventType, event.Quantity)
	return err
}

func GetSellerAnalytics(sellerID int) ([]models.SellerStats, error) {
	query := `
		SELECT 
			p.id, p.name,
			COUNT(CASE WHEN sa.event_type = 'view' THEN 1 END),
			COUNT(CASE WHEN sa.event_type = 'purchase' THEN 1 END),
			COALESCE(SUM(CASE WHEN sa.event_type = 'purchase' THEN sa.quantity END), 0),
			CASE 
				WHEN COUNT(CASE WHEN sa.event_type = 'purchase' THEN 1 END) > 0 
				THEN ROUND(COALESCE(SUM(CASE WHEN sa.event_type = 'purchase' THEN sa.quantity END), 0)::numeric / 
						   COUNT(CASE WHEN sa.event_type = 'purchase' THEN 1 END)::numeric, 2)
				ELSE 0 
			END
		FROM products p
		LEFT JOIN site_analytics sa ON p.id = sa.product_id
		WHERE p.seller_id = $1
		GROUP BY p.id, p.name;
	`

	rows, err := db.Query(query, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []models.SellerStats
	for rows.Next() {
		var s models.SellerStats
		if err := rows.Scan(&s.ProductID, &s.ProductName, &s.TotalViews, &s.TotalPurchases, &s.TotalUnitsSold, &s.AvgUnitsOrder); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}
	return stats, nil
}
func GetAdminAnalytics() (models.AdminStats, error) {
	var stats models.AdminStats

	activityQuery := `
		SELECT
			COUNT(DISTINCT CASE WHEN created_at >= NOW() - INTERVAL '1 day' THEN user_id END),
			COUNT(DISTINCT CASE WHEN created_at >= NOW() - INTERVAL '30 days' THEN user_id END)
		FROM site_analytics
		WHERE user_id IS NOT NULL;
	`
	err := db.QueryRow(activityQuery).Scan(&stats.ActiveToday, &stats.ActiveThisMonth)
	if err != nil {
		return stats, err
	}

	conversionQuery := `
		SELECT 
			COUNT(DISTINCT CASE WHEN event_type = 'register' THEN user_id END),
			COUNT(DISTINCT CASE WHEN event_type = 'purchase' AND user_id IN (
				SELECT user_id FROM site_analytics WHERE event_type = 'register'
			) THEN user_id END),
			ROUND(
				(COUNT(DISTINCT CASE WHEN event_type = 'purchase' AND user_id IN (
					SELECT user_id FROM site_analytics WHERE event_type = 'register'
				) THEN user_id END)::numeric / 
				NULLIF(COUNT(DISTINCT CASE WHEN event_type = 'register' THEN user_id END), 0)::numeric) * 100, 2
			)
		FROM site_analytics;
	`
	err = db.QueryRow(conversionQuery).Scan(&stats.TotalRegistered, &stats.RegisteredAndBought, &stats.ConversionRate)
	return stats, err
}
