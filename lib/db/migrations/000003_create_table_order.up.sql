create table if not exists orders (
    id serial primary key,
    collateral_token_id integer not null references token(id),
    debt_token_id integer not null references token(id),
    maturity_id integer not null references maturities(id),
    rate numeric(5,2) not null check (rate > 0),
    available_token integer not null check (available_token > 0),
    order_type varchar(255) check (order_type in ('LEND', 'BORROW')) not null
);