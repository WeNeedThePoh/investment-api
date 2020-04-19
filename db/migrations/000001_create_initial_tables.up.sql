-- CURRENCIES
CREATE TABLE currencies (
    id SERIAL,
    name VARCHAR(20) NOT NULL,
    symbol VARCHAR(3) NOT NULL
);

-- COUNTRIES
CREATE TABLE countries (
    id SERIAL,
    currency_id INT NOT NULL,
    name VARCHAR(25) NOT NULL,
    iso_code VARCHAR(3) NOT NULL
);

-- USERS
CREATE TABLE users (
    id SERIAL,
    currency_id INT,
    country_id INT NOT NULL,
    lang_id INT,
    email VARCHAR(75) NOT NULL,
    first_name VARCHAR(40) NOT NULL,
    last_name VARCHAR(40),
    password VARCHAR(255) NOT NULL,
    password_reset VARCHAR(255),
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- STOCKS
CREATE TABLE stocks (
    id SERIAL,
    country_id INT NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    company_name VARCHAR(50) NOT NULL,
    price FLOAT NOT NULL,
    min_price FLOAT NOT NULL,
    max_price FLOAT NOT NULL,
    daily_change FLOAT DEFAULT 0,
    daily_change_percentage FLOAT DEFAULT 0,
    year_change FLOAT DEFAULT 0,
    year_change_percentage FLOAT DEFAULT 0,
    div_yield FLOAT,
    div_share FLOAT,
    eps FLOAT
);

-- PORTFOLIOS
CREATE TABLE portfolios (
    id SERIAL,
    user_id INT NOT NULL,
    currency_id INT NOT NULL,
    name VARCHAR(40) NOT NULL,
    cost FLOAT DEFAULT 0,
    market_value FLOAT DEFAULT 0,
    total_change FLOAT DEFAULT 0,
    total_change_percentage FLOAT DEFAULT 0,
    daily_change FLOAT DEFAULT 0,
    daily_change_percentage FLOAT DEFAULT 0,
    unrealised_gain_loss FLOAT DEFAULT 0,
    realised_gain_loss FLOAT DEFAULT 0,
    expected_div_yield FLOAT DEFAULT 0,
    expected_div FLOAT DEFAULT 0,
    div_collected FLOAT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- PORTFOLIO STOCKS
CREATE TABLE portfolio_stocks (
    portfolio_id INT NOT NULL,
    stock_id INT NOT NULL,
    type VARCHAR(10) NOT NULL,
    shares FLOAT,
    avg_share_cost FLOAT,
    cost FLOAT,
    market_value FLOAT DEFAULT 0,
    total_change FLOAT DEFAULT 0,
    total_change_percentage FLOAT DEFAULT 0,
    daily_change FLOAT DEFAULT 0,
    daily_change_percentage FLOAT DEFAULT 0,
    unrealised_gain_loss FLOAT DEFAULT 0,
    realised_gain_loss FLOAT DEFAULT 0,
    expected_div_yield FLOAT DEFAULT 0,
    expected_div FLOAT DEFAULT 0,
    div_collected FLOAT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- TRANSACTIONS
CREATE TABLE portfolio_transactions (
    id SERIAL,
    portfolio_id INT NOT NULL,
    stock_id INT NOT NULL,
    type VARCHAR(10) NOT NULL,
    shares FLOAT,
    amount FLOAT,
    cost_per_share FLOAT,
    fees FLOAT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- HISTORY
CREATE TABLE portfolio_history (
    id SERIAL,
    portfolio_id INT NOT NULL,
    market_value FLOAT,
    gain_loss FLOAT,
    daily_gain_loss FLOAT  DEFAULT 0,
    div_collected FLOAT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW()
);
