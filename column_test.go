package main

import (
	"testing"
)

const (
	TestDBName = "test_db"
)

func Test_Column__Scan_CreditCards(t *testing.T) {
	const expected = 20
	config := config()
	db := open(config)
	defer db.Close()

	column := Column{db: db, Config: config, Database: TestDBName, Table: "creditcards", Name: "ccnumber"}
	results := make(chan Result)
	count := 0

	go func() {
		for result := range results {
			var _ = result
			count++
		}
	}()

	if err := column.Scan(results); err != nil {
		t.Fatal(err)
	}
	close(results)

	if count != expected {
		t.Errorf("Expected %d results but received %d", expected, count)
	}
}

func Test_Column__Scan_Posts(t *testing.T) {
	const expected = 1
	config := config()
	db := open(config)
	defer db.Close()

	column := Column{db: db, Config: config, Database: TestDBName, Table: "posts", Name: "body"}
	results := make(chan Result)
	count := -1

	go func() {
		for result := range results {
			var _ = result
			count++
		}
	}()

	if err := column.Scan(results); err != nil {
		t.Fatal(err)
	}
	results <- Result{} // For some reason we have to send another object otherwise no objects are received if only one is sent. Possible bug?
	close(results)

	if count != expected {
		t.Errorf("Expected %d results but received %d", expected, count)
	}
}
