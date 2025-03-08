CREATE TABLE IF NOT EXISTS blogs (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title text NOT NULL,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    content text NOT NULL,
    version integer NOT NULL DEFAULT 1
);