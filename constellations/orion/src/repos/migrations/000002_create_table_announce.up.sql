CREATE TABLE announcements
(
    id         int unsigned NOT NULL AUTO_INCREMENT,
    created_at datetime     NOT NULL,
    updated_at datetime     NOT NULL,
    deleted_at datetime,
    posted_at  datetime     NOT NULL,
    author     varchar(255) NOT NULL,
    message    text         NOT NULL,
    on_home_page  boolean,
    PRIMARY KEY (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;