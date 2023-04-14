package entity

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Options struct {
	Name     string `json:"name" validate:"required"`
	Driver   string `json:"driver" validate:"required"`
	Host     string `json:"host" validate:"required"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	Database string `json:"database" validate:"required"`
	SSLMode  string `json:"ssl_mode" default:"disable"`
	Schema   string `json:"schema" default:"public"`

	Port              int                 `json:"port" validate:"required"`
	ConnectionOptions []ConnectionOptions `json:"connection_options"`
	Retries           int                 `json:"retries"`
}

type Pool map[string]interface{}

type SqlPoolInterface interface {
	*sql.DB | *mongo.Client
}

type MongoClient = *mongo.Client
type SqlDB = *sql.DB

type ConnectionOptions func(db *sql.DB)
