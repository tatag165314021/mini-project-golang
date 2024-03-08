package database

import (
	"database/sql"
)

var (
	DbConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB) {
	// migrations := &migrate.PackrMigrationSource{
	// 	Box: packr.New("migrations", "./sql_migrations"),
	// }

	// n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	// if errs != nil {
	// 	panic(errs)
	// }
	DbConnection = dbParam

	// fmt.Println("Applied ", n, " migrations!")
}
