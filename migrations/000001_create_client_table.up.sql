CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS client (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    birthday DATE NOT NULL,
    gender TEXT NOT NULL,
    registration_date TIMESTAMP NOT NULL,
    address_id UUID
);
