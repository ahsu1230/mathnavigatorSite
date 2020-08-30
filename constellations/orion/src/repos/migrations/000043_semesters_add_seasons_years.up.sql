ALTER TABLE semesters
    ADD COLUMN season VARCHAR(16) NOT NULL AFTER semester_id,
    ADD COLUMN year INT NOT NULL AFTER season;