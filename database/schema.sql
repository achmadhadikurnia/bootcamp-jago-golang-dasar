-- 1. Tabel Categories
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

-- 2. Tabel Products
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    category_id INT REFERENCES categories(id) ON DELETE SET NULL
);

-- 3. Tabel Transactions
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    total_amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4. Tabel Transaction Details
CREATE TABLE IF NOT EXISTS transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id),
    quantity INT NOT NULL,
    subtotal INT NOT NULL
);

-- ================================================
-- Seed Data
-- ================================================

-- Insert Categories
INSERT INTO categories (name, description) VALUES
('Makanan', 'Produk makanan siap saji'),
('Minuman', 'Produk minuman'),
('Snack', 'Makanan ringan');

-- Insert Products
INSERT INTO products (name, price, stock, category_id) VALUES
('Indomie Goreng', 3500, 100, 1),
('Indomie Kuah', 3000, 80, 1),
('Indomie Rendang', 4000, 50, 1),
('Teh Botol Sosro', 5000, 60, 2),
('Aqua 600ml', 4000, 100, 2),
('Coca Cola', 7000, 40, 2),
('Chitato', 12000, 30, 3),
('Taro', 8000, 25, 3),
('Oreo', 10000, 35, 3);
