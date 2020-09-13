CREATE TABLE user_classes
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  datetime     NOT NULL,
    updated_at  datetime     NOT NULL,
    deleted_at  datetime,
    class_id    varchar(128) NOT NULL,
    user_id     int unsigned NOT NULL,
    account_id  int unsigned NOT NULL,
    state       int unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    FOREIGN KEY (class_id) REFERENCES classes (class_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (account_id) REFERENCES accounts (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;