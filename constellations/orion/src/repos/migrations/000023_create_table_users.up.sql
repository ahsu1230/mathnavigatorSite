CREATE TABLE users
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  datetime     NOT NULL,
    updated_at  datetime     NOT NULL,
    deleted_at  datetime,
    first_name  varchar(32)  NOT NULL,
    last_name   varchar(32)  NOT NULL,
    middle_name varchar(32),
    email       varchar(64)  NOT NULL,
    phone       varchar(24)  NOT NULL,
    is_guardian boolean      NOT NULL DEFAULT 0,
    account_id   int unsigned NOT NULL,
    notes       varchar(64),
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES accounts (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;