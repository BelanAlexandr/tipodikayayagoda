
DROP INDEX IF EXISTS idx_notifications_user_unread;
DROP INDEX IF EXISTS idx_products_price_id;
DROP INDEX IF EXISTS idx_product_offers_product_id;
DROP INDEX IF EXISTS idx_product_offers_seller_id;
DROP INDEX IF EXISTS idx_products_id;
DROP INDEX IF EXISTS idx_products_category_id;
DROP INDEX IF EXISTS idx_products_name_trgm;
DROP INDEX IF EXISTS idx_products_tsvector;

DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS product_offers;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS roledictionary;


DROP EXTENSION IF EXISTS pg_trgm;