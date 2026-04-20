package main

import (
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/dach-trier/env"
	"github.com/dach-trier/portal/internal/app"
)

func main() {
	if err := env.LoadFile(os.DirFS("."), ".env"); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			panic(err)
		}
	}

	// --
	// Server
	// --

	port := env.GetIntInRange(os.Getenv, "PORT", 8080, 1, 65535)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: app.New().Router(),
	}

	log.Println("listening on :" + strconv.Itoa(port))
	log.Fatal(server.ListenAndServe())
}
