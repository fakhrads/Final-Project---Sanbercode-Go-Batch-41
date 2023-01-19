-- +migrate Up
-- +migrate StatementsBegin

CREATE  TABLE tbl_orders ( 
	id                   SERIAL PRIMARY KEY,
	users_id             integer  NOT NULL,
	products_id          integer  NOT NULL,
	total_price          integer DEFAULT 0,
	total_bought         integer,
	created_at           date,
	updated_at           date,
	CONSTRAINT unq_tbl_orders_users_id UNIQUE ( users_id ),
	CONSTRAINT unq_tbl_orders_products_id UNIQUE ( products_id ) 
);

-- +migrate StatementEnd