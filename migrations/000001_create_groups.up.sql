CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE groups
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    parent_id UUID REFERENCES groups(id)
        ON DELETE CASCADE
);