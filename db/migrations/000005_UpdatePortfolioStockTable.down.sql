ALTER TABLE portfolio_stocks RENAME COLUMN symbol TO stock_id;
ALTER TABLE portfolio_stocks ALTER COLUMN stock_id TYPE INTEGER USING stock_id::INTEGER;

ALTER TABLE portfolio_transactions RENAME COLUMN symbol TO stock_id;
ALTER TABLE portfolio_transactions ALTER COLUMN stock_id TYPE INTEGER USING stock_id::INTEGER;
