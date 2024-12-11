CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE bookmarks (
    "bookmark_id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "p_id" UUID NOT NULL,
    "b_id" UUID NOT NULL,
    CONSTRAINT fk_product
    FOREIGN KEY ("p_id") 
    REFERENCES products("p_id")
    ON DELETE CASCADE, -- Automatically delete related bookmarks
    CONSTRAINT fk_buyer
    FOREIGN KEY ("b_id") 
    REFERENCES buyers("b_id") -- buyer id
);