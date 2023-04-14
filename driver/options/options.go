package options

import (
	"database/sql"
	"time"

	"github.com/pcpratheesh/dbmux/entity"
)

// SetConnMaxLifetime
func SetConnMaxLifetime(d time.Duration) entity.ConnectionOptions {
	return func(db *sql.DB) {
		db.SetConnMaxLifetime(d)
	}
}

func SetConnMaxIdleTime(d time.Duration) entity.ConnectionOptions {
	return func(db *sql.DB) {
		db.SetConnMaxIdleTime(d)
	}
}

func SetMaxOpenConns(n int) entity.ConnectionOptions {
	return func(db *sql.DB) {
		db.SetMaxOpenConns(n)
	}
}

func SetMaxIdleConns(n int) entity.ConnectionOptions {
	return func(db *sql.DB) {
		db.SetMaxIdleConns(n)
	}
}
