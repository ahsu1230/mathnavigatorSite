CREATE TABLE user_afhs
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  datetime     NOT NULL,
    updated_at  datetime     NOT NULL,
    deleted_at  datetime,
    user_id     int unsigned NOT NULL,
    afh_id      int unsigned NOT NULL,
    account_id  int unsigned NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (afh_id) REFERENCES ask_for_help (id),
    FOREIGN KEY (account_id) REFERENCES accounts (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;