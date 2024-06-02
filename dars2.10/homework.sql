CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);


BEGIN;
INSERT INTO users (username, email, password)
VALUES ('ali', 'ali561l@example.com', 'pass ali');
INSERT INTO products (name, description, price, stock_quantity)
VALUES ('olma', 'qizil', 10000.00, 10);
COMMIT;

BEGIN;

UPDATE users
SET username = 'shox', email = 'shox123@example.com', password = '1238499'
WHERE id = 1;

COMMIT;

BEGIN;

UPDATE products
SET description = 'yashil', price = 9900.00, stock_quantity = 20
WHERE id = 1;

COMMIT;

BEGIN;

DELETE FROM users WHERE id = 1;

COMMIT;


BEGIN;

DELETE FROM products WHERE id = 1;

COMMIT;
