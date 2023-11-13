-- Add additional column to organizations
ALTER TABLE organizations
ADD COLUMN province_id text DEFAULT '',
ADD COLUMN city_id text DEFAULT '',
ADD COLUMN district_id text DEFAULT '',
ADD COLUMN village_id text DEFAULT '';

UPDATE organizations
SET pic_phone = phone_number;

ALTER TABLE organizations
DROP COLUMN phone_number;
