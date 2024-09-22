CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE sellers (
    "s_id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "s_first_name" VARCHAR(20) NOT NULL,
    "s_last_name" VARCHAR(20) NOT NULL,
    "s_email" VARCHAR(40) NOT NULL UNIQUE,
    "s_password" VARCHAR(100) NOT NULL,
    "s_image_url" VARCHAR(100) NOT NULL,
    "s_address" VARCHAR(100) NOT NULL,
    "s_phone" VARCHAR(15) NOT NULL,
    "s_pan_card" VARCHAR(10) NOT NULL,
    "s_dob" DATE NOT NULL,
    "s_company_name" VARCHAR(20) ,
    "s_description" VARCHAR(256) NOT NULL,
    "is_verified" BOOLEAN DEFAULT TRUE NOT NULL,
    "is_email_verified" BOOLEAN DEFAULT TRUE NOT NULL,
    "s_gst_number" VARCHAR(15) ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_sellers_updated_at
BEFORE UPDATE ON sellers
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();