create table companies(
    id serial constraint companies_pk primary key,
    name varchar(50) not null,

    created_at timestamp not null default now(),
    updated_at timestamp not null,

    constraint company_name_uniqueness unique (name) 
);
