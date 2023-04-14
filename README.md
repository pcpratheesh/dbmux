**Package Name:** dbmux

**Description:**

dbmux is a Go package that provides a simple but powerful way to connect to multiple databases using different drivers. It abstracts away the underlying implementation details and allows you to switch between database drivers without changing your code.

The package uses the `database/sql` interface to provide a unified API for interacting with databases. It supports popular database drivers such as MySQL, PostgreSQL, and MongoDB, and can be extended to support other database drivers.

**Installation:**

You can install the package using go get command:

```sh
go get github.com/pcpratheesh/dbmux
```

**Usage:**

To use this package, you first need to import it into your project:

```go
import "github.com/pcpratheesh/dbmux"
```

Next, you can create a new instance of `MultiConn` struct by calling the `New` function and passing in any desired options:

```go
    mySQLConnection := entity.Options{
        Name :  "mysql_1",
        Driver:   driver.MYSQL,
        Host:     "localhost",
        Port:     3306,
        Username: "root",
        Password: "password",
        Database: "testdb",
    }

    postgresConnection := entity.Options{
        Name :  "psql_1",
        Driver:   driver.POSTGRES,
        Host:     "localhost",
        Port:     5432,
        Username: "user",
        Password: "password",
        Database: "testdb",
    }

    var dbConnections = []entity.Options{
        mySQLConnection, postgresConnection,
    }

    // this will initiate the connection
    multiConn := dbmux.New(dbConnections...)
```

Once you have an instance of `MultiConn`, you can retrieve a connection object from a specific driver's connection pool by calling the `GetConnection` method and providing the desired driver name:

```go
    psqlConnection, err := MultiConn.GetConnection("psql_1")
    if err != nil {
        panic(err)
    }

    psqlDB := dbmux.GetConnectionPool[entity.SqlDB](psqlConnection)

```

You can then use the returned connection (`psqlDB`) objects to execute SQL queries, perform transactions, and other database operations.


see [examples](examples/init.go) for more 


## entity.Options
| Key               | Description                                                                                               |
|-------------------|-----------------------------------------------------------------------------------------------------------|
| Name              | specifies a name for this connection.                                                                     |
| Driver            | specifies the type of database driver to use ( driver.MYSQL , driver.POSTGRES, driver.MONGO )             |
| Host              | specifies the host of the database server.                                                          |
| User              | specifies the user to be used to connect to the database server.                                          |
| Password          | specifies the password for the specified user.                                                            |
| Database          | specifies the name of the database to connect to.                                                         |
| Port              | specifies the port number on which the server is listening.                                         |
| ConnectionOptions | field is used to specify additional connection pool settings                                              |


### ConnectionOptions
|         Function Name      | Description                                                                                       |
|----------------------------|---------------------------------------------------------------------------------------------------|
| options.SetConnMaxLifetime | sets the maximum amount of time a connection can remain open before being closed.                 |
| options.SetConnMaxIdleTime | sets the maximum amount of time a connection can remain idle in the pool before being closed.     |
| options.SetMaxOpenConns    | sets the maximum number of open connections that can exist in the pool at any given time.         |
| options.SetMaxIdleConns    | sets the maximum number of idle (unused) connections that can exist in the pool at any given time |

see [example](examples/connection-options.go) 

## The importance of `dbmux.GetConnectionPool[entity.SqlDB](psqlConnection)`

The `dbmux.GetConnectionPool[entity.SqlDB](psqlConnection)` code is important because it retrieves a database connection pool object of a specific type that implements the `SqlPoolInterface`. In this case, it retrieves an instance of a PostgreSQL connection pool from the `psqlConnection` variable.

The `entity.SqlDB` type parameter specifies the type of the returned connection pool object. This object can then be used to perform various operations on the PostgreSQL database, such as executing queries or transactions, through the methods provided by the `database/sql.DB` struct.

Similarly, the `dbmux.GetConnectionPool[entity.MongoClient](mongoConnection)` code retrieves a connection pool object for a Mongo database connection.

By using `dbmux.GetConnectionPool` to retrieve the connection pool objects, we can centralize and abstract away the logic required to establish and manage connections to different types of databases. This simplifies the codebase, reduces code duplication, and makes it easier to maintain the application's database layer. 

Overall, the importance of `dbmux.GetConnectionPool[entity.SqlDB]` or `dbmux.GetConnectionPool[entity.MongClient]` lies in its ability to provide a seamless interface for connecting to different types of databases, allowing to focus on writing business logic rather than worrying about low-level database connection details.

## Contributing:

If you find any bugs or issues with the package, please report them on the project's GitHub page. Contributions are also welcome! If you'd like to contribute to the package, please submit a pull request with your changes.

### License:

This package is licensed under the MIT license. See the LICENSE file for more information.


