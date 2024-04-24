-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password CHAR(60) NOT NULL,
  customer_id INTEGER REFERENCES customers(customer_id)
);

-- +migrate StatementEnd