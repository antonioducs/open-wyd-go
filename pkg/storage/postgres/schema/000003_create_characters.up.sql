CREATE TABLE characters (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL REFERENCES accounts (id) ON DELETE CASCADE,

    name VARCHAR(16) NOT NULL UNIQUE,
    level INT NOT NULL DEFAULT 0,
    experience BIGINT NOT NULL DEFAULT 0,
    gold INT NOT NULL DEFAULT 0,
    guild_id INT NOT NULL DEFAULT 0,

    pos_x INT NOT NULL DEFAULT 0,
    pos_y INT NOT NULL DEFAULT 0,

    slot INT NOT NULL DEFAULT 0,
    class_id INT NOT NULL DEFAULT 0,

    status JSONB NOT NULL DEFAULT '{}'::JSONB,

    equipment JSONB NOT NULL DEFAULT '[]'::JSONB,

    inventory JSONB NOT NULL DEFAULT '[]'::JSONB,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_characters_account_id ON characters (account_id);
