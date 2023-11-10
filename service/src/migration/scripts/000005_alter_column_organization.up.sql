-- Add additional column to organizations
ALTER TABLE organizations
ADD COLUMN province text DEFAULT '',
ADD COLUMN city text DEFAULT '',
ADD COLUMN district text DEFAULT '',
ADD COLUMN village text DEFAULT '',
ADD COLUMN google_map text DEFAULT '',
ADD COLUMN pic_name text DEFAULT '',
ADD COLUMN pic_position text DEFAULT '',
ADD COLUMN pic_phone text DEFAULT '';
