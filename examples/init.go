package main

import (
	"fmt"
	"time"

	"github.com/pcpratheesh/dbmux"
	"github.com/pcpratheesh/dbmux/driver"
	"github.com/pcpratheesh/dbmux/driver/options"
	"github.com/pcpratheesh/dbmux/entity"
)

func main() {
	mysql_1 := entity.Options{
		Name:     "mysql_1",
		Driver:   driver.MYSQL,
		Host:     "localhost",
		User:     "root",
		Password: "root",
		Database: "db",
		Port:     3306,
		ConnectionOptions: []entity.ConnectionOptions{
			options.SetConnMaxLifetime(10 * time.Second),
			options.SetConnMaxIdleTime(30 * time.Second),
			options.SetMaxOpenConns(5),
			options.SetMaxIdleConns(2),
		},
	}

	psql_1 := entity.Options{
		Name:     "psql_1",
		Driver:   driver.POSTGRES,
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Database: "db",
		Port:     5433,
	}

	mongo_1 := entity.Options{
		Name:     "mongo_1",
		Driver:   driver.MONGO,
		Host:     "localhost",
		User:     "admin",
		Password: "admin",
		Database: "db",
		Port:     27017,
	}

	var dbConnections = []entity.Options{
		mysql_1, mongo_1, psql_1,
	}

	// this will initiate the connection
	dbmuxConn := dbmux.New(dbConnections...)

	// get the mongo connection pool
	{
		mongoConnection, err := dbmuxConn.GetConnection("mongo_1")
		if err != nil {
			panic(err)
		}

		mongoDB := dbmux.GetConnectionPool[entity.MongoClient](mongoConnection)
		_ = mongoDB
	}

	// get the psql connection pool
	{
		psqlConnection, err := dbmuxConn.GetConnection("psql_1")
		if err != nil {
			panic(err)
		}

		psqlDB := dbmux.GetConnectionPool[entity.SqlDB](psqlConnection)
		_ = psqlDB
	}

	// get the psql connection pool
	{
		mysqlConnection, err := dbmuxConn.GetConnection("mysql_1")
		if err != nil {
			panic(err)
		}
		mysqlDB := dbmux.GetConnectionPool[entity.SqlDB](mysqlConnection)
		_ = mysqlDB
	}

	fmt.Println("dbmux connection initiated")

}
