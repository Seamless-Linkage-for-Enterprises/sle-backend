DROP TRIGGER IF EXISTS update_product_updated_at ON products;

DROP FUNCTION IF EXISTS update_updated_at_column_product;

DROP EXTENSION IF EXISTS uuid;

DROP TABLE IF EXISTS products;
