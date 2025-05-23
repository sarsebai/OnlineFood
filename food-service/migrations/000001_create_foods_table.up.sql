CREATE TABLE IF NOT EXISTS foods (
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    maker_id INT,
    category_id INT,
    price DECIMAL(10, 2) NOT NULL
); 