-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transactions (
  transaction_id SERIAL PRIMARY KEY,
  account_id INTEGER REFERENCES accounts(account_id) NOT NULL,
  transaction_type VARCHAR(20) NOT NULL,  -- Deposit, Withdrawal, Transfer
  amount DECIMAL(10,2) NOT NULL,
  description VARCHAR(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd