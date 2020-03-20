CREATE TABLE classes
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  datetime     NOT NULL,
    updated_at  datetime     NOT NULL,
    deleted_at  datetime,
    program_id  varchar(64)  NOT NULL,
    class_key   varchar(64),
    class_id    varchar(192) NOT NULL UNIQUE,
    semester_id varchar(64)  NOT NULL,
    location_id varchar(64)  NOT NULL,
    times       varchar(64)  NOT NULL,
    start_date  date         NOT NULL,
    end_date    date         NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (program_id) REFERENCES programs (program_id),
    FOREIGN KEY (semester_id) REFERENCES semesters (semester_id),
    FOREIGN KEY (location_id) REFERENCES locations (loc_id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;
