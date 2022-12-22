package main

import (
	"context"
	"log"
	"os"

	"github.com/goark/errs"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog"
	"github.com/sRRRs-7/loose_style.git/cfg"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph"
)

var config cfg.Config
var err error
var pool *pgxpool.Pool

// create zerolog instance
func CreateLogger() zerolog.Logger {
	logger := zerolog.New(
		os.Stdout,
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	logger.Err(os.ErrInvalid).Send()
	logger.Error().Interface("error", errs.Wrap(os.ErrInvalid)).Msg("sample error")
	return zerolog.Logger{}
}

func pgxConnection() {
	// get logger
	logger := CreateLogger()

	// create connection pool for PostgreSQL service
	cfg, err := pgxpool.ParseConfig(config.DBsource)
	if err != nil {
		logger.Error().Interface("error", errs.Wrap(err)).Send()
		log.Fatalf("pgx configuration error: %v", err)
	}
	cfg.ConnConfig.Logger = zerologadapter.NewLogger(logger)
	cfg.ConnConfig.LogLevel = pgx.LogLevelDebug

	pool, err = pgxpool.ConnectConfig(context.TODO(), cfg)
	if err != nil {
		logger.Error().Interface("error", errs.Wrap(err)).Send()
		log.Fatalf("cannot connect to database: %v", err)
	}
}

// root function
func main() {
	config, err = cfg.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	// postgres connection
	pgxConnection()
	defer pool.Close()
	// initialize sqlc instance
	store := db.NewStore(pool)
	// all instance connect
	resolver, tokenMaker, err := graph.NewResolver(config, store, pool)
	if err != nil {
		log.Fatalf("new resolver error: %v", err)
	}

	resolver.GinRouter(tokenMaker)
}
