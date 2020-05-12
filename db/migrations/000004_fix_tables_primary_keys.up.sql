CREATE UNIQUE INDEX countries_id_uindex ON countries (id);
ALTER TABLE countries ADD CONSTRAINT countries_pk PRIMARY KEY (id);

CREATE UNIQUE INDEX currencies_id_uindex ON currencies (id);
ALTER TABLE currencies ADD CONSTRAINT currencies_pk PRIMARY KEY (id);

CREATE UNIQUE INDEX portfolio_history_id_uindex ON portfolio_history (id);
ALTER TABLE portfolio_history ADD CONSTRAINT portfolio_history_pk PRIMARY KEY (id);

CREATE UNIQUE INDEX portfolio_transactions_id_uindex ON portfolio_transactions (id);
ALTER TABLE portfolio_transactions ADD CONSTRAINT portfolio_transactions_pk PRIMARY KEY (id);

CREATE UNIQUE INDEX stocks_id_uindex ON stocks (id);
ALTER TABLE stocks ADD CONSTRAINT stocks_pk PRIMARY KEY (id);

CREATE UNIQUE INDEX portfolios_id_uindex ON portfolios (id);
ALTER TABLE portfolios ADD CONSTRAINT portfolios_pk PRIMARY KEY (id);

CREATE UNIQUE INDEX users_id_uindex ON users (id);
ALTER TABLE users ADD CONSTRAINT users_pk PRIMARY KEY (id);

CREATE UNIQUE INDEX portfolio_stocks_uindex ON portfolio_stocks (portfolio_id, stock_id);
ALTER TABLE portfolio_stocks ADD CONSTRAINT portfolio_stocks_pk PRIMARY KEY (portfolio_id, stock_id);
