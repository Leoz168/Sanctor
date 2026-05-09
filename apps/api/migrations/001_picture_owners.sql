ALTER TABLE pictures ADD COLUMN IF NOT EXISTS owner_type varchar(40);
ALTER TABLE pictures ADD COLUMN IF NOT EXISTS owner_id uuid;

UPDATE pictures
SET owner_type = 'post'
WHERE owner_type IS NULL;

UPDATE pictures
SET owner_id = post_id
WHERE owner_id IS NULL AND post_id IS NOT NULL;

ALTER TABLE pictures ALTER COLUMN owner_type SET NOT NULL;
ALTER TABLE pictures ALTER COLUMN owner_id SET NOT NULL;
ALTER TABLE pictures ALTER COLUMN post_id DROP NOT NULL;

CREATE INDEX IF NOT EXISTS idx_pictures_owner ON pictures (owner_type, owner_id);
