package main

import (
	"context"
	"ecom-local/internal/env"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()

	var dbFallbackString string = "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("TESTENV", dbFallbackString),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	conn, dbError := pgx.Connect(ctx, cfg.db.dsn)
	if dbError != nil {
		panic(dbError)
	}
	defer conn.Close(ctx)

	logger.Info("Connected to database", "dsn", cfg.db.dsn)

	app := application{
		config: cfg,
	}

	h := app.mount()

	err := app.run(h)
	if err != nil {
		log.Println("Server has failed to start")
		os.Exit(1)
	}
}
