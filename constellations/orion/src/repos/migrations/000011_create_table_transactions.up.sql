CREATE TABLE transactions 
(
    id            int unsigned not null auto_increment,
    created_at    datetime     not null,
    updated_at    datetime     not null,
    deleted_at    datetime,
    account_id    int unsigned not null,
    type          varchar(32)  not null,
    amount        int          not null,
    notes text,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES accounts (id)
) AUTO_INCREMENT = 1
DEFAULT CHARSET = UTF8MB4;