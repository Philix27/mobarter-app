CREATE TABLE bank_account (
  id   BIGSERIAL PRIMARY KEY,
  bank_name TEXT      NOT NULL,
  account_name TEXT      NOT NULL,
  account_no INTEGER      NOT NULL,
  user_id  INT REFERENCES users (id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);