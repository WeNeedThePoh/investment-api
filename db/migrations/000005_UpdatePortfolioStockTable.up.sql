ALTER TABLE portfolio_stocks RENAME COLUMN stock_id TO symbol;
ALTER TABLE portfolio_stocks ALTER COLUMN symbol TYPE VARCHAR(10) USING symbol::VARCHAR(10);

ALTER TABLE portfolio_transactions RENAME COLUMN stock_id TO symbol;
ALTER TABLE portfolio_transactions ALTER COLUMN symbol TYPE VARCHAR(10) USING symbol::VARCHAR(10);
