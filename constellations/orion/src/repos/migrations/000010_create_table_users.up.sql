CREATE TABLE users
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  datetime     NOT NULL,
    updated_at  datetime     NOT NULL,
    deleted_at  datetime,
    account_id  int unsigned NOT NULL,
    first_name  varchar(32)  NOT NULL,
    middle_name varchar(32),
    last_name   varchar(32)  NOT NULL,
    email       varchar(64)  NOT NULL,
    phone       varchar(24)  NOT NULL,
    is_admin_created boolean NOT NULL DEFAULT 0,
    is_guardian boolean      NOT NULL DEFAULT 0,
    school      varchar(128),
    graduation_year int unsigned,
    notes       varchar(64),
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES accounts (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;