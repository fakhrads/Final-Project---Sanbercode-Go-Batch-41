package database

import (
	"database/sql"
	"fmt"

	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func DbMigrate(dbParams *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}
	DbConnection = dbParams
	n, errs := migrate.Exec(dbParams, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	fmt.Println("Applied ", n, "migrations!")
}
