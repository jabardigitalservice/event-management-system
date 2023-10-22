-- Create the 'event' table
CREATE TABLE "event" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(100),
  "description" TEXT,
  "location_name" VARCHAR(100),
  "address" VARCHAR(100),
  "date" TIMESTAMP,
  "organization" INT,
  "user_id" INT,
  "createdAt" TIMESTAMP,
  "updatedAt" TIMESTAMP
);

-- Add foreign keys to the 'event' table
ALTER TABLE "event" ADD CONSTRAINT "fk_organization" FOREIGN KEY ("organization") REFERENCES "organization" ("id");