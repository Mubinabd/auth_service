CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR,
    username VARCHAR ,
    password VARCHAR,
    phone VARCHAR,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at BIGINT default 0
);

CREATE TABLE IF NOT EXISTS card (
    id UUID PRIMARY KEY,
    name VARCHAR,
    card_num VARCHAR(16),
    expired_at TIMESTAMP,
    amount INT default 0,
    password INT,
    card_type card_type,
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at BIGINT default 0
);

CREATE TABLE IF NOT EXISTS transaction (
    id UUID PRIMARY KEY,
    amount INT default 0,
    type TRANSACTION_TYPE,
    description TEXT,
    payment_type payment_type,
    from_card UUID REFERENCES card(id),
    to_card UUID default null REFERENCES card(id),
    payment_name varchar default null,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at BIGINT default 0
);

-- Enum definitions
CREATE TYPE card_type AS ENUM ('uzcard', 'humo', 'visa');
CREATE TYPE TRANSACTION_TYPE AS ENUM ('debit', 'credit');
create type payment_type as enum('transfer', 'utility', 'fine', 'tax')