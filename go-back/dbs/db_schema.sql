CREATE TYPE CurrencyPair AS ENUM ('CUSD_NGN', 'USDT_NGN', 'CELO_NGN');
CREATE TYPE TradeType AS ENUM ('BUY', 'SELL');
CREATE TYPE AgentStatus AS ENUM ('TRADE', 'BLOCKED', 'IN_ACTIVE', 'BANNED', 'SUSPENDED');
CREATE TYPE TransactionPurpose AS 
  ENUM (
    'AIRTIME', 
    'DATA', 
    'TRANSFER', 
    'WITHDRAW', 
    'REFUND'
  );
CREATE TYPE AdsStatus AS ENUM ('UP', 'DOWN');
CREATE TYPE KycValidationStatus AS ENUM ('NONE', 'APPROVED', 'PENDING', 'DECLINED');
CREATE TYPE OrderType AS ENUM ('BUY', 'SELL');
CREATE TYPE OrderStatus AS ENUM ('CREATED', 'COMPLETED', 'APPEAL');

CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name TEXT      NOT NULL,
  bio  TEXT
);

CREATE TABLE airtime (
  id   BIGSERIAL PRIMARY KEY,
  network TEXT      NOT NULL,
  country  TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE data_plans (
  id   BIGSERIAL PRIMARY KEY,
  network TEXT      NOT NULL,
  country  TEXT NOT NULL,
  amount  INTEGER NOT NULL,
  duration  TEXT NOT NULL,
  plan  TEXT NOT NULL,
  is_supported  TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE country (
  id   BIGSERIAL PRIMARY KEY,
  name TEXT   UNIQUE     NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bank_account (
  id   BIGSERIAL PRIMARY KEY,
  bank_name TEXT      NOT NULL,
  account_name TEXT      NOT NULL,
  account_no INTEGER      NOT NULL,
  user_id  INT REFERENCES users (id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE waitlist (
  id   BIGSERIAL PRIMARY KEY,
  email TEXT      NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE newsletter (
  id   BIGSERIAL PRIMARY KEY,
  email TEXT      NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
  id   BIGSERIAL PRIMARY KEY,
  name TEXT      NOT NULL,
  purpose  TransactionPurpose NOT NULL,
  amount  INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE agents (
  id   BIGSERIAL PRIMARY KEY,
  user_id  INT UNIQUE REFERENCES users (id),
  agent_status AgentStatus,
  trade_count INT,
  nick_name TEXT,
  instructions TEXT[],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ads (
  id   BIGSERIAL PRIMARY KEY,
  agent_id  INT REFERENCES agents (id),
  payment_method_id INT REFERENCES bank_account  (id),
  ad_status AdsStatus DEFAULT 'DOWN',
  currency_pair CurrencyPair,
  limit_upper DECIMAL,
  limit_lower DECIMAL,
  rate DECIMAL,
  instructions TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
  id   BIGSERIAL PRIMARY KEY,
  user_id  INT REFERENCES users (id),
  agent_id  INT REFERENCES users (id),
  order_type OrderType,
  order_status OrderStatus,
  currency_pair CurrencyPair,
  amount DECIMAL,
  rate DECIMAL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE customer_care (
  id   BIGSERIAL PRIMARY KEY,
  message TEXT      NOT NULL,
  user_id  TEXT,
  room_id  TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  wallets TEXT[],
  first_name TEXT,
  last_name TEXT,
  dob TIMESTAMP,
  email TEXT  UNIQUE  NOT NULL,
  phone TEXT  UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the kyc_credentials table
CREATE TABLE kyc_credentials (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE NOT NULL REFERENCES users (id),
    bvn TEXT      ,
    bvn_status KycValidationStatus  DEFAULT "NONE",
    bvn_verification_date TIMESTAMP,
    bvn_verified_by TEXT,
    nin TEXT,
    nin_status KycValidationStatus  DEFAULT "NONE",
    nin_verification_date TIMESTAMP,
    nin_verified_by TEXT,
    house_address TEXT,
    house_address_status KycValidationStatus  DEFAULT "NONE",
    house_address_verification_date TIMESTAMP,
    house_address_verified_by TEXT,
    utility_bill TEXT,
    utility_bill_status KycValidationStatus  DEFAULT "NONE",
    utility_bill_verification_date TIMESTAMP,
    utility_bill_verified_by TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--! FUnctions

-- Create a function to update `updated_at`
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to call the function on updates
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON table_name
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
