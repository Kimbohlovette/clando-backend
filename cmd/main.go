package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"

	"github.com/kimbohlovette/clando-backend/db"
	"github.com/kimbohlovette/clando-backend/models"
	"github.com/kimbohlovette/clando-backend/server"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	_ = godotenv.Load()

	var cfg models.Config
	help, err := conf.Parse("", &cfg)
	if err != nil {
		if err == conf.ErrHelpWanted {
			fmt.Println(help)
			os.Exit(0)
		}
		log.Fatal("parsing config:", err)
	}

	sslMode := "disable"
	if cfg.DB.SSLMode {
		sslMode = "require"
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, sslMode)

	if err := db.Migrate(dbURL, "db/migrations", logger); err != nil {
		log.Fatal("cannot run migrations:", err)
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer pool.Close()

	store := db.NewStore(pool)
	srv := server.NewServer(store)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := srv.Start(":" + cfg.Port); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
