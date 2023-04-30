package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library

	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

type handler interface {
	//logger()
	writer()
	reader()
}

type database struct {
}

type rawStorage struct {
	path       string
	fileFormat string
	totalSize  float64
}

type customerAddress struct {
	street string
	city   string
	state  string
}

type customer struct {
	id      int
	name    string
	age     int
	address customerAddress
}

func main() {
	var jsonRawStorage rawStorage
	var csvRawStorage rawStorage
	jsonRawStorage.path = "data_json/customer_20230412.json"
	jsonRawStorage.fileFormat = path.Ext(jsonRawStorage.path)

	csvRawStorage.path = "data_csv/customer_20230415.csv"
	jsonRawStorage.fileFormat = path.Ext(csvRawStorage.path)

	nameSQLiteFile := "sqlite-database.db" // Create SQLite file

	initializeDb(nameSQLiteFile)
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File

	defer sqliteDatabase.Close() // Defer Closing the database
	createTable(sqliteDatabase)  // Create Database Tables

	fileExtension := path.Ext(jsonRawStorage.path)
	if fileExtension == ".csv" {
		//extract := readCsvFile(filePath)
		extract_json := jsonRawStorage.reader()
		fmt.Println(extract_json)
	} else if fileExtension == ".json" {
		extract_json := jsonRawStorage.reader()
		fmt.Println(extract_json)
	}
}

func readCsvFile(filePath string) []customer {
	recordsCustomer := []customer{}
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data_extract, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	fmt.Println(data_extract)
	_, data_content := data_extract[0], data_extract[1:]
	for _, data := range data_content {
		fmt.Println(data[0])
		age_int, err := strconv.Atoi(data[0])
		if err != nil {
			log.Fatal("can not convert to int: ", age_int, err)
		}
		recordsCustomer = append(recordsCustomer, customer{age_int, data[1], 7, customerAddress{"", "", ""}})
		fmt.Println(recordsCustomer)
	}
	return recordsCustomer
}

func (r rawStorage) reader(fileNames ...string) customer {

	customerJson, err := os.Open(r.path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer customerJson.Close()

	byteJson, _ := ioutil.ReadAll(customerJson)

	var records map[string]interface{} // -> if we dont know the structure of the json
	err = json.Unmarshal(byteJson, &records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	//build_str :=
	for key, element := range records {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	nestedMap := records["address"].(map[string]interface{})
	customerReturn := customer{1, records["lastName"].(string), int(records["age"].(float64)), customerAddress{nestedMap["streetAddress"].(string), nestedMap["city"].(string), nestedMap["state"].(string)}}

	return customerReturn
}

func initializeDb(name string) {
	file, err := os.Create(name) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db initialized")
}

func createTable(db *sql.DB) {
	createCustomerTableSQL := `CREATE TABLE IF NOT EXISTS CUSTOMER (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"firstName" TEXT,
		"lastName" TEXT,
		"name" TEXT,
		"age" TEXT
		"address" TEXT		
		"streetAddress" TEXT,
        "city" TEXT,
        "state" "Louisiana" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create CUSTOMER table...")
	statement, err := db.Prepare(createCustomerTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("CUSTOMER table created")
}
