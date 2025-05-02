CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    parcel_id BIGINT NOT NULL,
    delivery_address TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

