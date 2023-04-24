create table users(
    id serial constraint users_pk primary key,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email varchar(70) not null,
    password_hash varchar not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null,

    constraint users_email_uniqueness unique (email) 
);
