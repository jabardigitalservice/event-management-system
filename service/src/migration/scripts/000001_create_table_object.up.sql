-- Create an ENUM type for status
CREATE TYPE status AS ENUM ('draft', 'unpublished', 'published');
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Create the "object" table with UUID for id and status as ENUM
CREATE TABLE objects (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(100),
    province text,
    city text,
    district text,
    village text,
    address text,
    google_map text,
    description text,
    banner text[],
    logo text,
    social_media jsonb, 
    organizer text,
    status status,
    created_at timestamp,
    updated_at timestamp
);
