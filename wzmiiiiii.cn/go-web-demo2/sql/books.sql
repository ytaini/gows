use go_test;
create table books(
    id int PRIMARY KEY auto_increment,
    title varchar(100) not null ,
    author varchar(100) not null ,
    price DOUBLE(11,2) not null ,
    sales int not null ,
    stock int not null ,
    img_path varchar(100) not null
)
