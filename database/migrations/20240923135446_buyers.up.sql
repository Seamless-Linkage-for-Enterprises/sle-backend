CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE buyers (
    "b_id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "b_first_name" VARCHAR(20) NOT NULL,
    "b_last_name" VARCHAR(20) NOT NULL,
    "b_phone" VARCHAR(15) NOT NULL UNIQUE,
    "b_password" VARCHAR(100) NOT NULL,
    "b_email" VARCHAR(40) NOT NULL,
    "b_image_url" VARCHAR(1000) NOT NULL,
    "b_address" VARCHAR(300) NOT NULL,
    "b_dob" DATE NOT NULL,
    "is_phone_verified" BOOLEAN DEFAULT FALSE NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column_buyers()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_buyers_updated_at
BEFORE UPDATE ON buyers
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column_buyers();