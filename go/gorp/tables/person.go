package person

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/gorp.v1"
)

type Person struct {
	Id      int64
	Created int64
	FName   string
	LName   string
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// InitTable -
func InitTable(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(Person{}, "person_test").SetKeys(true, "Id")
}

// Insert -
func Insert(dbmap *gorp.DbMap) {
	person := Person{
		Created: time.Now().UnixNano(),
		FName:   "dingrui",
		LName:   "y",
	}

	startTime := time.Now()
	err := dbmap.Insert(&person)
	elapsed := time.Since(startTime)

	fmt.Println("Spand Time:", elapsed)
	checkErr(err, "Insert failed")
}
