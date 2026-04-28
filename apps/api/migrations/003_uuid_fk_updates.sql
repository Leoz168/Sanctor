-- UP
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Drop foreign keys that block type changes.
ALTER TABLE IF EXISTS users DROP CONSTRAINT IF EXISTS users_institution_id_fkey;
ALTER TABLE IF EXISTS groups DROP CONSTRAINT IF EXISTS groups_created_by_fkey;
ALTER TABLE IF EXISTS user_groups DROP CONSTRAINT IF EXISTS user_groups_user_id_fkey;
ALTER TABLE IF EXISTS user_groups DROP CONSTRAINT IF EXISTS user_groups_group_id_fkey;
ALTER TABLE IF EXISTS group_institutions DROP CONSTRAINT IF EXISTS group_institutions_group_id_fkey;
ALTER TABLE IF EXISTS group_institutions DROP CONSTRAINT IF EXISTS group_institutions_institution_id_fkey;
ALTER TABLE IF EXISTS posts DROP CONSTRAINT IF EXISTS posts_user_id_fkey;
ALTER TABLE IF EXISTS posts DROP CONSTRAINT IF EXISTS posts_created_by_fkey;
ALTER TABLE IF EXISTS posts DROP CONSTRAINT IF EXISTS posts_updated_by_fkey;
ALTER TABLE IF EXISTS post_groups DROP CONSTRAINT IF EXISTS post_groups_post_id_fkey;
ALTER TABLE IF EXISTS post_groups DROP CONSTRAINT IF EXISTS post_groups_group_id_fkey;
ALTER TABLE IF EXISTS post_institutions DROP CONSTRAINT IF EXISTS post_institutions_post_id_fkey;
ALTER TABLE IF EXISTS post_institutions DROP CONSTRAINT IF EXISTS post_institutions_institution_id_fkey;
ALTER TABLE IF EXISTS post_bookmarks DROP CONSTRAINT IF EXISTS post_bookmarks_user_id_fkey;
ALTER TABLE IF EXISTS post_bookmarks DROP CONSTRAINT IF EXISTS post_bookmarks_post_id_fkey;
ALTER TABLE IF EXISTS pictures DROP CONSTRAINT IF EXISTS pictures_post_id_fkey;
ALTER TABLE IF EXISTS comments DROP CONSTRAINT IF EXISTS comments_post_id_fkey;
ALTER TABLE IF EXISTS comments DROP CONSTRAINT IF EXISTS comments_created_by_user_id_fkey;
ALTER TABLE IF EXISTS dm_group_users DROP CONSTRAINT IF EXISTS dm_group_users_group_id_fkey;
ALTER TABLE IF EXISTS dm_group_users DROP CONSTRAINT IF EXISTS dm_group_users_user_id_fkey;
ALTER TABLE IF EXISTS dm_messages DROP CONSTRAINT IF EXISTS dm_messages_group_id_fkey;
ALTER TABLE IF EXISTS dm_messages DROP CONSTRAINT IF EXISTS dm_messages_user_id_fkey;

-- Convert primary/foreign key columns to uuid.
ALTER TABLE IF EXISTS users
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN id TYPE uuid USING id::uuid,
    ALTER COLUMN institution_id TYPE uuid USING institution_id::uuid;

ALTER TABLE IF EXISTS institutions
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN id TYPE uuid USING id::uuid;

ALTER TABLE IF EXISTS groups
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN id TYPE uuid USING id::uuid,
    ALTER COLUMN created_by TYPE uuid USING created_by::uuid;

ALTER TABLE IF EXISTS user_groups
    ALTER COLUMN user_id TYPE uuid USING user_id::uuid,
    ALTER COLUMN group_id TYPE uuid USING group_id::uuid;

ALTER TABLE IF EXISTS group_institutions
    ALTER COLUMN group_id TYPE uuid USING group_id::uuid,
    ALTER COLUMN institution_id TYPE uuid USING institution_id::uuid;

