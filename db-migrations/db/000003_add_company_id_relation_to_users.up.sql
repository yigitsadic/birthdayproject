alter table users
add column if not exists company_id int not null
constraint users_company_id_fk references companies(id);
