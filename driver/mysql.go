package driver

import (
	"database/sql"
	"fmt"

	"github.com/pcpratheesh/dbmux/entity"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	db *sql.DB
}

func mysqlConn() *Mysql {
	return &Mysql{}
}

// initiate connection with mysql
func (ms *Mysql) Connect(param entity.Options) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		param.User, param.Password, param.Host, param.Port, param.Database)

	db, err := sql.Open(MYSQL, connectionString)
	if err != nil {
		return err
	}

	for _, opt := range param.ConnectionOptions {
		opt(db)
	}

	if err := db.Ping(); err != nil {
		return err
	}

	// setting the db
	ms.db = db

	return nil
}

// Close
func (ms *Mysql) Close() error {
	return ms.db.Close()
}

// Pool
func (ms *Mysql) Pool() interface{} {
	return ms.db
}
