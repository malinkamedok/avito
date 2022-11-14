CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists balance, reserve, report cascade;

create table if not exists balance (
    user_uuid uuid not null unique default uuid_generate_v4(),
    balance bigint not null check ( balance >= 0 )
);

create table if not exists reserve (
    id uuid unique default uuid_generate_v4(),
    user_uuid uuid not null,
    service_uuid uuid not null,
    order_uuid uuid not null,
    reserve bigint not null check ( reserve >= 0 )
);

create table if not exists report (
    id uuid unique default uuid_generate_v4(),
    user_uuid uuid not null,
    service_uuid uuid not null,
    order_uuid uuid not null,
    money_amount bigint not null
);

create or replace function update_balance(userUUID uuid, sum int) returns void as $$
UPDATE balance set balance = sum + (select balance from balance where user_uuid = userUUID)
where user_uuid = userUUID;
$$ language sql;

create or replace function reserve_money(userUUID uuid, serviceUUID uuid, orderUUID uuid, amount bigint) returns void as $$
update balance set balance = (select balance from balance where user_uuid = userUUID) - amount
where user_uuid = userUUID;
insert into reserve(user_uuid, service_uuid, order_uuid, reserve) values (userUUID, serviceUUID, orderUUID, amount);
$$ language sql;

create or replace function check_money(IN balanceUUID uuid, IN amount bigint, OUT res bool) as
'select (select balance from balance where user_uuid = balanceUUID) >= amount'
language sql;

create or replace function accept_income(userUUID uuid, serviceUUID uuid, orderUUID uuid, amount bigint) returns void as $$
insert into report(user_uuid, service_uuid, order_uuid, money_amount) values (userUUID, serviceUUID, orderUUID, amount);
delete from reserve where user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount
and id in (select id from reserve where user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount LIMIT 1);
$$ language sql;

create or replace function check_required_reserve(IN userUUID uuid, IN serviceUUID uuid, IN orderUUID uuid, IN amount bigint, OUT res bool) as
'select EXISTS(SELECT * FROM reserve WHERE user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount)'
language sql;

create or replace function user_to_user_money_transfer(first_userUUID uuid, second_userUUID uuid, amount bigint) returns void as $$
update balance set balance = (select balance from balance where user_uuid = first_userUUID) - amount
where user_uuid = first_userUUID;
update balance set balance = amount + (select balance from balance where user_uuid = second_userUUID)
where user_uuid = second_userUUID;
$$ language sql;