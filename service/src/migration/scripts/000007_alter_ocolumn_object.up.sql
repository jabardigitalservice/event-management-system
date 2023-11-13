-- Add additional column to objects
ALTER TABLE objects
ADD COLUMN province_id text DEFAULT '',
ADD COLUMN city_id text DEFAULT '',
ADD COLUMN district_id text DEFAULT '',
ADD COLUMN village_id text DEFAULT '';
