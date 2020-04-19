INSERT INTO currencies (name, symbol)
VALUES
('Euro', 'EUR'),
('British Pound', 'GBP'),
('US Dollar', 'USD');

INSERT INTO countries (currency_id, name, iso_code)
VALUES
(1, 'Portugal', 'PT'),
(2, 'United Kingdom', 'UK'),
(3, 'United States', 'US');

-- STOCKS TEST DATA
INSERT INTO stocks (country_id, symbol, company_name, price, min_price, max_price, div_yield, div_share, eps)
VALUES
(3, 'AAPL', 'Apple inc.', 284.43, 170.27, 327.85, 1.07, 0.77, 12.60),
(3, 'MSFT', 'Microsoft Corporation', 171.88, 119.01, 190.70, 1.23, 0.51, 5.74);
