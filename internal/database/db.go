package database

import (
	"database/sql"
	"time"

	"github.com/safecornerscoffee/echo-mvc/internal/config"

	_ "github.com/lib/pq"
)

func NewDB(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 360; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		return nil, err
	}

	return db, err
}
