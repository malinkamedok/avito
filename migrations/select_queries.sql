create or replace procedure update_balance(userUUID uuid, sum int)
    language sql
as $$
UPDATE balance set balance = sum + (select balance from balance where user_uuid = userUUID)
where user_uuid = userUUID;
$$;

call update_balance('7e931dee-aa59-4687-a18a-44ea256fc733', 100);