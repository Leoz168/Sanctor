-- UP
CREATE TABLE IF NOT EXISTS dm_groups (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS dm_group_users (
    group_id UUID NOT NULL REFERENCES dm_groups(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (group_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_dm_group_users_user_id ON dm_group_users(user_id);

CREATE TABLE IF NOT EXISTS dm_messages (
    id UUID PRIMARY KEY,
    group_id UUID NOT NULL REFERENCES dm_groups(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    message_time TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_dm_messages_group_id_message_time ON dm_messages(group_id, message_time DESC);

-- DOWN
DROP TABLE IF EXISTS dm_messages;
DROP TABLE IF EXISTS dm_group_users;
DROP TABLE IF EXISTS dm_groups;
