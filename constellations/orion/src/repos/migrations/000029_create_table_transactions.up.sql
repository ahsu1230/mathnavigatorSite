CREATE TABLE transactions 
(
    id           int unsigned not null auto_increment,
    created_at   datetime     not null,
    amount       int          not null,
    payment_type varchar(16)  not null,
    primary key (id),
    foreign key (account_id) references accounts (id)
)