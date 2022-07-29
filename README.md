<p align="center">
  <img 
    src="https://github.com/syke99/images/blob/main/dynaQ.png?raw=true"
  >
</p>

[![Go Reference](https://pkg.go.dev/badge/github.com/syke99/dynaQ.svg)](https://pkg.go.dev/github.com/syke99/dynaQ)
[![go reportcard](https://goreportcard.com/badge/github.com/syke99/dynaQ)](https://goreportcard.com/report/github.com/syke99/dynaQ)
[![License](https://img.shields.io/github/license/syke99/dynaQ)](https://github.com/syke99/dynaQ/blob/master/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/syke99/dynaQ)</br>
An extension for Go's sql package in the standard library to support dynamic queries directly from the database, as well as on individual database connections


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
    
    dq := dynaQ.NewDynaQ(db)
    
    // after creating a dynamic querier, just pass in whatever
    // query string you want, and whatever variable amouont of query
    // arguments you need.
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
    // Query the database
    rows, err := dq.DatabaseQuery("select * from testTable where id in (@p1, @p2, @p3, @p4)", 1, 2, 4, 7)
    if err != nil {
        panic(err)
    }
    
    // create a boolean to keep track of whether or not there's a new row to be checked
    newRow := true
	
    fmt.Println("-----------------")
    for newRow {
    	// the first value returned by rows.NextRow()
	// is a bool signaling whether there is another
	// row following the second value returned by
	// rows.NextRow(), which represents the current
	// row for this loop
    	if ok, row := rows.NextRow(); !ok {
		newRow = false
    	}
    	fmt.Println(fmt.Sprintf("row: %d", row.CurrentRow))
    	fmt.Println("-----------------")
    	// create a boolean to keep track of whether or not there's a new column to be checked
	newColumn := true
	for newColumn {
		if ok, column := row.NextColumn(); !ok {
			newColumn = false
		}
            	fmt.Println(fmt.Sprintf("column: %s, value: %v (type: %s)", column.Name, fmt.Sprintf("%v", column.Value), column.Type))
        }
	fmt.Println("-----------------")
    }
    
    //
    // this will output:
    // -----------------
    // row: 1
    // -----------------
    // column: id, value: 1 (type: int64)
    // column: name, value: ab (type: string)
    // column: cost, value: 2.10 (type: float64)
    // column: available, value: true (type: bool)
    // column: created-date, value: 2018-01-18 05:43 (type: time.Time)
    // -----------------
    // row: 2
    // -----------------
    // column: id, value: 2 (type: int64)
    // column: name, value: cd (type string)
    // column: cost, value: 1.55 (type float64)
    // column: available, value: false (type: bool)
    // column: created-date, value: 2018-01-14 06:28 (type: time.Time)
    // -----------------
    // row: 3
    // -----------------
    // column: id, value: 4 (type: int64)
    // column: name, value: gh (type: string)
    // column: cost, value: 2.76 (type: float64)
    // column: available, value: true (type: bool)
    // column: created-date, value: 2018-09-04 15:09 (type: time.Time)
    // -----------------
    // row: 4
    // -----------------
    // column: id, value: 7 (type: int64)
    // column: name, value: mn (type: string)
    // column: cost, value: 2.99 (type: float64)
    // column: available, value: false (type: bool)
    // column: created-date, value: 2019-02-11 06:22 (type: time.Time)
    // -----------------
}
```

### Connections, Transactions, and PreparedStatements

dynaQ also allows for using dynamic queries on database connections. After creating your dynamic querier with `dynaQ.NewDynaQ(db *sql.DB)`,
you can call `NewDqConn(conn *sql.Conn)` before querying to query on a specific database connection

### dynaQ Usage Ideas

1. SSR with html templates and dynamic queries
2. Creating dynamic slices with values of various types from Database queries
3. And more!! The possibilities are endless!!

Who?
====

This library was developed by Quinn Millican ([@syke99](https://github.com/syke99))


## License

This repo is under the MIT license, see [LICENSE](LICENSE) for details.
