-- Revert sellers table
ALTER TABLE sellers
ALTER COLUMN s_image_url TYPE VARCHAR(100);

-- Revert buyers table
ALTER TABLE buyers
ALTER COLUMN b_image_url TYPE VARCHAR(100);

-- Revert products table
ALTER TABLE products
ALTER COLUMN p_image TYPE VARCHAR(100);
