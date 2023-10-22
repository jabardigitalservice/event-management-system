-- Create the 'organization' table
CREATE TABLE "organization" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(100),
  "email" VARCHAR(100),
  "address" VARCHAR(100),
  "phone_number" VARCHAR(30),
  "logo" VARCHAR,
  "createdAt" TIMESTAMP,
  "updatedAt" TIMESTAMP
);
