ALTER TABLE semesters
    ADD COLUMN published_at datetime AFTER deleted_at;