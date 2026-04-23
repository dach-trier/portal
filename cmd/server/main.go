package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dach-trier/env"
	"github.com/dach-trier/portal/internal/app"
	"github.com/dach-trier/portal/internal/repo"
	"github.com/dach-trier/portal/internal/repo/sqlc"
)

func main() {
	if err := env.LoadFile(os.DirFS("."), ".env"); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			panic(err)
		}
	}

	// --
	// Database
	// --

	var pool *pgxpool.Pool
	var err error

	ctx := context.Background()

	// database should be exposed over a local network and not publicly
	// available, so sslmode=disable is what we want here
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	if pool, err = pgxpool.New(ctx, dsn); err != nil {
		panic(err)
	}

	defer pool.Close()

	if err = pool.Ping(ctx); err != nil {
		panic(err)
	}

	repos := repo.Bundle{
		Projects: sqlc_repo.NewProjectRepository(pool),
	}

	// --
	// Server
	// --

	port := env.GetIntInRange(os.Getenv, "PORT", 8080, 1, 65535)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: app.New(repos).Router(),
	}

	log.Println("listening on :" + strconv.Itoa(port))
	log.Fatal(server.ListenAndServe())
}
