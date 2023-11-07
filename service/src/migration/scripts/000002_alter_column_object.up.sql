-- Add additional column to object
ALTER TABLE objects
ADD province text,
ADD city text,
ADD district text,
ADD village text,
ADD google_map text;
