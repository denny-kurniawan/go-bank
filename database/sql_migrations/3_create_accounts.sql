-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE accounts (
  account_id SERIAL PRIMARY KEY,
  customer_id INTEGER NOT NULL REFERENCES customers(customer_id),
  account_number VARCHAR(20) UNIQUE NOT NULL,
  balance DECIMAL(10,2) DEFAULT 0.00 NOT NULL
);

-- +migrate StatementEnd