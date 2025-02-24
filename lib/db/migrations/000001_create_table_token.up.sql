create table if not exists token (
    id serial primary key,
    name varchar(255) not null,
    symbol varchar(255) not null,
    address varchar(255) unique not null
);