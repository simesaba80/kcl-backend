create table if not exists users(
id serial primary key,
uid varchar(100) not null unique,
name varchar(20) not null,
sex varchar(5) not null,
email varchar(100) not null unique,
height int not null,
age int not null,
job varchar(50) not null,
created_at timestamp default current_timestamp
);

create table if not exists momentum(
id serial primary key,
user_id varchar(100) not null,
steps int not null,
calories int not null,
distance int not null,
max_heart_rate int not null,
min_heart_rate int not null,
date date not null
);

create table if not exists sleep(
id serial primary key,
user_id varchar(100) not null,
hours int not null,
started_at timestamp not null,
ended_at timestamp not null,
deep_sleep int not null,
light_sleep int not null,
rem_sleep int not null,
wake int not null,
date date not null
);

create table if not exists meal(
id serial primary key,
user_id varchar(100) not null,
meal_name varchar(100) not null,
caloriesper100g int not null,
date date not null
);