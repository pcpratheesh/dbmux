package dbmux

import (
	"database/sql"
	"testing"

	"github.com/pcpratheesh/dbmux/entity"
)

func TestAddToPool(t *testing.T) {
	// Create a new MultiConn instance
	multi := &MultiConn{
		conns: make(entity.Pool),
	}

	// Create a new mock database connection to add to the pool
	db := &sql.DB{}

	// Create a new Params struct to use as input
	params := entity.Options{
		Name:   "test_db",
		Driver: "sqlite3",
		// URL:    "file::memory:?cache=shared",
	}

	// Call AddToPool with the input params and database
	err := multi.AddToPool(params, db)

	// Check that there was no error returned
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the database was added to the pool
	if _, ok := multi.conns[params.Name]; !ok {
		t.Errorf("expected database to be added to the pool")
	}

	// Call AddToPool again with the same input params and database
	err = multi.AddToPool(params, db)

	// Check that an error was returned indicating the name is already taken
	if err == nil {
		t.Errorf("expected error due to name already existing in the pool")
	}
}

type Person struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
