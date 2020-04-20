ALTER TABLE sessions
    ADD COLUMN published_at datetime AFTER deleted_at;