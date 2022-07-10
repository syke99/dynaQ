![dynaQ logo](https://github.com/syke99/images/blob/main/dynaQ.png?raw=true)

[![Go Reference](https://pkg.go.dev/badge/github.com/syke99/dynaQ.svg)](https://pkg.go.dev/github.com/syke99/dynaQ)
[![go reportcard](https://goreportcard.com/badge/github.com/syke99/dynaQ)](https://goreportcard.com/report/github.com/syke99/dynaQ)</br>
An extension for Go's sql package in the standard library to support dynamic queries directly from the database, as well as on database connections, prepared statements and transactions


How do I use dynaQ?
====

### Installation

```
go get github.com/syke99/dynaQ
```

### Basic Usage

```go
package main

import (
	"database/sql"
	"fmt"

	"github.com/syke99/dynaQ"
	_ "github.com/go-sql-driver/mysql"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
	
  // create an instance of your database
    db, err := sql.Open("mysql", dsn(""))
    if err != nil {
	    panic(err)
    }
    defer db.Close()
    
    // create a new dynamic querier with your database instance,
    // you can either pass in your own time format to match the
    // format of time values stored in your database, or if they
    // match the format "2006-01-02 15:04", you can pass in ""
    // and dynaQ will default to that format
    dq := dynaQ.NewDynaQ(db, "")
    
    // after creating a dynamic querier, just pass in whatever
    // query string you want, and whatever variable amouont of query
    // arguments you need.
    // you can query for multiple rows, or just one row at once
    //    
    // testTable:
    // __| id | name | cost | available |   created-date   |__
    // ---------------------------------------------------------
    //   |  1 |  ab  | 2.10 |   true    | 2018-01-18 05:43 |
    //   |  2 |  cd  | 1.55 |   false   | 2018-01-14 06:28 |
    //   |  3 |  ef  | 3.78 |   true    | 2018-06-27 09:59 |
    //   |  4 |  gh  | 2.76 |   true    | 2018-09-04 15:09 |
    //   |  5 |  ij  | 8.13 |   true    | 2019-01-01 23:43 |
    //   |  6 |  kl  | 4.45 |   false   | 2019-01-19 10:14 |
    //   |  7 |  mn  | 2.99 |   false   | 2019-02-11 06:22 |
    //
    // single row:
    row, err := dq.DatabaseQueryRow("select * from testTable where id in (@p1, @p2, @p3, @p4)", 1, 2, 4, 7)
    if err != nil {
	    panic(err)
    }
    
    // a dynaQ column holds the type in field
    // <columnVariable>.Type for easy type
    // assertion later on
    for _, column := range row {
	     fmt.Sprintf("%s: %v (type: %s)", column.Name, fmt.Sprintf("%v", column.Value), column.Type)
	     fmt.Println("-----------------")
    }
    //
    // this will output:
    // -----------------
    // id: 1 (type: int64)
    // name: ab (type: string)
    // cost: 2.10 (type: float64)
    // available: true (type: bool)
    // created-date: 2018-01-18 05:43 (type: time.Time)
    // -----------------
	
	
    // multiple rows:
    rows, err := dq.DatabaseQuery("select * from testTable where id in (@p1, @p2, @p3, @p4)", 1, 2, 4, 7)
    if err != nil {
        panic(err)
    }
    
    newRow := true
	
    fmt.Println("-----------------")
    for newRow {
    	ok, row := rows.NextRow()
	if !ok {
		newRow = false
    	}
	for _, column := range row {
            	fmt.Sprintf("%s: %v", column.Name, fmt.Sprintf("%v", column.Value))
            	fmt.Println("-----------------")
        }
    }
    
    //
    // this will output:
    // -----------------
    // id: 1 (type: int64)
    // name: ab (type: string)
    // cost: 2.10 (type: float64)
    // available: true (type: bool)
    // created-date: 2018-01-18 05:43 (type: time.Time)
    // -----------------
    // id: 2 (type: int64)
    // name: cd (type string)
    // cost: 1.55 (type float64)
    // available: false (type: bool)
    // created-date: 2018-01-14 06:28 (type: time.Time)
    // -----------------
    // id: 4 (type: int64)
    // name: gh (type: string)
    // cost: 2.76 (type: float64)
    // available: true (type: bool)
    // created-date: 2018-09-04 15:09 (type: time.Time)
    // -----------------
    // id: 7 (type: int64)
    // name: mn (type: string)
    // cost: 2.99 (type: float64)
    // available: false (type: bool)
    // created-date: 2019-02-11 06:22 (type: time.Time)
    // -----------------
}
```

### Connections, Transactions, and PreparedStatements

dynaQ also allows for using dynamic queries on database connections, database transactions, and prepared statements
after creating your dynamic querier with `dynaQ.NewDynaQ(db *sql.DB)`, you can call `NewDqConn(conn *sql.Conn)`,
`NewDqTransaction(tx *sql.Tx)`, or `NewDqPreparedStatement(query string)` respectively before querying

### dynaQ Usage Ideas

1. SSR with html templates and dynamic queries
2. Creating dynamic slices with values of various types from Database queries
3. And more!! The possibilities are endless!!

Who?
====

This library was developed by Quinn Millican (@syke99)


## License

This repo is under the MIT license, see [LICENSE](LICENSE) for details.
