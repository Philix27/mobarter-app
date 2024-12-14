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
