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