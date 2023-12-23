CREATE TABLE IF NOT EXISTS "orders" (
  id VARCHAR(255),
  price DECIMAL(8,2),
  tax DECIMAL(8,2),
  final_price DECIMAL(8,2)
);