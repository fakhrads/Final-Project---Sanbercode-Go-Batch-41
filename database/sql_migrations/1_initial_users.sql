-- +migrate Up
-- +migrate StatementsBegin

CREATE TABLE tbl_users ( 
	id                   SERIAL PRIMARY KEY  ,
	username             varchar(100)    ,
	"password"           varchar(255)    ,
	created_at           date    ,
	updated_at           date    
 );


-- +migrate StatementEnd