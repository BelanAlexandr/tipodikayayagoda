CREATE TABLE roledictionary (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL UNIQUE
);


INSERT INTO roledictionary (id, role_name) VALUES 
(38, 'client'),
(27, 'seller'),
(16, 'admin');


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
    category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE RESTRICT
);


CREATE TABLE product_offers (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    seller_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    price REAL NOT NULL,
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


CREATE INDEX idx_notifications_user_unread ON notifications (user_id, is_read);