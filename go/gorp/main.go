package main

import (
	mysql "go-gorp-demo/db"
	person "go-gorp-demo/tables"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {
	dbmap := mysql.InitDB()
	defer dbmap.Db.Close()

	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	// Register Table
	person.InitTable(dbmap)

	// Create Table
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	// Specific operation
	person.Insert(dbmap)
}
