ALTER TABLE user_afh
    ADD COLUMN created_at datetime NOT NULL AFTER id,
    ADD COLUMN updated_at datetime NOT NULL AFTER created_at,
    ADD COLUMN deleted_at datetime AFTER updated_at;