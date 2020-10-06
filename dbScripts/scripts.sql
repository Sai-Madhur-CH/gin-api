create schema config;

create table config.user_details(
user_id serial primary key,
name varchar not null,
email varchar not null,
password_hash varchar not null,
created_at timestamp,
updated_at timestamp
);