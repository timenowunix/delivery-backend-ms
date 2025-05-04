CREATE TABLE IF NOT EXISTS deliveries (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    courier_id INTEGER,
    status TEXT NOT NULL CHECK (status IN ('pending', 'assigned', 'in_transit', 'delivered', 'failed')),
    priority TEXT NOT NULL CHECK (priority IN ('normal', 'express')),
    delivery_address TEXT NOT NULL,
    estimated_delivery_time TIMESTAMP NOT NULL,
    delivered_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);