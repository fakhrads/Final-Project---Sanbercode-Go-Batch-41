-- +migrate Up
-- +migrate StatementsBegin

ALTER TABLE tbl_orders ADD CONSTRAINT fk_tbl_orders_tbl_users FOREIGN KEY ( users_id ) REFERENCES tbl_users( id ) ON DELETE CASCADE;
ALTER TABLE tbl_orders ADD CONSTRAINT fk_tbl_orders_tbl_products FOREIGN KEY ( products_id ) REFERENCES tbl_products( id ) ON DELETE CASCADE;
-- +migrate StatementEnd
