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
    primary key (id),
    foreign key (account_id) references accounts (id)
)