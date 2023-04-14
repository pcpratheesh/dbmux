// Package dbmux provides a way to manage multiple database connections in a single place.
// It provides a MultiConn type that holds a pool of database connections, and methods to add new database connections to the pool,
// retrieve a connection from the pool, and get the connection pool.

package dbmux

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pcpratheesh/dbmux/driver"
	"github.com/pcpratheesh/dbmux/entity"
)

type MultiConn struct {
	conns entity.Pool
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Creates and returns a new MultiConn object, which manages a pool of database connections
func New(options ...entity.Options) *MultiConn {
	var multiConn = &MultiConn{
		conns: make(entity.Pool),
	}

	for _, opt := range options {
		err := multiConn.AddConnection(opt)
		if err != nil {
			panic(err)
		}
	}

	return multiConn
}

// AddConnection
// Adds a new database connection to the pool of connections.
// It takes an entity.Options structure as an argument that specifies the details of the connection (name, driver, host, port, etc).
// If a connection with the same name already exists in the pool, it returns an error indicating that the name is already taken.
// Otherwise, it connects to the database using the specified driver and adds the connection to the pool.
func (multi *MultiConn) AddConnection(option entity.Options) error {
	err := validate.Struct(option)
	if err != nil {
		return err
	}

	// if not, choose driver handler
	driverHandler, err := driver.ChooseDriver(option.Driver)
	if err != nil {
		return err
	}

	// connect the db
	err = driverHandler.Connect(option)
	if err != nil {
		return err
	}

	// add connection to pool
	conn := driverHandler.Pool()
	err = multi.AddToPool(option, conn)
	if err != nil {
		return err
	}

	// add to db pool
	return nil
}

// AddToPool
// This adds a new database connection to the pool of connections
// The function checks if a connection with the same name already exists in the pool.
// If it does, it returns an error indicating that the name is already taken
func (multi *MultiConn) AddToPool(option entity.Options, db interface{}) error {
	if _, ok := multi.conns[option.Name]; ok {
		return fmt.Errorf("name %s already exists for another connection in the pool", option.Name)
	}
	multi.conns[option.Name] = db

	return nil
}

// GetConnection
// returns a database connection object given its name.
// It takes in a string parameter name which represents the name of the connection that needs to be retrieved.
// This method can be used to retrieve a database connection object from the pool using its name.
// If the connection is not present in the pool, it returns an error indicating that the connection could not be found.
func (multi *MultiConn) GetConnection(name string) (interface{}, error) {
	if conn, ok := multi.conns[name]; ok {
		return conn, nil
	} else {
		return nil, fmt.Errorf("unable to find the db connection %s in the pool", name)
	}
}

// GetConnectionPool
// function performs a type assertion to convert the connection interface type into a value of type T.
// The type assertion ensures that connection implements the SqlPoolInterface, otherwise, it will panic
func GetConnectionPool[T entity.SqlPoolInterface](connection interface{}) T {
	return connection.(T)
}
