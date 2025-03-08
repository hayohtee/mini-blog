CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    username text UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT now(),
);
