package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	// Create a JSON logger with minimum level DEBUG.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Load .env file.
	if err := godotenv.Load(); err != nil {
		logger.Error(fmt.Sprintf("unable to load .env file: %s", err.Error()))
		os.Exit(1)
	}

	// Read the addr and dsn flag from the command-line
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsb", os.Getenv("DATABASE_DSN"), "Database DSN")
	flag.Parse()

	// Open database connection
	logger.Info("opening database connection")
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(fmt.Sprintf("error opening database connection: %s", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	app := application{
		logger: logger,
	}

	srv := http.Server{
		Addr:         *addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
		Handler:      app.routes(),
	}

	logger.Info("starting server", slog.String("addr", srv.Addr))
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("error starting up the server", slog.String("message", err.Error()))
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
