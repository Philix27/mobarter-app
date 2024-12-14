CREATE TABLE transactions (
  id   BIGSERIAL PRIMARY KEY,
  user_id  INT UNIQUE REFERENCES users (id),
  name TEXT      NOT NULL,
  purpose  TransactionPurpose NOT NULL,
  amount  INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);