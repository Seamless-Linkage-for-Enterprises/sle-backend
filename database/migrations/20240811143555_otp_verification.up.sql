CREATE TABLE otp_verification (
    "id" BIGSERIAL NOT NULL PRIMARY KEY,
    "s_id" UUID NOT NULL,
    "email" VARCHAR(100) NOT NULL,
    "otp" VARCHAR(6) NOT NULL,
    "expires_at" INT NOT NULL
)