package app

import (
	"app/pkg/config"
	"app/pkg/database"
	"app/pkg/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

// loggers, redis and other things can be added here, application state

type App struct {
	Config  *config.Config
	Db      *pgxpool.Pool
	Queries *sqlc.Queries
}

func NewApp(cfg *config.Config) (*App, error) {

	if err := database.RunMigrations(cfg.DB); err != nil {
		panic("failed to migrate db")
	}

	//setup database Connection Pool
	db, err := database.NewPostgresPool(cfg.DB)

	if err != nil {
		panic("Error starting up new DB Pool")
	}

	queries := sqlc.New(db)

	return &App{
		Config:  cfg,
		Db:      db,
		Queries: queries,
	}, nil

}
