package database

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	allMigrations, err := migrations.FindMigrations()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found %d migrations\n", len(allMigrations))

	for _, m := range allMigrations {
		fmt.Printf("Migration: %s\n", m.Id)
	}

	n, err := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Printf("Error: Could not migrate the database: %v\n", err)
		panic(err)
	}

	DbConnection = dbParam

	fmt.Printf("Applied %d migrations!\n", n)
}
