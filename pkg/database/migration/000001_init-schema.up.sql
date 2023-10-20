create table if not exists users(
    id serial primary key,
    first_name varchar(200),
    last_name varchar(200),
    patronymic varchar(200),
    age int,
    gender varchar(10),
    nationality varchar(10),
    nationality_probability double precision
);