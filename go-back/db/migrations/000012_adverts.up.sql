CREATE TABLE adverts (
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