create table if not exists employees(
    id serial constraint employees_pk primary key,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email varchar(70) not null,
    birth_day int not null,
    birth_month int not null,
    company_id int not null,

    created_at timestamp not null default now(),
    updated_at timestamp not null,

    constraint employee_companies_fk foreign key(company_id) references companies(id),
    constraint employees_email_uniqueness unique (email) 
);