ALTER TABLE IF EXISTS posts
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN id TYPE uuid USING id::uuid,
    ALTER COLUMN user_id TYPE uuid USING user_id::uuid,
    ALTER COLUMN created_by TYPE uuid USING created_by::uuid,
    ALTER COLUMN updated_by TYPE uuid USING updated_by::uuid;

ALTER TABLE IF EXISTS post_groups
    ALTER COLUMN post_id TYPE uuid USING post_id::uuid,
    ALTER COLUMN group_id TYPE uuid USING group_id::uuid;

ALTER TABLE IF EXISTS post_institutions
    ALTER COLUMN post_id TYPE uuid USING post_id::uuid,
    ALTER COLUMN institution_id TYPE uuid USING institution_id::uuid;

ALTER TABLE IF EXISTS post_bookmarks
    ALTER COLUMN user_id TYPE uuid USING user_id::uuid,
    ALTER COLUMN post_id TYPE uuid USING post_id::uuid;

ALTER TABLE IF EXISTS pictures
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN id TYPE uuid USING id::uuid,
    ALTER COLUMN post_id TYPE uuid USING post_id::uuid;

ALTER TABLE IF EXISTS comments
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN id TYPE uuid USING id::uuid,
    ALTER COLUMN post_id TYPE uuid USING post_id::uuid,
    ALTER COLUMN created_by_user_id TYPE uuid USING created_by_user_id::uuid;

ALTER TABLE IF EXISTS dm_groups
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN id TYPE uuid USING id::uuid;

ALTER TABLE IF EXISTS dm_group_users
    ALTER COLUMN group_id TYPE uuid USING group_id::uuid,
    ALTER COLUMN user_id TYPE uuid USING user_id::uuid;

ALTER TABLE IF EXISTS dm_messages
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN id TYPE uuid USING id::uuid,
    ALTER COLUMN group_id TYPE uuid USING group_id::uuid,
    ALTER COLUMN user_id TYPE uuid USING user_id::uuid;

-- Recreate foreign keys using uuid columns.
ALTER TABLE IF EXISTS users
    ADD CONSTRAINT users_institution_id_fkey FOREIGN KEY (institution_id) REFERENCES institutions(id) ON DELETE SET NULL;

ALTER TABLE IF EXISTS groups
    ADD CONSTRAINT groups_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT;

ALTER TABLE IF EXISTS user_groups
    ADD CONSTRAINT user_groups_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    ADD CONSTRAINT user_groups_group_id_fkey FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS group_institutions
    ADD CONSTRAINT group_institutions_group_id_fkey FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    ADD CONSTRAINT group_institutions_institution_id_fkey FOREIGN KEY (institution_id) REFERENCES institutions(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS posts
    ADD CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    ADD CONSTRAINT posts_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    ADD CONSTRAINT posts_updated_by_fkey FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE IF EXISTS post_groups
    ADD CONSTRAINT post_groups_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    ADD CONSTRAINT post_groups_group_id_fkey FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS post_institutions
    ADD CONSTRAINT post_institutions_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    ADD CONSTRAINT post_institutions_institution_id_fkey FOREIGN KEY (institution_id) REFERENCES institutions(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS post_bookmarks
    ADD CONSTRAINT post_bookmarks_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    ADD CONSTRAINT post_bookmarks_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS pictures
    ADD CONSTRAINT pictures_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS comments
    ADD CONSTRAINT comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    ADD CONSTRAINT comments_created_by_user_id_fkey FOREIGN KEY (created_by_user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS dm_group_users
    ADD CONSTRAINT dm_group_users_group_id_fkey FOREIGN KEY (group_id) REFERENCES dm_groups(id) ON DELETE CASCADE,
    ADD CONSTRAINT dm_group_users_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS dm_messages
    ADD CONSTRAINT dm_messages_group_id_fkey FOREIGN KEY (group_id) REFERENCES dm_groups(id) ON DELETE CASCADE,
    ADD CONSTRAINT dm_messages_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

-- DOWN
-- No-op: UUID conversion is irreversible without data loss.
