package main

import (
	"testing"
)

func Test_MySQL__Columns(t *testing.T) {
	const expected = 2
	config := config()
	db := open(config)
	m := MySQL{config}
	defer db.Close()

	columns := m.Columns(db)

	if len(columns) != expected {
		t.Errorf("Expected %d columns but found %d: ", expected, len(columns), columns)
	}
}

func Test_MySQL__Scan(t *testing.T) {
	const expected = 20
	config := config()
	m := MySQL{config}
	results := make(chan Result)
	count := 0

	go func() {
		for result := range results {
			var _ = result
			count++
		}
	}()

	if err := m.Scan(results); err != nil {
		t.Fatal(err)
	}
	close(results)

	if count != expected {
		t.Errorf("Expected %d results, received %d", expected, count)
	}
}
