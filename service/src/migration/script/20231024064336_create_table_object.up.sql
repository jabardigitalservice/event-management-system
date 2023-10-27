-- Create an ENUM type for status
CREATE TYPE status AS ENUM ('draft', 'unpublished', 'published');

-- Create the "object" table with UUID for id and status as ENUM
CREATE TABLE object (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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



