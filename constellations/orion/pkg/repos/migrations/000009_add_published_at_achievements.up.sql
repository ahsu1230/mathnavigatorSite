ALTER TABLE achievements
    ADD COLUMN published_at datetime AFTER deleted_at;