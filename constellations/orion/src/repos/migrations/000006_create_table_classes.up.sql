CREATE TABLE classes
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  datetime     NOT NULL,
    updated_at  datetime     NOT NULL,
    deleted_at  datetime,
    published_at datetime,
    program_id  varchar(64)  NOT NULL,
    semester_id varchar(64)  NOT NULL,
    class_key   varchar(64),
    class_id    varchar(128) NOT NULL UNIQUE,
    location_id varchar(64)  NOT NULL,
    times_str   varchar(128)  NOT NULL,
    google_class_code  varchar(16),
    full_state  tinyint unsigned NOT NULL DEFAULT 0,
    price_per_session int unsigned,
    price_lump_sum int unsigned,
    payment_notes text,
    PRIMARY KEY (id),
    FOREIGN KEY (program_id) REFERENCES programs (program_id),
    FOREIGN KEY (semester_id) REFERENCES semesters (semester_id),
    FOREIGN KEY (location_id) REFERENCES locations (location_id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;
