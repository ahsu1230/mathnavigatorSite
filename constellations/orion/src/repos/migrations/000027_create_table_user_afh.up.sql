CREATE TABLE user_afh
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    user_id     int unsigned NOT NULL,
    afh_id      int unsigned NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (afh_id) REFERENCES ask_for_help (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;