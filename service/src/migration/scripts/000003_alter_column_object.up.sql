-- Add additional column to object
ALTER TABLE objects
ADD organizer_email text DEFAULT '',
ADD organizer_phone text DEFAULT '';
