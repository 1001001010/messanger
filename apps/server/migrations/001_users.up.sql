CREATE EXTENSION IF NOT EXISTS pgcrypto;


CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    username VARCHAR(32) UNIQUE,

    email VARCHAR(255) NOT NULL UNIQUE,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,

    password_hash TEXT NOT NULL,

    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64),

    bio VARCHAR(140),

    file_id TEXT,

    last_seen TIMESTAMPTZ,
    is_online BOOLEAN NOT NULL DEFAULT FALSE,

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);