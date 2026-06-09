CREATE TABLE IF NOT EXISTS site_analytics (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT,
    product_id BIGINT,
    seller_id BIGINT,
    event_type VARCHAR(50) NOT NULL,
    quantity INT DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_analytics_seller_type ON site_analytics(seller_id, event_type);
CREATE INDEX IF NOT EXISTS idx_analytics_user_type ON site_analytics(user_id, event_type);