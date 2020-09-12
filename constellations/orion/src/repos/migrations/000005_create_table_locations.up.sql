CREATE TABLE locations
(
    id         int unsigned NOT NULL AUTO_INCREMENT,
    created_at datetime     NOT NULL,
    updated_at datetime     NOT NULL,
    deleted_at datetime,
    location_id     varchar(64)  NOT NULL UNIQUE,
    title      varchar(64) NOT NULL,
    street     varchar(255),
    city       varchar(64),
    state      varchar(64),
    zipcode    varchar(64),
    room       varchar(64),
    is_online  tinyint unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;