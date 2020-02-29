CREATE TABLE semesters
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  bigint(20)   NOT NULL,
    updated_at  bigint(20)   NOT NULL,
    deleted_at  datetime,
    semester_id varchar(64)  NOT NULL UNIQUE,
    title       varchar(64)  NOT NULL,
    PRIMARY KEY (id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;
