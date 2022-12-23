use go_test;
create table users(
    id varchar(36) primary key,
    username varchar(20) unique not null,
    passwd varchar(20) not null,
    email varchar(50)
);