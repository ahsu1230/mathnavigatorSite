CREATE TABLE users
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  bigint(20)   NOT NULL,
    updated_at  bigint(20)   NOT NULL,
    deleted_at  datetime,
    first_name  varchar(32)  NOT NULL,
    last_name   varchar(32)  NOT NULL,
    middle_name varchar(32),
    email       varchar(64)  NOT NULL,
    phone       varchar(24)  NOT NULL,
    isGuardian  boolean      NOT NULL,
    guardianId  int unsigned NOT NULL,
    PRIMARY KEY (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;
