CREATE INDEX IF NOT EXISTS resources_title_idx ON resources USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS resources_tags_idx ON resources USING GIN (tags);