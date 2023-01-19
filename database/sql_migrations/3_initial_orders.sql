-- +migrate Up
-- +migrate StatementsBegin

CREATE  TABLE tbl_orders ( 
	id                   SERIAL PRIMARY KEY,
	users_id             integer  NOT NULL,
	products_id          integer  NOT NULL,
	total_price          integer DEFAULT 0,
	total_bought         integer,
	created_at           date,
	updated_at           date
);

-- +migrate StatementEnd