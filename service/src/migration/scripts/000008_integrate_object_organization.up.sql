ALTER TABLE objects
ADD COLUMN organization_id UUID DEFAULT 'adac8742-73a4-4d03-8ab0-60945c3ffda0',
ADD CONSTRAINT fk_organization
    FOREIGN KEY (organization_id)
    REFERENCES organizations(id);
