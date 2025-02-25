create table if not exists tokenized (
    id serial primary key,
    quote_token_id integer not null references token(id),
    base_token_id integer not null references token(id),
    price integer not null check (price > 0),
    maturity_id integer not null references maturities(id),
    rate numeric(5,2) not null check (rate > 0),
    amount integer not null check (amount > 0),
    volume integer not null check (volume > 0),
    order_type varchar(255) check (order_type in ('BUY', 'SELL')) not null
);