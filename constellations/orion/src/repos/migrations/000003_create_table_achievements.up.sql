CREATE TABLE achievements
(
    id           int unsigned NOT NULL AUTO_INCREMENT,
    created_at   datetime     NOT NULL,
    updated_at   datetime     NOT NULL,
    deleted_at   datetime,
    year         int unsigned NOT NULL,
    message      text         NOT NULL,
    position     int unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;
