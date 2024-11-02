create table if not exists users(
id serial primary key,
username varchar(50) not null,
email varchar(100) not null unique,
password_hash varchar(255) not null,
created_at timestamp default current_timestamp
);