insert into companies (name, created_at, updated_at)
values ('Acme Inc.', now(), now());

insert into employees(first_name, last_name, email, birth_day, birth_month, company_id, updated_at, created_at)
values (
    'Yigit', 'Sadic', 'yigit@google.com', 15, 4,
    (select companies.id from companies where companies.name = 'Acme Inc.' limit 1),
    now(), now()
);

insert into users(first_name, last_name, email, password_hash, company_id, updated_at, created_at)
values(
    'Yigit', 'Sadic', 'sadic@google.com',
    '$2a$14$QHCciI5qzpoXF0tnD/4uSOo0U9kMgIeKxESzFESq6Annk38Z1gZoi',
    (select companies.id from companies where companies.name = 'Acme Inc.' limit 1),
    now(), now()
);
