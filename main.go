package main

import (
	"fmt"

	//"strings"
	"path"
	//"reflect"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

/*
To-Do:
- add type mapping: then maps for tables does not need anymore "string"
- for function "prepareSQL":
- - should work for different types (view, table) -> interface?
- - need type of database for different SQL styles
- - implement type mapping from Go-Types to DB-types
- table-struct: how to best create table columns and types?

*/
// Every different databse systems needs individual function for creating Statement
// -> different sql languages

// special tables need to inherit execute from the database

/*
how to structure classes for different db's
- Database (abstract)
- - Postgres (or abstract on this level)
- - SQLite
- - ...

- methods for DB: (interfaces)
- - execute
- - prepare
- - read/ query

Classes
Tables: (tables abstract?):
- LandingTables
- tables
- views

methods:
- insert
- create
- merge
- update

Datatypes
- Postgres
  - VARCHAR
  - STRING
  - ...

// Abstract Class:
// Tables:
// -> uses method "write to db"

// Sub Classes:
// special tables;
// implement interfaces for specific types
*/
type handler interface {
	//logger()
	writer()
	reader()
}

var db database

func main() {
	fmt.Println("test")
	var jsonRawStorage rawStorage
	var csvRawStorage rawStorage
	jsonRawStorage.path = "data_json/"
	jsonRawStorage.fileFormat = ".json"

	csvRawStorage.path = "data_csv/customer_20230415.csv"
	csvRawStorage.fileFormat = path.Ext(csvRawStorage.path)

	db.path = "./db/"
	db.nameSQLiteFile = "sqlite-database.db"
	db.instance = db.initializeDb()
	defer db.instance.Close() // Defer Closing the database

	// table:
	var customerA customer
	var customerB customer

	var ordersA order
	var ordersB order

	customerA.Id = 1
	customerA.Firstname = "Jose"
	customerA.Lastname = "Al"
	customerA.Age = 36

	customerB.Id = 2
	customerB.Firstname = "Allen"
	customerB.Lastname = "Cuck"
	customerB.Age = 36

	ordersA.Id = 1
	ordersA.Firstname = "Jose"
	ordersA.Lastname = "Al"
	ordersA.Object = "Book"
	ordersA.Amount = 5
	ordersA.Shipped = true

	ordersB.Id = 2
	ordersB.Firstname = "Bold"
	ordersB.Lastname = "Eric"
	ordersB.Object = "Movie"
	ordersB.Amount = 1
	ordersB.Shipped = true

	customerTable := customerTable{}

	orderTable := orderTable{}
	orderTable.C = []order{ordersA, ordersB}
	customerTable.C = []customer{customerA, customerB}
	//db.createTable(orderTable)
	db.createTable(orderTable)
	//ReadStruct(orderTable)
	db.insert(orderTable)

	//query := "Select * from ordertable"
	//res, _ := db.instance.Exec(query)
	//fmt.Println(res)

}
