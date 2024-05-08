create table cars(
    id uuid primary key default gen_random_uuid(),
    model varchar,
    year int,
    num varchar,
    color varchar,
    owner varchar,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at bigint default 0
);

-- create index <index_name> on <table_name>(<col_name>)
-- single column
create index idx_cars_id on cars(id);

-- multi column
create index idx_cars_id_num on cars(id);

create unique index idx_cars_unq_id on cars(id, name);
create unique index idx_cars_unq_id on cars(name);
