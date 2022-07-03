package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Jay-Gardiner/go-simple-api/pkg/api"
	"github.com/Jay-Gardiner/go-simple-api/pkg/app"
	repository "github.com/Jay-Gardiner/go-simple-api/pkg/repo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Application entry point
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Encountered startup error: %s\\n", err)
		os.Exit(1)
	}
}

// Responsible for initialisation of DB connection, routes, etc
func run() error {
	// Setup database connection
	connectionString := "postgres://postgres:9PHA%24%24%25X6atiGiu%25C@192.168.88.3/go-simple-api?sslmode=disable"
	db, err := setupDatabase(connectionString)

	if err != nil {
		return err
	}

	// create storage dependency
	storage := repository.NewStorage(db)

	// create services
	dataService := api.NewDataService(storage)

	// Setup route dependency
	router := gin.Default()
	router.Use(cors.Default())

	// create the server
	server := app.NewServer(router, dataService)

	// run the server
	err = server.RunServer()

	// return errors
	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
