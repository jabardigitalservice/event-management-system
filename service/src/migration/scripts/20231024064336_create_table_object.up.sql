-- Create Enum for status
CREATE TYPE status AS ENUM ('draft', 'unpublished', 'published');

-- Create table "object" with the structure provided
CREATE TABLE object (
    id serial PRIMARY KEY,
    name varchar(100),
    address text,
    description text,
    banner text[],
    logo text,
    social_media jsonb, 
    organizer text,
    status status,
    created_at timestamp,
    updated_at timestamp
);



