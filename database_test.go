package main

import (
	"testing"
)

func Test_Database__Scan(t *testing.T) {
	const expected = 20
	config := config()
	d := Database{config}
	results := make(chan Result)
	count := 0

	go func() {
		for result := range results {
			var _ = result
			count++
		}
	}()

	if err := d.Scan(results); err != nil {
		t.Fatal(err)
	}
	close(results)

	if count != expected {
		t.Errorf("Expected %d results, received %d", expected, count)
	}
}
