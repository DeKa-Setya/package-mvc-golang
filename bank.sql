create table accounts(
    id int(10) not null auto_increment primary key,
    id_account varchar(20) not null,
    name varchar(100) not null,
    password varchar(255) not null,
    account_number int(10) not null,
    saldo int(11) default 0
)

create table auth(
    id int(10) not null auto_increment primary key,
    name varchar(100) not null,
    password varchar(255) not null
)

create table transactions(
    id int(10) not null auto_increment primary key,
    transaction_type int(2) not null,
    transaction_description text null,
    sender int(10) not null,
    recipient int(10) not null,
    amount int(11) not null,
    time_stamp timestamp default current_timestamp
)