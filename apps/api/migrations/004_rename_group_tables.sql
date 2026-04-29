-- UP
-- Rename group tables to community tables.
ALTER TABLE IF EXISTS groups RENAME TO communities;
ALTER TABLE IF EXISTS user_groups RENAME TO user_communities;
ALTER TABLE IF EXISTS group_institutions RENAME TO community_institutions;
ALTER TABLE IF EXISTS post_groups RENAME TO post_communities;

-- Rename group_id columns to community_id.
ALTER TABLE IF EXISTS user_communities RENAME COLUMN group_id TO community_id;
ALTER TABLE IF EXISTS community_institutions RENAME COLUMN group_id TO community_id;
ALTER TABLE IF EXISTS post_communities RENAME COLUMN group_id TO community_id;

-- Drop legacy constraint names and recreate with community naming.
ALTER TABLE IF EXISTS communities DROP CONSTRAINT IF EXISTS groups_created_by_fkey;
ALTER TABLE IF EXISTS user_communities DROP CONSTRAINT IF EXISTS user_groups_user_id_fkey;
ALTER TABLE IF EXISTS user_communities DROP CONSTRAINT IF EXISTS user_groups_group_id_fkey;
ALTER TABLE IF EXISTS community_institutions DROP CONSTRAINT IF EXISTS group_institutions_group_id_fkey;
ALTER TABLE IF EXISTS community_institutions DROP CONSTRAINT IF EXISTS group_institutions_institution_id_fkey;
ALTER TABLE IF EXISTS post_communities DROP CONSTRAINT IF EXISTS post_groups_post_id_fkey;
ALTER TABLE IF EXISTS post_communities DROP CONSTRAINT IF EXISTS post_groups_group_id_fkey;

ALTER TABLE IF EXISTS communities
    ADD CONSTRAINT communities_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT;

ALTER TABLE IF EXISTS user_communities
    ADD CONSTRAINT user_communities_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    ADD CONSTRAINT user_communities_community_id_fkey FOREIGN KEY (community_id) REFERENCES communities(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS community_institutions
    ADD CONSTRAINT community_institutions_community_id_fkey FOREIGN KEY (community_id) REFERENCES communities(id) ON DELETE CASCADE,
    ADD CONSTRAINT community_institutions_institution_id_fkey FOREIGN KEY (institution_id) REFERENCES institutions(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS post_communities
    ADD CONSTRAINT post_communities_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    ADD CONSTRAINT post_communities_community_id_fkey FOREIGN KEY (community_id) REFERENCES communities(id) ON DELETE CASCADE;

-- DOWN
-- Revert community tables back to group tables.
ALTER TABLE IF EXISTS post_communities DROP CONSTRAINT IF EXISTS post_communities_post_id_fkey;
ALTER TABLE IF EXISTS post_communities DROP CONSTRAINT IF EXISTS post_communities_community_id_fkey;
ALTER TABLE IF EXISTS community_institutions DROP CONSTRAINT IF EXISTS community_institutions_community_id_fkey;
ALTER TABLE IF EXISTS community_institutions DROP CONSTRAINT IF EXISTS community_institutions_institution_id_fkey;
ALTER TABLE IF EXISTS user_communities DROP CONSTRAINT IF EXISTS user_communities_user_id_fkey;
ALTER TABLE IF EXISTS user_communities DROP CONSTRAINT IF EXISTS user_communities_community_id_fkey;
ALTER TABLE IF EXISTS communities DROP CONSTRAINT IF EXISTS communities_created_by_fkey;

ALTER TABLE IF EXISTS user_communities RENAME COLUMN community_id TO group_id;
ALTER TABLE IF EXISTS community_institutions RENAME COLUMN community_id TO group_id;
ALTER TABLE IF EXISTS post_communities RENAME COLUMN community_id TO group_id;

ALTER TABLE IF EXISTS communities RENAME TO groups;
ALTER TABLE IF EXISTS user_communities RENAME TO user_groups;
ALTER TABLE IF EXISTS community_institutions RENAME TO group_institutions;
ALTER TABLE IF EXISTS post_communities RENAME TO post_groups;

ALTER TABLE IF EXISTS groups
    ADD CONSTRAINT groups_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT;

ALTER TABLE IF EXISTS user_groups
    ADD CONSTRAINT user_groups_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    ADD CONSTRAINT user_groups_group_id_fkey FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS group_institutions
    ADD CONSTRAINT group_institutions_group_id_fkey FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    ADD CONSTRAINT group_institutions_institution_id_fkey FOREIGN KEY (institution_id) REFERENCES institutions(id) ON DELETE CASCADE;

ALTER TABLE IF EXISTS post_groups
    ADD CONSTRAINT post_groups_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    ADD CONSTRAINT post_groups_group_id_fkey FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE;
