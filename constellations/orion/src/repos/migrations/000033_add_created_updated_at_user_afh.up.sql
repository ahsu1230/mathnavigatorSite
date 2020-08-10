ALTER TABLE user_afh
    ADD COLUMN created_at datetime NOT NULL,
    ADD COLUMN updated_at datetime NOT NULL,
    ADD COLUMN deleted_at datetime;