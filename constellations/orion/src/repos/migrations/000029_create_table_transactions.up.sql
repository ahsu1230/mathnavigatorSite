CREATE TABLE transactions 
(
    id            int unsigned not null auto_increment,
    created_at    datetime     not null,
    updated_at    datetime     not null,
    deleted_at    datetime,
    amount        int          not null,
    payment_type  varchar(16)  not null,
    payment_notes text,
    account_id    int unsigned not null,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES accounts (id)
) AUTO_INCREMENT = 1
DEFAULT CHARSET = UTF8MB4;