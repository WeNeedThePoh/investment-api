DROP INDEX countries_id_uindex;
ALTER TABLE countries DROP CONSTRAINT countries_pk;

DROP INDEX currencies_id_uindex;
ALTER TABLE currencies DROP CONSTRAINT currencies_pk;

DROP INDEX portfolio_history_id_uindex;
ALTER TABLE portfolio_history DROP CONSTRAINT portfolio_history_pk;

DROP INDEX portfolio_transactions_id_uindex;
ALTER TABLE portfolio_transactions DROP CONSTRAINT portfolio_transactions_pk;

DROP INDEX stocks_id_uindex;
ALTER TABLE stocks DROP CONSTRAINT stocks_pk;

DROP INDEX portfolios_id_uindex;
ALTER TABLE portfolios DROP CONSTRAINT portfolios_pk;

DROP INDEX users_id_uindex;
ALTER TABLE users DROP CONSTRAINT users_pk;

DROP INDEX portfolio_stocks_uindex;
ALTER TABLE portfolio_stocks DROP CONSTRAINT portfolio_stocks_pk;
