CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists balance, reserve, report cascade;

create table if not exists balance (
    user_uuid uuid not null unique default uuid_generate_v4(),
    balance bigint not null
);

create table if not exists reserve (
    user_uuid uuid not null unique default uuid_generate_v4(),
    reserve bigint not null
);

create table if not exists report (
    id uuid unique default uuid_generate_v4(),
    user_uuid uuid not null,
    service_uuid uuid not null,
    order_uuid uuid not null,
    total_cost bigint not null
);