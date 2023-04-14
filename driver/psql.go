package driver

import (
	"database/sql"
	"fmt"

	"github.com/pcpratheesh/dbmux/entity"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func postgresConn() *Postgres {
	return &Postgres{}
}

// initiate connection with postgres
func (ps *Postgres) Connect(param entity.Options) error {
	if param.SSLMode == "" {
		param.SSLMode = "disable"
	}

	if param.Schema == "" {
		param.Schema = "public"
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s search_path=%s",
		param.Host, param.Port, param.User, param.Password, param.Database, param.SSLMode, param.Schema)

	db, err := sql.Open(POSTGRES, connectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	ps.db = db

	return nil
}

func (ps *Postgres) Close() error {
	return ps.db.Close()
}

func (ms *Postgres) Pool() interface{} {
	return ms.db
}
