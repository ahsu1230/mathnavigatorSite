ALTER TABLE classes
    ADD COLUMN published_at datetime AFTER deleted_at;