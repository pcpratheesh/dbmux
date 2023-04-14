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

	var dbConnections = []entity.Options{
		mysql_1,
	}

	// this will initiate the connection
	dbmuxConn := dbmux.New(dbConnections...)

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
