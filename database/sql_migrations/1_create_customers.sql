-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE customers (
  customer_id SERIAL PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50),
  email VARCHAR(100) UNIQUE NOT NULL,
  phone_number VARCHAR(20)
);

-- +migrate StatementEnd