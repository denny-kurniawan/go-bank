-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE sessions (
  session_id CHAR(36) PRIMARY KEY,  -- Use a UUID for unique identifier
  user_id INTEGER REFERENCES users(user_id) NOT NULL,
  token TEXT NOT NULL, 
  expires_at TIMESTAMP WITH TIME ZONE DEFAULT (CURRENT_TIMESTAMP + INTERVAL '10 minutes'),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd