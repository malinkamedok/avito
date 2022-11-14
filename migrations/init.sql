CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE serviceEnum AS ENUM ('Append Balance', 'Service transport posting payment', 'Service real estate posting payment', 'Receiving money from user', 'Sending money to user');

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
    service_uuid uuid,
    service_name serviceEnum not null,
    order_uuid uuid,
    money_amount bigint not null,
    operation_date timestamp not null
);

create or replace function update_balance(userUUID uuid, sum bigint) returns void as $$
UPDATE balance set balance = sum + (select balance from balance where user_uuid = userUUID)
where user_uuid = userUUID;
insert into report(user_uuid, service_name, money_amount, operation_date) values (userUUID, 'Append Balance', sum, CURRENT_TIMESTAMP);
$$ language sql;

create or replace function reserve_money(userUUID uuid, serviceUUID uuid, orderUUID uuid, amount bigint) returns void as $$
update balance set balance = (select balance from balance where user_uuid = userUUID) - amount
where user_uuid = userUUID;
insert into reserve(user_uuid, service_uuid, order_uuid, reserve) values (userUUID, serviceUUID, orderUUID, amount);
$$ language sql;

create or replace function check_money(IN balanceUUID uuid, IN amount bigint, OUT res bool) as
'select (select balance from balance where user_uuid = balanceUUID) >= amount'
language sql;

create or replace function accept_income(userUUID uuid, serviceUUID uuid, serviceName serviceEnum, orderUUID uuid, amount bigint) returns void as $$
insert into report(user_uuid, service_uuid, service_name, order_uuid, money_amount, operation_date) values (userUUID, serviceUUID, serviceName, orderUUID, amount, CURRENT_TIMESTAMP);
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
insert into report(user_uuid, service_name, money_amount, operation_date) values (first_userUUID, 'Sending money to user', amount, CURRENT_TIMESTAMP),
                                                                                 (second_userUUID, 'Receiving money from user', amount, CURRENT_TIMESTAMP);
$$ language sql;

create or replace function unreserve_money(userUUID uuid, serviceUUID uuid, orderUUID uuid, amount bigint) returns void as $$
delete from reserve where user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount
and id in (select id from reserve where user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount LIMIT 1);
UPDATE balance set balance = amount + (select balance from balance where user_uuid = userUUID)
where user_uuid = userUUID;
$$ language sql;

-- Задача: необходимо предоставить метод получения списка транзакций с комментариями
-- откуда и зачем были начислены/списаны средства с баланса.
-- Необходимо предусмотреть пагинацию и сортировку по сумме и дате.


--select service_name, money_amount, operation_date from report where user_uuid = '9dcb3c84-139f-4b80-b3ca-115a8c64fc60' order by operation_date;

create or replace function create_new_balance(userUUID uuid, userBalance bigint) returns void as $$
INSERT INTO balance(user_uuid, balance) VALUES(userUUID, userBalance);
insert into report(user_uuid, service_name, money_amount, operation_date) values (userUUID, 'Append Balance', userBalance, CURRENT_TIMESTAMP);
$$ language sql;