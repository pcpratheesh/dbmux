package driver

import (
	"fmt"

	"github.com/pcpratheesh/dbmux/entity"
)

const (
	POSTGRES = "postgres"
	MYSQL    = "mysql"
	MONGO    = "mongo"
)

type driver interface {
	Connect(entity.Options) error
	Close() error
	Pool() interface{}
}

// ChooseDriver
// The ChooseDriver function is used to retrieve a database driver's connection object based on the given driver name.
// It takes a string parameter driver, which represents the type of the database driver.
// The purpose of this function is to provide a convenient way to establish a connection to various types of databases
// without exposing the underlying implementation details.
// It encapsulates the logic required to select the appropriate connection method for a specific database driver and
// returns an instance of the corresponding driver's connection object.
func ChooseDriver(driver string) (driver, error) {
	switch driver {
	case MYSQL:
		return mysqlConn(), nil
	case POSTGRES:
		return postgresConn(), nil
	case MONGO:
		return mongoConn(), nil
	default:
		return nil, fmt.Errorf("%s driver not borded yet", driver)
	}
}
