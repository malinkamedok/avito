CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists balance, reserve, report cascade;

create table if not exists balance (
    user_uuid uuid not null unique default uuid_generate_v4(),
    balance bigint not null check (balance >= 0)
);

create table if not exists reserve (
    id serial primary key,
    user_uuid uuid not null,
    reserve bigint not null check ( reserve >= 0 )
);

create table if not exists report (
    id uuid unique default uuid_generate_v4(),
    user_uuid uuid not null,
    service_uuid uuid not null,
    order_uuid uuid not null,
    total_cost bigint not null
);

create or replace function update_balance(userUUID uuid, sum int) returns void as $$
UPDATE balance set balance = sum + (select balance from balance where user_uuid = userUUID)
where user_uuid = userUUID;
$$ language sql;

create or replace function reserve_money(balanceUUID uuid, reserveUUID uuid, amount bigint) returns void as $$
update balance set balance = (select balance from balance where user_uuid = balanceUUID) - amount
where user_uuid = balanceUUID;
update reserve set reserve = amount + (select reserve from reserve where user_uuid = reserveUUID)
where user_uuid = reserveUUID;
$$ language sql;

create or replace function check_money(IN balanceUUID uuid, IN amount bigint, OUT res bool) as
    'select (select balance from balance where user_uuid = balanceUUID) >= amount'
     language sql;