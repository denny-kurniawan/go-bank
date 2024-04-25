-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE accounts (
  account_id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(user_id) NOT NULL,
  account_number VARCHAR(10) UNIQUE NOT NULL,
  balance DECIMAL(12,2) DEFAULT 100000.00 NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd