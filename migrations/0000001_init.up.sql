CREATE TYPE currency_enum AS ENUM (
    'USD',
    'RUB',
    'EUR',
    'JPY'
);

CREATE TYPE log_action_enum AS ENUM (
    'GET_ALL_CURRENCY',
    'GET_CURRENT_CURRENCY'
);

CREATE TABLE current_currency
(
    id            bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    currency_type currency_enum            NOT NULL UNIQUE,
    rate          float8                   NOT NULL,
    created_at    timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE historic_currency
(
    id            bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    currency_type currency_enum            NOT NULL,
    currency_date date                     NOT NULL,
    rate          float8                   NOT NULL,
    created_at    timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ix_historic_currency_date ON historic_currency (currency_date);

CREATE TABLE action_log
(
    id         bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    action     log_action_enum          NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
