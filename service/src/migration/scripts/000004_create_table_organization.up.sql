CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(100),
    email varchar(255),
    address text,
    phone_number text,
    description text,
    logo text,
    created_at timestamp,
    updated_at timestamp
);
