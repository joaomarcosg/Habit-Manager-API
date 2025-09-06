package mysqlstore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MSUserStore struct {
	Queries *Queries
	Pool    *sql.DB
}

func NewMSUserStore(pool *sql.DB) MSUserStore {
	return MSUserStore{
		Queries: New(pool),
		Pool:    pool,
	}
}
