CREATE TABLE IF NOT EXISTS Items (
    item_id SERIAL PRIMARY KEY,
    item_uuid VARCHAR(36) UNIQUE NOT NULL,
    item_name VARCHAR(30) UNIQUE NOT NULL CHECK (TRIM(item_name) <> '')
);