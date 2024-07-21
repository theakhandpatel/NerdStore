CREATE EXTENSION IF NOT EXISTS citext;
CREATE TABLE IF NOT EXISTS resources (
    id bigserial PRIMARY KEY,  
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    link text NOT NULL,
    tags text[],
    version integer NOT NULL DEFAULT 1
);

