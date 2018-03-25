package config

import (
	"github.com/rubenv/sql-migrate"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)



func MigrateDatabase(){
	migrations := &migrate.FileMigrationSource{
		Dir: "config/migration",
	}

	migrate.SetTable("migration_log")

	db, err := sql.Open("sqlite3", "wahlblock.db")
	if err != nil {
		// Handle errors!
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		panic("Can not connect to sqlite database!")
	}

	if n > 0 {
		log.Println("INFO: Applied %d migrations", n)
	} else {
		log.Println("INFO: none migrations made")
	}
}