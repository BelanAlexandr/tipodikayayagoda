CREATE EXTENSION IF NOT EXISTS pg_trgm;


CREATE TABLE roledictionary (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    pass VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    secondname VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    role INTEGER NOT NULL REFERENCES roledictionary(id) ON DELETE RESTRICT
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL, 
    img_url VARCHAR(255),
    offer BOOLEAN NOT NULL DEFAULT FALSE,
    min_price NUMERIC(10, 2) NOT NULL,
    category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
    name_tsvector TSVECTOR GENERATED ALWAYS AS (to_tsvector('russian', name)) STORED
);

CREATE TABLE product_offers (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    seller_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    price NUMERIC(10, 2) NOT NULL,
    count INTEGER NOT NULL,
    CONSTRAINT unique_product_seller UNIQUE(product_id, seller_id)
);

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);



CREATE INDEX idx_products_tsvector ON products USING gin(name_tsvector);

CREATE INDEX IF NOT EXISTS idx_products_name_trgm 
ON products USING gin (name gin_trgm_ops);


CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_id ON products(id);
CREATE INDEX idx_product_offers_seller_id ON product_offers(seller_id);
CREATE INDEX idx_product_offers_product_id ON product_offers(product_id);
CREATE INDEX idx_products_price_id ON products(min_price, id);

CREATE INDEX idx_notifications_user_unread ON notifications (user_id) WHERE is_read = FALSE;


INSERT INTO roledictionary (id, role_name) VALUES 
(16, 'admin'),
(27, 'seller'),
(38, 'client')
ON CONFLICT (id) DO NOTHING;