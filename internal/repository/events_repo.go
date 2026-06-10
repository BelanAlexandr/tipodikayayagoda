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
			sa.product_id,
			p.name,
			COUNT(CASE WHEN sa.event_type = 'view' THEN 1 END) AS total_views,
			COUNT(CASE WHEN sa.event_type = 'purchase' THEN 1 END) AS total_purchases,
			COALESCE(SUM(CASE WHEN sa.event_type = 'purchase' THEN sa.quantity END), 0) AS total_units_sold,
			CASE 
				WHEN COUNT(CASE WHEN sa.event_type = 'purchase' THEN 1 END) > 0 
				THEN ROUND(COALESCE(SUM(CASE WHEN sa.event_type = 'purchase' THEN sa.quantity END), 0)::numeric / 
						   COUNT(CASE WHEN sa.event_type = 'purchase' THEN 1 END)::numeric, 2)
				ELSE 0 
			END AS avg_units_per_order
		FROM site_analytics sa
		INNER JOIN products p ON sa.product_id = p.id
		WHERE sa.seller_id = $1 AND sa.product_id IS NOT NULL
		GROUP BY sa.product_id, p.name;
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
        COUNT(DISTINCT user_id) FILTER (WHERE created_at >= NOW() - INTERVAL '1 day') AS active_today,
        COUNT(DISTINCT user_id) FILTER (WHERE created_at >= NOW() - INTERVAL '30 days') AS active_this_month
    FROM site_analytics
    WHERE user_id IS NOT NULL;
`
	err := db.QueryRow(activityQuery).Scan(&stats.ActiveToday, &stats.ActiveThisMonth)
	if err != nil {
		return stats, err
	}

	conversionQuery := `
    WITH user_flags AS (
        SELECT 
            user_id,
            BOOL_OR(event_type = 'register') AS has_registered,
            BOOL_OR(event_type = 'purchase') AS has_purchased
        FROM site_analytics
        WHERE user_id IS NOT NULL
        GROUP BY user_id
    )
    SELECT 
        COUNT(*) FILTER (WHERE has_registered) AS registered_users,
        COUNT(*) FILTER (WHERE has_registered AND has_purchased) AS registered_buyers,
        ROUND(
            (COUNT(*) FILTER (WHERE has_registered AND has_purchased))::numeric / 
            NULLIF(COUNT(*) FILTER (WHERE has_registered), 0)::numeric * 100, 2
        ) AS conversion_rate
    FROM user_flags;
`
	err = db.QueryRow(conversionQuery).Scan(&stats.TotalRegistered, &stats.RegisteredAndBought, &stats.ConversionRate)
	return stats, err
}
