package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/dach-trier/env"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var err error

	if err = env.LoadFile(os.DirFS("."), ".env"); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			panic(err)
		}
	}

	// database should be exposed over a local network and not publicly
	// available, so sslmode=disable is what we want here
	dsn := fmt.Sprintf(
		"pgx5://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var m *migrate.Migrate

	if m, err = migrate.New("file://database/migrations", dsn); err != nil {
		panic(err)
	}

	m.Up()
}
