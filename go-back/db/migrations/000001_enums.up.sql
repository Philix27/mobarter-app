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

