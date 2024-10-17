CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
    "p_id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "p_name" VARCHAR(100) NOT NULL,
    "p_category" VARCHAR(100) NOT NULL,
    "p_brand" VARCHAR(100) NOT NULL,
    "p_status" BOOLEAN DEFAULT TRUE,
    "p_description" VARCHAR(500) NOT NULL,
    "p_quantity" INT NOT NULL,
    "p_image" VARCHAR(500) NOT NULL,
    "p_price" INT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "s_id" UUID NOT NULL,
    CONSTRAINT fk_seller
    FOREIGN KEY ("s_id") 
    REFERENCES sellers("s_id")
);


CREATE OR REPLACE FUNCTION update_updated_at_column_product()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_product_updated_at
BEFORE UPDATE ON sellers
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column_product();