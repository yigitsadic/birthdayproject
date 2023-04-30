drop table if exists users;
drop table if exists companies;
drop table if exists employees;

-- users
create table users(
    id int,
    company_id int,
    first_name varchar,
    last_name varchar,
    email varchar,
	password_hash varchar,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

insert into users(first_name, last_name, email, password_hash, company_id, id, created_at, updated_at) values('john', 'doe', 'johndo@google.com', '$2a$14$QHCciI5qzpoXF0tnD/4uSOo0U9kMgIeKxESzFESq6Annk38Z1gZoi', 1, 1, now(), now());

-- companies
create table companies(
    id int,
    name varchar,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

insert into companies (name, id, created_at, updated_at) values('Acme Inc.', 1, now(), now());

-- employees
create table employees(
    id serial,
    company_id int,
    first_name varchar,
    last_name varchar,
    email varchar,
    birth_day int,
    birth_month int,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

insert into employees (first_name, last_name, email, birth_day, birth_month, company_id, id, created_at, updated_at)
values ('yigit', 'sadic', 'yigit@google.com', 13, 2, 1, 1, now(), now());

