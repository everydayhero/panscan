package main

import (
	"testing"
)

const (
	TestDBName = "test_db"
)

func TestGetColumns(t *testing.T) {
	const expected = 8
	config := config()
	db := open(config)
	defer db.Close()

	columns := GetColumns(db, config)

	if len(columns) != expected {
		t.Errorf("Expected %d columns but found %d", expected, len(columns))
	}
}

func TestScanOnCreditCards(t *testing.T) {
	const expected = 22
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

func TestScanOnPosts(t *testing.T) {
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
