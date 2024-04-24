-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transactions (
  transaction_id SERIAL PRIMARY KEY,
  account_id INTEGER NOT NULL REFERENCES accounts(account_id),
  transaction_type VARCHAR(20) CHECK (transaction_type IN ('Deposit', 'Withdrawal', 'Transfer')),
  amount DECIMAL(10,2) NOT NULL,
  transaction_date TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  description TEXT
);


-- +migrate StatementEnd