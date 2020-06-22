CREATE TABLE ask_for_help
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  datetime     NOT NULL,
    updated_at  datetime     NOT NULL,
    deleted_at  datetime,
    title       text         NOT NULL,
    date        date         NOT NULL,
    timeString  varchar(64)  NOT NULL,
    subject     text         NOT NULL,
    location_id varchar(64)  NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (location_id) REFERENCES locations (location_id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;