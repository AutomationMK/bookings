package dbrepo

import (
	"github.com/AutomationMK/bookings/internal/config"
	"github.com/AutomationMK/bookings/internal/repository"
	"github.com/jackc/pgx/v5"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *pgx.Conn
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *pgx.Conn
}

func NewPostgresRepo(conn *pgx.Conn, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
