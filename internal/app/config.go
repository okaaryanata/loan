package app

import (
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AppConfig struct {
	Host   string
	Port   string
	InitDB bool
	DB     *pgxpool.Pool
}

func (app *AppConfig) InitService() {
	app.Host = os.Getenv("APP_HOST")
	app.Port = os.Getenv("APP_PORT")
	app.InitDB, _ = strconv.ParseBool(os.Getenv("DB_INIT_TABLE"))

	// init db
	app.InitPostgresDB()
}
