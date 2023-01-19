-- +migrate Up
-- +migrate StatementsBegin

CREATE  TABLE tbl_products ( 
	id                   SERIAL PRIMARY KEY,
	product_name         varchar(100)  NOT NULL,
	product_description  varchar,
	product_stock        integer DEFAULT 0,
	product_price        integer DEFAULT 0,
	created_at           date,
	updated_at           date
 );

-- +migrate StatementEnd