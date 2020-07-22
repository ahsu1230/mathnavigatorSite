CREATE TABLE userclass
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  datetime     NOT NULL,
    updated_at  datetime     NOT NULL,
    deleted_at  datetime,
    user_id     int unsigned NOT NULL,
    class_id    varchar(192) NOT NULL,
    account_id  int unsigned NOT NULL,
    state       int unsigned NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES accounts (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (class_id) REFERENCES classes (class_id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;