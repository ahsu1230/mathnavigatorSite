ALTER TABLE programs
    ADD COLUMN published_at datetime AFTER deleted_at;