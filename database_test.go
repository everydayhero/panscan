package main

import (
	"testing"
)

func TestScan(t *testing.T) {
	const expected = 24
	config := config()
	d := NewDatabase(config)
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
