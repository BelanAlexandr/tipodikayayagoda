CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    pass VARCHAR(255) NOT NULL,
    role INTEGER NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price REAL NOT NULL,
    count INTEGER NOT NULL,
    seller_id INTEGER NOT NULL,
    img_url VARCHAR(255)
);

CREATE TABLE roledictionary (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL
);